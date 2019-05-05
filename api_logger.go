package logapi

// Logger create binder interface
type Logger interface {
	// ----
	IsDebugEnabled() bool

	IsInfoEnabled() bool

	IsWarnEnabled() bool

	IsErrorEnabled() bool

	Debug(msg string, logbean StructBean)

	Info(msg string, logbean StructBean)

	Warn(msg string, logbean StructBean)

	Error(msg string, logbean StructBean)
}

// LoggerFactory define the base logger factory manager
type LoggerFactory interface {

	// NewStructBean define
	NewLoggerBean() StructBean

	// CreateLogger get the base logger
	GetLogger() Logger
}
