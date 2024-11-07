package config

// Config struct for structuring the config data, can be extended accordingly
type Config struct {
	Database struct {
		URL     string `yaml:"url"`
		DB      string `yaml:"db"`
		Timeout int    `yaml:"timeout"`
	} `yaml:"database"`
	Server struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`
}

// NewConfig returns a hardcoded configuration for simplicity
func NewConfig() *Config {
	config := Config{}
	config.Database.URL = "mongodb://admin:admin@mongo:27017/?retryWrites=true&w=majority&authSource=admin"
	config.Database.DB = "catalog"
	config.Database.Timeout = 5

	config.Server.Port = "8080"
	config.Server.Host = "localhost"
	return &config
}
