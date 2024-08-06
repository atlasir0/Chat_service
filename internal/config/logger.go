package config

import (
	"errors"
	"os"
)

type LoggerConfig interface {
	LoggerLevel() string
}

var _ LoggerConfig = (*loggerConfig)(nil)

const (
	loggerLevel = "LOGGER_LEVEL"
)

type loggerConfig struct {
	loggerLevel string
}

// NewLoggerConfig - ...
func NewLoggerConfig() (*loggerConfig, error) {
	loggerLevel := os.Getenv(loggerLevel)
	if len(loggerLevel) == 0 {
		return nil, errors.New("loggerLevel not found")
	}

	return &loggerConfig{
		loggerLevel: loggerLevel,
	}, nil
}

func (cfg *loggerConfig) LoggerLevel() string {
	return cfg.loggerLevel
}
