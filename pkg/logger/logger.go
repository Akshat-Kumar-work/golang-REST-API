package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var globalLogger *zap.Logger

type LogConfig struct {
	Level  string `mapstructure:"level"`
	Format string `mapstructure:"format"`
}

// Initialize sets up the global zap logger
// Initialize sets up the global Zap logger using values from config
func Initialize(config LogConfig) error {
	var zapConfig zap.Config

	// 1️⃣ Set default log level to INFO
	level := zapcore.InfoLevel

	// Try to read log level from config (e.g. "debug", "info", "error")
	// If it fails, fallback to INFO
	// log level info, will print info,warn,error will ignore debug
	// log level error will print error, will ignore info and debug
	// log level debug will print everything
	if err := level.UnmarshalText([]byte(config.Level)); err != nil {
		level = zapcore.InfoLevel
	}

	// 2️⃣ Choose log format
	// - "json" → production logs (machine readable)
	// - anything else → development logs (human readable)
	if config.Format == "json" {
		zapConfig = zap.NewProductionConfig()
	} else {
		zapConfig = zap.NewDevelopmentConfig()
		// Show log levels in color for easier reading in local/dev
		zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	// 3️⃣ Apply the log level (debug / info / error)
	zapConfig.Level = zap.NewAtomicLevelAt(level)

	// 4️⃣ Add common fields that appear in every log
	zapConfig.InitialFields = map[string]interface{}{
		"service": "students-api",
	}

	// 5️⃣ Build the logger
	logger, err := zapConfig.Build(
		zap.AddCaller(),      //enable caller info
		zap.AddCallerSkip(1), //skip logger package(will not show logger file as log caller)
	)
	if err != nil {
		return err
	}

	// 6️⃣ Save logger as a global logger
	globalLogger = logger
	return nil
}

func GetLogger() *zap.Logger {
	if globalLogger == nil {
		// Fallback to development logger if not initialized
		globalLogger, _ = zap.NewDevelopment()
	}
	return globalLogger
}

// Sync flushes any buffered log entries
func Sync() error {
	if globalLogger != nil {
		return globalLogger.Sync()
	}
	return nil
}

// Info logs an info message
func Info(msg string, fields ...zap.Field) {
	GetLogger().Info(msg, fields...)
}

// Error logs an error message
func Error(msg string, fields ...zap.Field) {
	GetLogger().Error(msg, fields...)
}

// Debug logs a debug message
func Debug(msg string, fields ...zap.Field) {
	GetLogger().Debug(msg, fields...)
}

// Warn logs a warning message
func Warn(msg string, fields ...zap.Field) {
	GetLogger().Warn(msg, fields...)
}

// Fatal logs a fatal message and exits
func Fatal(msg string, fields ...zap.Field) {
	GetLogger().Fatal(msg, fields...)
	os.Exit(1)
}
