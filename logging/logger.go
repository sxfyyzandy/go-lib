package logging

import (
	"fmt"

	"github.com/cihub/seelog"
)

var (
	infoLogger   ILogger
	accessLogger ILogger
)

func InitLogger(infoLogCfgFile, accessLogCfgFile string) {
	initZlog(infoLogCfgFile, accessLogCfgFile)
}

type Zlogger struct {
	Prefix string
	Level  int
	seelog.LoggerInterface
}

func GetAccessLogger() ILogger {
	return accessLogger
}

func GetInfoLogger() ILogger {
	return infoLogger
}

func initZlog(infoLogCfgFile, accessLogCfgFile string) {
	var err error

	infoLogger, err = newZloggerFromFile(infoLogCfgFile)
	if err != nil {
		panic(err)
	}

	accessLogger, err = newZloggerFromFile(accessLogCfgFile)
	if err != nil {
		panic(err)
	}
}

func newZloggerFromFile(fileName string) (ILogger, error) {

	fmt.Println("newZloggerFromFile, filename: ", fileName)
	loggerInterface, err := seelog.LoggerFromConfigAsFile(fileName)

	loggerInterface.SetAdditionalStackDepth(1)

	if err != nil {

		fmt.Println("new loggerinterface failed: ", err)
		return nil, err
	}

	logger := &Zlogger{
		Level:           LEVEL_INFO,
		LoggerInterface: loggerInterface,
	}

	return logger, nil
}

func (z *Zlogger) SetLevel(level int) {
	z.Level = level
}

func (z *Zlogger) SetPrefix(prefix string) {
	z.Prefix = prefix
}

func (z *Zlogger) Debug(v ...interface{}) {
	z.LoggerInterface.Debug(v...)
}

func (z *Zlogger) Debugf(format string, v ...interface{}) {
	z.LoggerInterface.Debugf(format, v...)
}

func (z *Zlogger) Info(v ...interface{}) {
	z.LoggerInterface.Info(v...)
}

func (z *Zlogger) Infof(format string, v ...interface{}) {
	z.LoggerInterface.Infof(format, v...)
}

func (z *Zlogger) Warn(v ...interface{}) {
	z.LoggerInterface.Warn(v...)
}

func (z *Zlogger) Warnf(format string, v ...interface{}) {
	z.LoggerInterface.Warnf(format, v...)
}

func (z *Zlogger) Error(v ...interface{}) {
	z.LoggerInterface.Error(v...)
}

func (z *Zlogger) Errorf(format string, v ...interface{}) {
	z.LoggerInterface.Errorf(format, v...)
}

func (z *Zlogger) Critical(v ...interface{}) {
	z.LoggerInterface.Critical(v...)
}

func (z *Zlogger) Criticalf(format string, v ...interface{}) {
	z.LoggerInterface.Criticalf(format, v...)
}

func (z *Zlogger) Trace(v ...interface{}) {
	z.LoggerInterface.Trace(v...)
}

func (z *Zlogger) Tracef(format string, v ...interface{}) {
	z.LoggerInterface.Tracef(format, v...)
}
