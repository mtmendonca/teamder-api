package api

import "os"

// Config defines the service configuration
type Config struct {
	APIPort       string
	MongoEndpoint string
	MongoDatabase string
}

// LoadConfig populates struct with environment variables or default values
func LoadConfig() *Config {
	port, found := os.LookupEnv("API_PORT")
	if !found {
		port = "3001"
	}

	mongoEndpoint, found := os.LookupEnv("MONGO_ENDPOINT")
	if !found {
		mongoEndpoint = "mongodb://localhost:27017"
	}

	mongoDatabase, found := os.LookupEnv("MONGO_DATABASE")
	if !found {
		mongoDatabase = "teamder"
	}

	return &Config{
		APIPort:       port,
		MongoEndpoint: mongoEndpoint,
		MongoDatabase: mongoDatabase,
	}
}
