package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/joho/godotenv/autoload"
)

// Service represents a service that manages a database connection.
type Service interface {
	// Close terminates the database connection.
	Close() error

	// GetDB returns the database connection pool to be used by SQLC queries.
	GetDB() *pgxpool.Pool
}

type service struct {
	db *pgxpool.Pool
}

var (
	dburl      = os.Getenv("DB_URL")
	dbInstance *service
)

// NewService initializes the database connection pool and returns a Service interface.
func NewService() Service {
	if dbInstance != nil {
		return dbInstance
	}

	// Configure connection pool
	config, err := pgxpool.ParseConfig(dburl)
	if err != nil {
		log.Fatalf("Unable to parse connection string: %v", err)
	}

	config.MaxConns = 10

	// Create the connection pool
	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v", err)
	}

	dbInstance = &service{
		db: pool,
	}
	return dbInstance
}

// GetDB provides access to the underlying DB connection pool for SQLC queries.
func (s *service) GetDB() *pgxpool.Pool {
	return s.db
}

// Close closes the database connection pool.
func (s *service) Close() error {
	log.Printf("Closing database connection pool: %s", dburl)
	s.db.Close()
	return nil
}
