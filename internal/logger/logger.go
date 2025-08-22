// logger/logger.go
package logger

import (
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var log *zap.Logger

// Config holds the logger configuration
type Config struct {
	Level      string
	FilePath   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
}

// InitLogger initializes the logger with the given configuration
func InitLogger(cfg *Config) error {
	if cfg == nil {
		cfg = &Config{
			Level: "info",
		}
	}

	// Set default values
	if cfg.MaxSize == 0 {
		cfg.MaxSize = 100 // MB
	}
	if cfg.MaxBackups == 0 {
		cfg.MaxBackups = 3
	}
	if cfg.MaxAge == 0 {
		cfg.MaxAge = 28 // days
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	var writer zapcore.WriteSyncer
	if cfg.FilePath != "" {
		// Buat direktori jika belum ada
		if err := os.MkdirAll(getDir(cfg.FilePath), 0755); err != nil {
			return err
		}

		lumberJackLogger := &lumberjack.Logger{
			Filename:   cfg.FilePath,
			MaxSize:    cfg.MaxSize,
			MaxBackups: cfg.MaxBackups,
			MaxAge:     cfg.MaxAge,
			Compress:   cfg.Compress,
		}
		writer = zapcore.AddSync(lumberJackLogger)
	} else {
		writer = zapcore.AddSync(os.Stdout)
	}

	var level zapcore.Level
	if err := level.UnmarshalText([]byte(cfg.Level)); err != nil {
		level = zapcore.InfoLevel
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		writer,
		level,
	)

	log = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	return nil
}

// Helper: ambil direktori dari file path
func getDir(path string) string {
	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == '/' || path[i] == '\\' {
			return path[:i]
		}
	}
	return "."
}

// Fungsi logging
func Debug(msg string, args ...interface{}) {
	if len(args) > 0 {
		msg = fmt.Sprintf(msg, args...)
	}
	log.Debug(msg)
}

func Info(msg string, args ...interface{}) {
	if len(args) > 0 {
		msg = fmt.Sprintf(msg, args...)
	}
	log.Info(msg)
}

func Infof(format string, args ...interface{}) {
	log.Info(fmt.Sprintf(format, args...))
}

func Warn(msg string, args ...interface{}) {
	if len(args) > 0 {
		msg = fmt.Sprintf(msg, args...)
	}
	log.Warn(msg)
}

func Error(msg string, err error) {
	if err != nil {
		log.Error(msg, zap.Error(err))
	} else {
		log.Error(msg)
	}
}

func Errorf(format string, args ...interface{}) {
	log.Error(fmt.Sprintf(format, args...))
}

func Fatal(msg string, err error) {
	if err != nil {
		log.Fatal(msg, zap.Error(err))
	} else {
		log.Fatal(msg)
	}
}

func Fatalf(format string, args ...interface{}) {
	log.Fatal(fmt.Sprintf(format, args...))
}

func WithFields(fields map[string]interface{}) *zap.Logger {
	zapFields := make([]zap.Field, 0, len(fields))
	for k, v := range fields {
		zapFields = append(zapFields, zap.Any(k, v))
	}
	return log.With(zapFields...)
}

func WithError(err error) *zap.Logger {
	return log.With(zap.Error(err))
}

func WithTime(t time.Time) *zap.Logger {
	return log.With(zap.Time("time", t))
}
