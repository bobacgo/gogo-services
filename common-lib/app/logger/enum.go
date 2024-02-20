package logger

import "strings"

type LogLevel string

const (
	LogLevel_Debug LogLevel = "debug"
	LogLevel_Error LogLevel = "error"
	LogLevel_Info  LogLevel = "info"
)

var logLevelMap = map[string]LogLevel{
	"debug": LogLevel_Debug,
	"error": LogLevel_Error,
	"info":  LogLevel_Info,
}

func (l LogLevel) String() string {
	return string(l)
}

func StringToLevel(level string) LogLevel {
	lower := strings.ToLower(level)
	logLevel := logLevelMap[lower]
	return logLevel
}
