package logger

import (
	"fmt"
	"sync"

	"go.uber.org/zap"
)

var once sync.Once

var logger *zap.SugaredLogger

type Config struct {
	Environment string
}

// FromConfig returns sugared logger provided with a configuration.
func FromConfig(config *Config) (*zap.SugaredLogger, error) {
	var err error

	// Create and set package scoped logger
	once.Do(func() {
		var l *zap.Logger

		switch config.Environment {
		case "dev":
			l, err = zap.NewDevelopment()

		case "prod":
			l, err = zap.NewProduction()
		default:
			err = fmt.Errorf("no such environment %q", config.Environment)
		}

		logger = l.Sugar()
	})

	// Error Handling
	if err != nil {
		return nil, err
	}

	return logger, nil
}
