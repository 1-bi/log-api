package logapi

import (
	"log"
)

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
	staticLogger = logger

	return factory
}

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
