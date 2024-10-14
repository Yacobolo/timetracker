package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/joho/godotenv/autoload"
)

type Service interface {
	Close() error

	GetDB() *pgxpool.Pool
}

type service struct {
	db  *pgxpool.Pool
	dsn string
}

var (
	dbInstance *service
)

func NewService(dsn string) Service {
	if dbInstance != nil {
		return dbInstance
	}

	// Configure connection pool
	config, err := pgxpool.ParseConfig(dsn)
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
		db:  pool,
		dsn: dsn,
	}
	return dbInstance
}

// GetDB provides access to the underlying DB connection pool for SQLC queries.
func (s *service) GetDB() *pgxpool.Pool {
	return s.db
}

// Close closes the database connection pool.
func (s *service) Close() error {
	log.Printf("Closing database connection pool: %s", s.dsn) // Access dsn from the struct
	s.db.Close()
	return nil
}
