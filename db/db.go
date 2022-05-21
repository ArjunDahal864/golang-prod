package db

import (
	"database/sql"
	"fmt"
	"prod/conf"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/lib/pq"
)

type Database interface {
	PostgresDB() *sql.DB
	// add more database methods if needed
	MockDB() *sql.DB
}

type PostgresDatabase struct {
	DB conf.SQLDatabase
}

type MockDB struct{}

func (s *PostgresDatabase) Connect() *sql.DB {
	db, err := sql.Open(s.DB.Driver, fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", s.DB.User, s.DB.Password, s.DB.Host, s.DB.Port, s.DB.Name))
	if err != nil {
		panic(err)
	}
	return db
}

func (s *MockDB) Connect() *sql.DB {
	db, _, err := sqlmock.New()
	if err != nil {
		panic(err)
	}
	return db
}

func Postgres(d Database) *sql.DB {
	return d.PostgresDB()
}
