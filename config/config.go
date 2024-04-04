package config

import (
	"os"
	"strconv"
)

type Server struct {
	Hostname        string
	Port            int
	ShutdownTimeout int32
}

type FeatureFlag struct{}

type Config struct {
	FeatureFlag  FeatureFlag
	DBConnection string
	Server       Server
}

const (
	keyHostname        = "HOSTNAME"
	keyPort            = "PORT"
	keyShutdownTimeout = "SHUTDOWN_TIMEOUT"

	keyDBConnection = "DB_CONNECTION"
)

const (
	defaultHostname        = ""
	defaultPort            = 8080
	defaultShutdownTimeout = 10

	defaultDBConnection = "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable"
)

func New() Config {
	return Config{
		Server: Server{
			Hostname:        envString(keyHostname, defaultHostname),
			Port:            envInt(keyPort, defaultPort),
			ShutdownTimeout: envInt(keyShutdownTimeout, defaultShutdownTimeout),
		},
		FeatureFlag:  FeatureFlag{},
		DBConnection: envString(keyDBConnection, defaultDBConnection),
	}
}

func envString(key string, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok && value != "" {
		return value
	}
	return defaultValue
}

func envInt(key string, defaultValue int) int {
	if value, ok := os.LookupEnv(key); ok && value != "" {
		intValue, err := strconv.Atoi(value)
		if err != nil {
			return defaultValue
		}
		return intValue
	}
	return defaultValue
}
