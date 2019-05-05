package logapi

// StructLoggerRegister the interface provide the factory register
type StructLoggerRegister interface {

	// CreateLogger return the instance logger with multiopts
	CreateLogger(loggerName string, multiopts ...Option) (Logger, error)

	// CreateStructBean
	CreateStructBean() StructBean
}
