package logapi

import "log"

// define delvel const for object
const (
	DEVEL_DEBUG = 1
	DEVEL_INFO  = 2
	DEVEL_WARN  = 3
	DEVEL_ERROR = 4
	DEVEL_FATAL = 5
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

//
func NewLoggerBean() StructBean {
	return _globalLoggerBean.CreateStructBean()
}

// LoggerFactoryRegister create new logger factory manager by logger provider register
func LoggerFactoryRegister(reg StructLoggerRegister, opts ...Option) error {

	_globalOpts = opts

	_globalLoggerBean = reg

	return nil

}
