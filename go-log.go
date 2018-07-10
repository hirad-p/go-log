package main

import (
	"fmt"
	"log"

	"github.com/jb-hirad/go-log/config"
	"github.com/jb-hirad/go-log/logger"
	"github.com/jb-hirad/go-log/model"
)

// Create returns a logger that can be used to log
func Create(configuration model.Config) *logger.Logger {
	if configuration.Config == "" {
		log.Fatal("Must supply a configuration file. Check the example file if needed")
	}

	if configuration.Log == "" {
		fmt.Println("Log path not supplied, logs will not be persisted")
	}

	logger := logger.Logger{
		Levels: config.Read(configuration.Config),
		Path:   configuration.Log,
	}
	return &logger
}

func main() {
	config := model.Config{
		Config: "config/config.json",
	}
	goLogger := Create(config)

	goLogger.Log(0, "Hello There")
}
