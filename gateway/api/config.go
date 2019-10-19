package api

import "os"

// Config defines the service configuration
type Config struct {
	APIPort       string
	EventEndpoint string
	JWTSecret     string
}

// LoadConfig populates struct with environment variables or default values
func LoadConfig() *Config {
	APIPort, found := os.LookupEnv("API_PORT")
	if !found {
		panic("api port required - API_PORT")
	}

	EventEndpoint, found := os.LookupEnv("EVENT_SERVICE_ENDPOINT")
	if !found {
		panic("Event service endpoint required - EVENT_SERVICE_ENDPOINT")
	}

	return &Config{
		APIPort:       APIPort,
		EventEndpoint: EventEndpoint,
	}
}
