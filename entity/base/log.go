package base

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

func init() {
	logConf := ConfInstance.Log
	level, err := logrus.ParseLevel(logConf.Level)
	if err != nil {
		panic(err)
	}
	logrus.SetLevel(level)
	logger := &lumberjack.Logger{
		Filename:   logConf.Filename,
		MaxSize:    logConf.MaxSize,
		MaxAge:     logConf.MaxAge,
		MaxBackups: 0,
		LocalTime:  true,
	}
	logrus.SetReportCaller(logConf.ReportCaller)
	if !logConf.OutputToConsole {
		logrus.SetOutput(logger)
	}
}
