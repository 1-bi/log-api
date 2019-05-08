package logapi

import (
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

// NewStructBean create new struct bean for struct logger message
func NewStructBean() StructBean {
	return _globalLoggerBean.CreateStructBean()
}

// LoggerFactoryRegister create new logger factory manager by logger provider register
func RegisterLoggerFactory(reg StructLoggerRegister, opts ...Option) ([]Logger, error) {

	_globalLoggerBean = reg

	loggers, err := reg.CreateLogger(opts...)

	if err != nil {
		return nil, err
	}

	// --- init logger holder  ----
	for _, logVal := range loggers {
		_logPatternHolder[logVal.GetName()] = logVal
	}

	return loggers, nil

}

var _logPatternHolder = make(map[string]Logger)

//  define holder
func findCloseLoggerByLoggerPattern(loggerName string) Logger {

	// --- find and match logger pattern
	var partedNames = strings.Split(loggerName, ".")
	var patternSize = len(partedNames) - 1
	var parentParteds = make([]string, patternSize)

	for i := 0; i < patternSize; i++ {
		parentParteds[i] = partedNames[i]
	}
	var parentPattern = strings.Join(parentParteds, ".")
	var runtimeLogger = _logPatternHolder[parentPattern]

	// --- get the main logger ,if not logger mapper
	if runtimeLogger == nil && patternSize == 0 {
		// --- get the main logger directory
		runtimeLogger = _logPatternHolder["main"]
	}

	if runtimeLogger == nil {
		// create new logger pattern
		runtimeLogger = findCloseLoggerByLoggerPattern(parentPattern)
	}
	return runtimeLogger

}

// GetLogger define the custom logger , loggername is mark for identifing logger function . get close patten logger
func GetLogger(loggerName string) Logger {

	// get logger by id
	var runtimeLogger = _logPatternHolder[loggerName]

	if runtimeLogger != nil {
		return runtimeLogger
	}

	runtimeLogger = findCloseLoggerByLoggerPattern(loggerName)

	// ---- get root logger ---
	if runtimeLogger == nil && _logPatternHolder[loggerName] == nil {
		// use default logger
		runtimeLogger = _logPatternHolder["main"]
	}

	return runtimeLogger

}
