package logger

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func InitializeLogger(runMode string) {
	var err error
	var config zap.Config

	if runMode == "production" {
		config = zap.NewProductionConfig()
	} else {
		config = zap.NewDevelopmentConfig()
	}

	// Logger Config
	encoderConfig := zap.NewProductionEncoderConfig()
	// Encoder TimeStamp Config
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// Encoder Level Config
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// Encoder Message Config
	encoderConfig.MessageKey = "message"
	// Encoder Stacktrace Config
	// encoderConfig.StacktraceKey = "" // Disable Stacktrace

	config.EncoderConfig = encoderConfig

	logger, err = config.Build()
	if err != nil {
		log.Fatalf("[Error]->Failed to initialize logger : %s", err)
	}
	
}


func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

func Warning(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	logger.Fatal(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	logger.Panic(msg, fields...)
}