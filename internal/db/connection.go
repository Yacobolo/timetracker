package db

import (
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
)

// Service represents a service that manages a database connection.
type Service interface {
	// Close terminates the database connection.
	Close() error

	// GetDB returns the database connection to be used by SQLC queries.
	GetDB() *sqlx.DB
}

type service struct {
	db *sqlx.DB
}

var (
	dburl      = os.Getenv("DB_URL")
	dbInstance *service
)

// NewService initializes the database connection and returns a Service interface.
func NewService() Service {
	if dbInstance != nil {
		return dbInstance
	}

	db := sqlx.MustOpen("sqlite3", dburl)

	// Optional: Set connection pooling parameters
	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	dbInstance = &service{
		db: db,
	}
	return dbInstance
}

// GetDB provides access to the underlying DB object for SQLC queries.
func (s *service) GetDB() *sqlx.DB {
	return s.db
}

// Close closes the database connection.
func (s *service) Close() error {
	log.Printf("Closing database connection: %s", dburl)
	return s.db.Close()
}
