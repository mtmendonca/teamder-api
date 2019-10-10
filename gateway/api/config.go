package api

import "os"

// Config defines the service configuration
type Config struct {
	APIPort         string
	AccountEndpoint string
}

// LoadConfig populates struct with environment variables or default values
func LoadConfig() *Config {
	port, found := os.LookupEnv("API_PORT")
	if !found {
		port = "3000"
	}

	accountEndpoint, found := os.LookupEnv("ACCOUNT_SERVICE_ENDPOINT")
	if !found {
		accountEndpoint = "localhost:3001"
	}

	return &Config{
		APIPort:         port,
		AccountEndpoint: accountEndpoint,
	}
}
