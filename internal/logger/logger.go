package logger

import (
	"fmt"

	"go.uber.org/zap"
)

type Config struct {
	Environment string
}

// FromConfig returns sugared logger provided with a configuration.
func FromConfig(config *Config) (*zap.SugaredLogger, error) {
	// Establish
	var logger *zap.Logger
	var err error

	switch config.Environment {
	case "dev":
		logger, err = zap.NewDevelopment()
		if err != nil {
			return nil, err
		}
	case "prod":
		logger, err = zap.NewProduction()
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("no such environment %q", config.Environment)
	}

	return logger.Sugar(), nil
}
