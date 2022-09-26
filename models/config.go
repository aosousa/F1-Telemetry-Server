package models

import (
	"encoding/json"
	"os"

	"github.com/aosousa/F1-Telemetry-Server/logger"
)

// Config struct contains all the necessary configurations to run the application (e.g., Database config)
type Config struct {
	UDPPort uint16 `json:"udpPort"`  // Port to use in the UDP server
	DB      DB     `json:"database"` // Database configuration
}

// DB struct contains the database configuration
type DB struct {
	Host     string `json:"host"`     // Database hostname or IP address
	Port     string `json:"port"`     // Port in which the database is running
	User     string `json:"user"`     // User used for authentication in the database
	Password string `json:"password"` // Password used for authentication in the database
	Schema   string `json:"schema"`   // Database schema used
}

// CreateConfig adds information from a configuration file to a Config struct
func CreateConfig() Config {
	var config Config
	jsonFile, err := os.ReadFile("./config.json")
	if err != nil {
		logger.Error(err.Error())
	}

	err = json.Unmarshal(jsonFile, &config)
	if err != nil {
		logger.Error(err.Error())
	}

	if (DB{} == config.DB) {
		logger.Error("Database configuration missing")
	}

	return config
}
