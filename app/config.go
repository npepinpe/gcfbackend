package app

// Mapping for the main application configuration structure in application.yml
type Configuration struct {
	Honeybadger hbConfig     `yaml:"honeybadger"`
	LogLevel    string       `yaml:"logLevel"`
	Server      serverConfig `yaml:"server"`
}

// Mapping for the sub honeybadger part
type hbConfig struct {
	Enabled bool `yaml:"enabled"`
}

// Mapping for server configuration
type serverConfig struct {
	Host string // can be blank
	Port string // can be blank
}

func (config *Configuration) ServerAddress() string {
	if len(config.Server.Port) > 0 {
		return config.Server.Host + ":" + config.Server.Port
	} else {
		return config.Server.Host
	}
}
