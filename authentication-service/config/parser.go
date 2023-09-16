package config

import (
	"fmt"
	"os"
	"path/filepath"

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
	content := readFile(configFile)

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
	content := readFile(configFile)

	cfg := &DBConfig{}

	err := yaml.Unmarshal(*content, &cfg)
	if err != nil {
		panic(fmt.Sprintf("error: %v", err))
	}

	return cfg
}

// read file from a disk
func readFile(fileName string) *[]byte {
	// identify current file directory
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	content, err := os.ReadFile(dir + "/ops/" + fileName)
	if err != nil {
		panic(fmt.Sprintf("error: %v", err))
	}

	return &content
}

func CheckEnvironment(folderName string) string {
	env := os.Getenv("ENV")

	switch env {
	case "development":
		return folderName + "/config-dev.yaml"
	case "staging":
		return folderName + "/config-staging.yaml"
	case "production":
		return folderName + "/config-production.yaml"
	default:
		return ""
	}
}
