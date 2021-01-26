package logger

import (
	"os"
	"sync"

	"github.com/rs/zerolog"
)

// Logger ...
type Logger struct {
	*zerolog.Logger
}

var (
	logger Logger
	once   sync.Once
)

// Get creates Logger
func Get(logLevel string) *Logger {
	once.Do(func() {
		zeroLogger := zerolog.New(os.Stderr).With().Timestamp().Logger()
		switch logLevel {
		case "debug":
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
		case "info":
			zerolog.SetGlobalLevel(zerolog.InfoLevel)
		case "warning":
			zerolog.SetGlobalLevel(zerolog.WarnLevel)
		case "error":
			zerolog.SetGlobalLevel(zerolog.ErrorLevel)
		case "fatal":
			zerolog.SetGlobalLevel(zerolog.FatalLevel)
		default:
			zerolog.SetGlobalLevel(zerolog.InfoLevel)
		}
		logger = Logger{&zeroLogger}
	})
	return &logger
}
