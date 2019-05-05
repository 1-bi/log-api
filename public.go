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
	return _globalLogger
}

//
func NewLoggerBean() LoggerBean {
	return _globalLoggerBean
}

// NewLoggerFactory create new logger factory manager by logger provider register
func NewLoggerFactory(reg FactoryRegister, opts ...Option) LoggerFactory {

	logger, err := reg.CreateLogger(opts...)

	if err != nil {
		log.Panic(err)
		return nil
	}

	// create another object
	var factory = &baseLoggerFactory{logger, reg}

	// --- register logger object to instance ,gloal define
	_globalLogger = logger

	reg.CreateLoggerBean()

	return factory
}
