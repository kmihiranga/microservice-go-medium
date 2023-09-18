package log

import (
	"sync"

	"authentication-service/config"

	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

var logger *zap.Logger
var once sync.Once

func GetLogger() *zap.Logger {
	once.Do(func() {
		// read logger configurations using 
		configFile := config.CheckEnvironment("log")
		content := config.ReadFile(configFile)

		var cfg zap.Config
		err := yaml.Unmarshal(*content, &cfg)
		if err != nil {
			panic(err)
		}

		logger, err = cfg.Build()
		if err != nil {
			panic(err)
		}
		defer logger.Sync()
		logger.Info("Logger Initialized.")
	})
	return logger
}