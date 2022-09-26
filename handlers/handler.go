// Package handlers separates the application's basic functionality into different files (handlers)
//
// This main handler is responsible for the initial configuration of the application (e.g., loading database information into a reusable struct)
package handlers

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/aosousa/F1-Telemetry-Server/logger"
	"github.com/aosousa/F1-Telemetry-Server/models"
)

const ()

var (
	config models.Config
	db     *sql.DB
)

// InitConfig adds information from a configuration file to a Config struct
// that will be used throughout the application
func InitConfig() {
	config = models.CreateConfig()
}

// InitDatabase establishes a connection to the database with parameters in the
// configuration struct that was previously created.
func InitDatabase() {
	var err error
	sqlStmt := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DB.User, config.DB.Password, config.DB.Host, config.DB.Port, config.DB.Schema)

	// establish connection to the database
	db, err = sql.Open("mysql", sqlStmt)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	db.SetConnMaxLifetime(time.Minute * 30)
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(5)

	// https://github.com/go-sql-driver/mysql/wiki/Examples#a-word-on-sqlopen
	// because sql.Open does not actually open a connection and does not return an error
	// we need to call db.Ping in order to check if the connection was established correctly
	// and handle any errors here
	err = db.Ping()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
