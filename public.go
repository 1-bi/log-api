package logapi

import (
	"log"
	"strings"
)

// define delvel const for object
const (
	DEBUG = 1
	INFO  = 2
	WARN  = 3
	ERROR = 4
	FATAL = 5
)

// open public api source in this file

func getLogger(loggerName string) Logger {

	logger, err := _globalLoggerBean.CreateLogger(loggerName, _globalOpts...)

	if err != nil {
		log.Panic(err)
		return nil
	}

	return logger
}

// NewStructBean create new struct bean for struct logger message
func NewStructBean() StructBean {
	return _globalLoggerBean.CreateStructBean()
}

// LoggerFactoryRegister create new logger factory manager by logger provider register
func RegisterLoggerFactory(reg StructLoggerRegister, opts ...Option) (Logger, error) {

	_globalOpts = opts

	_globalLoggerBean = reg

	logger, err := reg.CreateLogger("main", opts...)

	if err != nil {
		return nil, err
	}

	return logger, nil

}

var _logPatternHolder = make(map[string]Logger)

//  define holder
func InitLoggerPattern(loggerNames []string, defaultMainLogger Logger) {
	// use init logger
	for _, logName := range loggerNames {
		_logPatternHolder[logName] = getLogger(logName)
	}

	_logPatternHolder["main"] = defaultMainLogger

}

// GetLogger define the custom logger , loggername is mark for identifing logger function
func GetLogger(loggerName string) Logger {

	// get logger by id
	var runtimeLogger = _logPatternHolder[loggerName]

	if runtimeLogger != nil {
		return runtimeLogger
	}

	var parentParteds = make([]string, 0)

	// --- find and match logger pattern
	var partedNames = strings.Split(loggerName, ".")

	for i := 0; i < len(partedNames)-1; i++ {
		parentParteds = append(parentParteds, partedNames[i])
	}
	var parentPattern = strings.Join(partedNames, ".")
	runtimeLogger = _logPatternHolder[parentPattern]
	if runtimeLogger == nil {
		// create new logger pattern
		_logPatternHolder[parentPattern] = getLogger(parentPattern)
		runtimeLogger = _logPatternHolder[parentPattern]
	}

	// ---- get root logger ---
	if runtimeLogger == nil && _logPatternHolder[loggerName] == nil {
		// use default logger
		runtimeLogger = _logPatternHolder["main"]
	}

	return runtimeLogger

}
