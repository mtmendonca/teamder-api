package api

import "os"

// Config defines the service configuration
type Config struct {
	APIPort         string
	AccountEndpoint string
	EventEndpoint   string
	JWTSecret       string
}

// LoadConfig populates struct with environment variables or default values
func LoadConfig() *Config {
	APIPort, found := os.LookupEnv("API_PORT")
	if !found {
		panic("api port required - API_PORT")
	}

	AccountEndpoint, found := os.LookupEnv("ACCOUNT_SERVICE_ENDPOINT")
	if !found {
		panic("Account service endpoint required - ACCOUNT_SERVICE_ENDPOINT")
	}

	EventEndpoint, found := os.LookupEnv("EVENT_SERVICE_ENDPOINT")
	if !found {
		panic("Event service endpoint required - EVENT_SERVICE_ENDPOINT")
	}

	JWTSecret, found := os.LookupEnv("JWT_SECRET_KEY")
	if !found {
		panic("jwt secret key required - JWT_SECRET_KEY")
	}

	return &Config{
		APIPort:         APIPort,
		AccountEndpoint: AccountEndpoint,
		EventEndpoint:   EventEndpoint,
		JWTSecret:       JWTSecret,
	}
}
