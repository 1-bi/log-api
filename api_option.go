package logapi

//  Option this Logger option is a interface for every implement package
type Option interface {

	// GetLevel return the base level object
	GetLevel() string

	// SetLoggerPattern define logger pattern for match loggger name
	SetLoggerPattern(loggerPattern string)

	// GetLoggerPattern get logger pattern
	GetLoggerPattern() string

	// SetAdditivity default is true , set the value is false . log info would be output in private .
	SetAdditivity(newadditivity bool)

	// GetAdditivity get the Additivity status
	GetAdditivity() bool

	// GetProperties get the append value
	GetProperties() map[string]string

	// GetAppender get the all appender
	GetAppenders() map[string]Appender
}
