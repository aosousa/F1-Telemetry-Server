package main

import (
	"fmt"

	h "github.com/aosousa/F1-Telemetry-Server/handlers"
	"github.com/aosousa/F1-Telemetry-Server/logger"
)

const version = "0.0.1"

func init() {
	// set up Config struct before setting up the server
	fmt.Println("Configuration file: Loading")
	h.InitConfig()
	fmt.Printf("Configuration file: OK\n\n")

	// initialize database through Config struct
	fmt.Println("System database: Checking")
	h.InitDatabase()
	fmt.Printf("System database: OK\n\n")

	fmt.Printf("F1 Telemetry Server V%s running\n", version)
}

func main() {
	logger.Debug("Test")
}
