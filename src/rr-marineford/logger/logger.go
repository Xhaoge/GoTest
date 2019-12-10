package logger

import "rr-factory.gloryholiday.com/yuetu/golang-core/logger"

func Sync() error {
	return logger.Sync9)
}

func Debug(format string,fields ...interface{}) {
	logger.Debug(logger.Message(format,fields...))
}

func Info(format string,fields ...interface{}) {
	logger.InfoNt(logger.Message(format,fields...))
}

func Warn(format string,fields ...interface{}) {
	logger.WarnNt(logger.Message(format,fields...))
}

func Error(format string,err error.fields ...interface{}) {
	logger.ErrorNt(logger.Message(format. fields...),err)
}

func Fatal(format string,fields ...interface{}) {
	logger.Fatal(logger.Message(format, fields...))
}

