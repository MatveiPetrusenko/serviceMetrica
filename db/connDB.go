package db

import (
	"database/sql"
	"fmt"
	"serviceMetrica/internal/config"
	"sync"

	_ "github.com/lib/pq"
)

var (
	instanceConnection *sql.DB
	onceConnection     sync.Once
)

func New() *sql.DB {
	onceConnection.Do(func() {
		instanceConnection, _ = ConnectToDB()
	})

	return instanceConnection
}

// ConnectToDB connect to database
func ConnectToDB() (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s host=%s port=%s",
		config.New().PostgreSQL.User,
		config.New().PostgreSQL.Password,
		config.New().PostgreSQL.DbName,
		config.New().PostgreSQL.SSLMode,
		config.New().PostgreSQL.Host,
		config.New().PostgreSQL.Port)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("Failed to connect database: %v\n", err)
	}

	err = pingDatabase(db)
	if err != nil {
		return nil, fmt.Errorf("Ping method did not executed %v\n", err)
	}

	return db, nil
}

// pingDatabase ping connection
func pingDatabase(db *sql.DB) error {
	err := db.Ping()
	if err != nil {
		return fmt.Errorf("Unable to ping connection: %v\n", err)
	}

	return nil
}
