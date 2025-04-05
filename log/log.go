package log

import (
	"os"

	"github.com/charmbracelet/log"
)

type Logger struct {
	logger *log.Logger
}

func Initialize() Logger {
	return Logger{
		logger: log.NewWithOptions(os.Stderr, log.Options{}),
	}
}

// TODO: calling this should write to log file and notify the user
func (l Logger) Close() error {
	return nil
}

// TODO: write to log file whenever any part of the program calls the Write method
func (l Logger) Write(level log.Level, message string) error {
	return nil
}
