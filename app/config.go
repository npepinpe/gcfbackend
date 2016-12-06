package app

import (
	"fmt"
	"net"
)

const (
	ConfigDbPasswordKey     string = "database_password"
	ConfigDbUsernameKey     string = "database_username"
	ConfigHoneybadgerAPIKey string = "honeybadger_api_key"
)

// Mapping for the main application configuration structure in application.yml
type Configuration struct {
	Honeybadger hbConfig     `yaml:"honeybadger"`
	LogLevel    string       `yaml:"logLevel"`
	Server      serverConfig `yaml:"server"`
	Database    dbConfig     `yaml:"database"`
}

// Mapping for the sub honeybadger part
type hbConfig struct {
	Enabled bool `yaml:"enabled"`
}

// Mapping for server configuration
type serverConfig struct {
	Host string `yaml:"host"` // can be blank
	Port string `yaml:"port"` // can be blank
}

// Mapping for database configuration
type dbConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
	Address  string `yaml:"address"`
	Protocol string `yaml:"protocol"`
}

func (config *serverConfig) Address() string {
	return net.JoinHostPort(config.Host, config.Port)
}

func (config *dbConfig) DataSourceName() (dataSourceName string) {
	return fmt.Sprintf("%s:%s@%s(%s)/%s", config.User,
		config.Password, config.Protocol, config.Address, config.Dbname)
}
