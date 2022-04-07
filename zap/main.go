package main

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.SugaredLogger

func init() {
	core := zapcore.NewCore(getEncoder(), getLogWriter(), zapcore.DebugLevel)
	logger := zap.New(core, zap.AddCaller())
	log = logger.Sugar()

}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./test.log",
		MaxSize:    10,
		MaxBackups: 20,
		MaxAge:     30,
		Compress:   true,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	return zapcore.NewConsoleEncoder(encoderConfig)
}

func main() {
	// log := zap.NewExample()
	// log, _ := zap.NewDevelopment()
	// log, _ := zap.NewProduction()

	for i := 0; i < 1000000; i++ {
		log.Debug("This is debug message")
		log.Info("This is info message")
		log.Info("This is info message w/ fields", zap.Int("age", 24), zap.String("name", "Newton"))
		log.Warn("This is warn message")
		log.Error("This is error message")
	}
	// log.Panic("This is panic message")
}
