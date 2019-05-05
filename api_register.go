package logapi

// FactoryRegister the interface provide the factory register
type FactoryRegister interface {

	// CreateLogger return the instance logger with multiopts
	CreateLogger(multiopts ...Option) (Logger, error)

	CreateLoggerBean() LoggerBean
}
