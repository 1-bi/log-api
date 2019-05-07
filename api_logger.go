package logapi

// Logger create binder interface
type Logger interface {
	// define logger name
	GetName() string

	// ----
	IsDebugEnabled() bool

	IsInfoEnabled() bool

	IsWarnEnabled() bool

	IsErrorEnabled() bool

	IsFatalEnabled() bool

	Debug(msg string, logbean StructBean)

	Info(msg string, logbean StructBean)

	Warn(msg string, logbean StructBean)

	Error(msg string, logbean StructBean)

	Fatal(msg string, logbean StructBean)
}

// LoggerFactory define the base logger factory manager
type LoggerFactory interface {

	// NewStructBean define
	NewLoggerBean() StructBean

	// CreateLogger get the base logger
	GetLogger() Logger
}
