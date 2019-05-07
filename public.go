package logapi

import "log"

// define delvel const for object
const (
	DEBUG = 1
	INFO  = 2
	WARN  = 3
	ERROR = 4
	FATAL = 5
)

// open public api source in this file

// GetLogger define the custom logger , loggername is mark for identifing logger function
func GetLogger(loggerName string) Logger {

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
