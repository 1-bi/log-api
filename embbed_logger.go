package logapi

import (
	"github.com/1-bi/log-api/embbed"
	"log"
)

type logger struct {
	name         string
	runtimeLevel byte

	additivity bool
}

func NewEmbbedLogger(runtimeLevel byte, loggerName string) *logger {

	log := new(logger)
	log.runtimeLevel = runtimeLevel
	log.name = loggerName

	return log
}

func (log *logger) GetName() string {
	return log.name
}

func (log *logger) SetParentLogger(parentLogger Logger) {

}

func (log *logger) GetParentLogger() Logger {
	return nil
}

func (log *logger) IsDebugEnabled() bool {
	return log.runtimeLevel == DEBUG
}

func (log *logger) IsInfoEnabled() bool {
	return log.runtimeLevel == INFO
}

func (log *logger) IsWarnEnabled() bool {
	return log.runtimeLevel == WARN
}

func (log *logger) IsErrorEnabled() bool {
	return log.runtimeLevel == ERROR
}

func (log *logger) IsFatalEnabled() bool {
	return log.runtimeLevel == FATAL
}

// Debug debug logger message object
func (myself *logger) Debug(msg string, msgObj StructBean) {

	embbed.Trace.Println(msg)
}

func (myself *logger) Info(msg string, msgObj StructBean) {
	embbed.Info.Println(msg)
}

func (myself *logger) Warn(msg string, msgObj StructBean) {
	embbed.Warning.Println(msg)
}

func (myself *logger) Error(msg string, msgObj StructBean) {
	embbed.Error.Println(msg)

}

func (myself *logger) Fatal(msg string, msgObj StructBean) {
	log.Fatal(msg)
}
