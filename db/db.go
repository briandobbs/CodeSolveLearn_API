package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
)

// Database interface defines the common methods for interacting with the database
type Database interface {
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}

// SQLDB struct implements the Database interface
type SQLDB struct {
	*sql.DB
}

// InitDB initializes a database connection and returns an SQLDB struct that satisfies the Database interface
func InitDB(user, password, dbname, host string, port int) (*SQLDB, error) {
	// Data Source Name (DSN) format: user:password@tcp(host:port)/dbname
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbname)

	// Open a database connection
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Verify the connection
	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Database connection established")
	return &SQLDB{db}, nil
}

// QueryContext satisfies the Database interface for querying the database
func (s *SQLDB) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return s.DB.QueryContext(ctx, query, args...)
}

// ExecContext satisfies the Database interface for executing queries that don't return rows
func (s *SQLDB) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return s.DB.ExecContext(ctx, query, args...)
}
