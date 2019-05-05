package logapi

// Option this option is a interface for every implement package
type Option interface {

	// GetLevel return the base level object
	GetLevel() string

	GetProperties() map[string]string

	// GetAppender get the all appender
	GetAppenders() map[string]Appender
}
