package logger_test

import (
	"runtime"
	"testing"

	"github.com/gogoclouds/gogo-services/framework/app/logger"
)

var logConf logger.Config

func init() {
	logConf = logger.NewConfig(
		logger.WithLevel(logger.LogLevel_Error),
		logger.WithTimeFormat("2006-01-02 15:04:05.000"),
		logger.WithFilename("server"),
		logger.WithFileMaxSize(10),
		logger.WithFileMaxAge(6*30),
		logger.WithFileCompress(true),
	)
	logger.InitZapLogger(logConf)

	//logger.InitSlogLogger(logConf)
}

func TestLogger(t *testing.T) {
	logger.Debug("The is ", "Debug", "debug")
	logger.Info("The is ", "Info", "info")
	logger.Error("The is ", "Error", "error")

	go logConf.SetLevel(logger.LogLevel_Debug)
	runtime.Gosched()
	logger.Debugf("The is %s", "Debugf")
	logger.Infof("The is %s", "Infof")
	logger.Errorf("The is %s", "Errorf")

	// output:
	// 2023-11-06 23:55:03.638	INFO	logger/logger_test.go:28	The is 	{"Info": "info"}
	// 2023-11-06 23:55:03.638	ERROR	logger/logger_test.go:29	The is 	{"Error": "error"}
	// 2023-11-06 23:55:03.638	DEBUG	logger/logger_test.go:33	The is Debugf
	// 2023-11-06 23:55:03.638	INFO	logger/logger_test.go:34	The is Infof
	// 2023-11-06 23:55:03.638	ERROR	logger/logger_test.go:35	The is Errorf
}
