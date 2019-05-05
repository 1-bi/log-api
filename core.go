package logapi

var _globalLogger Logger
var _globalLoggerBean LoggerBean

// inner function and struct  source in this file

// baseLoggerFactory create new default logger factory
type baseLoggerFactory struct {
	logger   Logger
	register FactoryRegister
}

func (myself *baseLoggerFactory) NewLoggerBean() LoggerBean {

	return myself.register.CreateLoggerBean()
}

func (myself *baseLoggerFactory) GetLogger() Logger {
	return myself.logger
}
