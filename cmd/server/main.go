package main

import (
	"fmt"
	"log"

	"github.com/lutfiharidha/google-trace/internal/config"
	"github.com/lutfiharidha/google-trace/pkg/infrasturcture/server"
	"github.com/lutfiharidha/google-trace/pkg/shared/logger"
)

func main() {
	config := config.GetConfig()

	fmt.Println(config.Logger)

	logConfig := logger.Configuration{
		EnableConsole:     config.Logger.Console.Enable,
		ConsoleJSONFormat: config.Logger.Console.JSON,
		ConsoleLevel:      config.Logger.Console.Level,
		EnableFile:        config.Logger.File.Enable,
		FileJSONFormat:    config.Logger.File.JSON,
		FileLevel:         config.Logger.File.Level,
		FileLocation:      config.Logger.File.Path,
	}

	if err := logger.NewLogger(logConfig, logger.InstanceZapLogger); err != nil {
		log.Fatalf("Could not instantiate log %v", err)
	}

	logConfigPanic := logger.Configuration{
		EnableConsole:     config.Logger.ConsolePanic.Enable,
		ConsoleJSONFormat: config.Logger.ConsolePanic.JSON,
		ConsoleLevel:      config.Logger.ConsolePanic.Level,
		EnableFile:        config.Logger.FilePanic.Enable,
		FileJSONFormat:    config.Logger.FilePanic.JSON,
		FileLevel:         config.Logger.FilePanic.Level,
		FileLocation:      config.Logger.FilePanic.Path,
	}

	if err := logger.NewLoggerPanic(logConfigPanic, logger.InstanceZapLogger); err != nil {
		log.Fatalf("Could not instantiate log %v", err)
	}

	server.RunServer()
}
