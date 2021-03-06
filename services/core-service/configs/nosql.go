package config

type NoSQLConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBType   string `yaml:"dbtype"`
}

type NoSQLConfigs []NoSQLConfig
