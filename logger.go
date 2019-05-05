package logapi

import "time"

// define delvel const for object
const (
	DEVEL_DEBUG = 1
	DEVEL_INFO  = 2
	DEVEL_WARN  = 3
	DEVEL_ERROR = 4
	DEVEL_FATAL = 5
)

// define all api interface

// Option this option is a interface for every implement package
type Option interface {

	// GetLevel return the base level object
	GetLevel() string

	GetProperties() map[string]string

	// GetAppender get the all appender
	GetAppenders() map[string]Appender
}

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

// FactoryRegister the interface provide the factory register
type FactoryRegister interface {

	// CreateLogger return the instance logger with multiopts
	CreateLogger(multiopts ...Option) (Logger, error)

	CreateLoggerBean() LoggerBean
}

// LoggerFactory define the base logger factory manager
type LoggerFactory interface {

	// NewLoggerBean define
	NewLoggerBean() LoggerBean

	// CreateLogger get the base logger
	GetLogger() Logger
}

// LoggerBean Message object is logger info use common
type LoggerBean interface {
	LogBinary(key string, value []byte)

	LogByteString(key string, values []byte)

	LogByteStringArray(key string, values [][]byte)

	// LogString log message string
	LogString(string, string)

	LogStringArray(key string, values []string)

	LogBool(key string, value bool)

	LogBoolArray(key string, values []bool)

	LogInt8(string, int8)

	LogInt8Array(key string, nums []int8)

	LogInt(string, int)

	LogIntArray(key string, nums []int)

	LogInt16(key string, value int16)

	LogInt16Array(key string, nums []int16)

	LogInt32(key string, value int32)

	LogInt32Array(key string, nums []int32)

	LogInt64(key string, value int64)

	LogInt64Array(key string, nums []int64)

	LogUint8(key string, value uint8)

	LogUint8Array(key string, nums []uint8)

	LogUint(key string, value uint)

	LogUintArray(key string, nums []uint)

	LogUint16(key string, value uint16)

	LogUint16Array(key string, nums []uint16)

	LogUint32(key string, value uint32)

	LogUint32Array(key string, nums []uint32)

	LogFloat32(key string, value float32)

	LogFloat32Array(key string, values []float32)

	LogFloat64(key string, value float64)

	LogFloat64Array(key string, values []float64)

	LogDuration(key string, value time.Duration)

	LogDurationArray(key string, values []time.Duration)

	LogTime(key string, value time.Time)

	LogTimeArray(key string, values []time.Time)
}
