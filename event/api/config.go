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
	APIPort, found := os.LookupEnv("API_PORT")
	if !found {
		panic("Api port required - API_PORT")
	}

	MongoEndpoint, found := os.LookupEnv("MONGO_ENDPOINT")
	if !found {
		panic("Mongo endpoint is required - MONGO_ENDPOINT")
	}

	MongoDatabase, found := os.LookupEnv("MONGO_DATABASE")
	if !found {
		panic("Mongo database is required - MONGO_DATABASE")
	}

	return &Config{
		APIPort:       APIPort,
		MongoEndpoint: MongoEndpoint,
		MongoDatabase: MongoDatabase,
	}
}
