package embbed

import "github.com/1-bi/log-api"

type logger struct {

	name         string
	runtimeLevel byte

	additivity   bool
}

func NewEmbbedLogger(runtimeLevel byte , loggerName string ) *logger {

	log := new (logger)
	log.runtimeLevel = runtimeLevel
	log.name = loggerName

	return log
}

func (log *logger) GetName() string {
	return log.name
}

func (log *logger) SetParentLogger(parentLogger logapi.Logger) {

}

func (log *logger) GetParentLogger() logapi.Logger {
	return nil
}


func (log *logger) IsDebugEnabled() bool {
	return log.runtimeLevel == logapi.DEBUG
}

func (log *logger) IsInfoEnabled() bool {
	return log.runtimeLevel == logapi.INFO
}

func (log *logger) IsWarnEnabled() bool {
	return log.runtimeLevel == logapi.WARN
}

func (log *logger) IsErrorEnabled() bool {
	return log.runtimeLevel == logapi.ERROR
}

func (log *logger) IsFatalEnabled() bool {
	return log.runtimeLevel == logapi.FATAL
}

// Debug debug logger message object
func (myself *logger) Debug(msg string, msgObj logapi.StructBean) {



}

func (myself *logger) Info(msg string, msgObj logapi.StructBean) {


}

func (myself *logger) Warn(msg string, msgObj logapi.StructBean) {


}

func (myself *logger) Error(msg string, msgObj logapi.StructBean) {


}

func (myself *logger) Fatal(msg string, msgObj logapi.StructBean) {


}