package config

// config holds all other config structs
type Config struct {
	AppConfig *AppConfig
	DBConfig  *DBConfig
}

// app config holds application configurations
type AppConfig struct {
	Name                string `yaml:"name"`
	Host                string `yaml:"host"`
	Port                int    `yaml:"port"`
	Env                 string `yaml:"env"`
	ReadTimeout         uint   `yaml:"read-timeout"`
	WriteTimeout        uint   `yaml:"write-timeout"`
	ShutdownWaitTimeout uint   `yaml:"shutdown-wait-timeout"`
}

// db config holds database configurations
type DBConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}
