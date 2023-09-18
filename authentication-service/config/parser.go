package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

// Parse parses all configuration to a single config object
func Parse() *Config {
	return &Config{
		AppConfig: parseAppConfig(),
		DBConfig:  parseDBConfigs(),
	}
}

// parse application configurations
func parseAppConfig() *AppConfig {
	configFile := CheckEnvironment("app")
	content := ReadFile(configFile)

	cfg := &AppConfig{}

	err := yaml.Unmarshal(*content, &cfg)
	if err != nil {
		panic(fmt.Sprintf("error: %v", err))
	}

	return cfg
}

// parse db configurations
func parseDBConfigs() *DBConfig {
	configFile := CheckEnvironment("db")
	content := ReadFile(configFile)

	cfg := &DBConfig{}

	err := yaml.Unmarshal(*content, &cfg)
	if err != nil {
		panic(fmt.Sprintf("error: %v", err))
	}

	return cfg
}

func CheckEnvironment(folderName string) string {
	env := os.Getenv("ENV")

	switch env {
	case "staging":
		return folderName + "/config-staging.yaml"
	case "production":
		return folderName + "/config-production.yaml"
	default:
		return folderName + "/config-dev.yaml"
	}
}
