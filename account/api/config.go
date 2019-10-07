package api

import "os"

// Config defines the service configuration
type Config struct {
	APIPort           string
	MongoEndpoint     string
	MongoDatabase     string
	GoogleAPIClientID string
	JWTSecret         string
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

	GoogleAPIClientID, found := os.LookupEnv("GOOGLE_API_CLIENT_ID")
	if !found {
		panic("Google api client id is required - GOOGLE_API_CLIENT_ID")
	}

	JWTSecret, found := os.LookupEnv("JWT_SECRET_KEY")
	if !found {
		panic("jwt secret is required - JWT_SECRET_KEY")
	}

	return &Config{
		APIPort:           APIPort,
		MongoEndpoint:     MongoEndpoint,
		MongoDatabase:     MongoDatabase,
		GoogleAPIClientID: GoogleAPIClientID,
		JWTSecret:         JWTSecret,
	}
}
