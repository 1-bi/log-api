package logapi

var staticLogger Logger
var StaticLoggerBean LoggerBean

// GetLogger define the custom logger ,  input module to load objectg
func GetLogger(modName string) Logger {
	return staticLogger
}

func GenerateLoggerBean() LoggerBean {
	return StaticLoggerBean
}
