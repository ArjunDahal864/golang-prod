package db

import (
	"fmt"
	"golang-prod/conf"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database interface {
	ConnectToSqlDatabase() *gorm.DB
	ConnectToPostgresDatabase() *gorm.DB
}

func New(config conf.Config) *Connection {
	return &Connection{
		config: config,
	}
}

type Connection struct {
	config conf.Config
}

func (conf Connection) ConnectToSQliteDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("sql.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func (conf Connection) ConnectToPostgresDatabase() *gorm.DB {
	database := conf.config.Database
	dburl := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", database.User, database.Password, database.Host, database.Port, database.Name)
	db, err := gorm.Open(postgres.Open(dburl), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
