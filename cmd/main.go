package main

import (
	"github.com/RexGreenway/CollectionApp/internal/logger"
)

func main() {
	// Establish Logging
	logger, err := logger.FromConfig(&logger.Config{Environment: "dev"})
	if err != nil {
		panic(err)
	}

	logger.Info("starting collection application")

	// - Create Application
	// - Start gRPC Server

	logger.Info("exiting collection application")
}
