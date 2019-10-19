package api

// Config defines the service configuration
type Config struct {
	APIPort   string
	JWTSecret string
}

// LoadConfig populates struct with environment variables or default values
func LoadConfig() *Config {

	return &Config{
		APIPort: "3000",
	}
}
