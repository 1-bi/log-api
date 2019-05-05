package logapi

/**
 * create binder interface
 */
type Logger interface {
	// ----
	IsDebugEnabled() bool

	IsInfoEnabled() bool

	IsWarnEnabled() bool

	IsErrorEnabled() bool

	Debug(msg string, logbean LoggerBean)

	Info(msg string, logbean LoggerBean)

	Warn(msg string, logbean LoggerBean)

	Error(msg string, logbean LoggerBean)
}

// LoggerFactory define the base logger factory manager
type LoggerFactory interface {

	// NewLoggerBean define
	NewLoggerBean() LoggerBean

	// CreateLogger get the base logger
	GetLogger() Logger
}
