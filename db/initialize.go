package db

import (
	"fmt"
	"prod/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(c config.Config) *gorm.DB {
	ofig := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", c.DBPassword, c.DBUser, c.DBHost, c.DBPort, c.DBName)
	db, err := gorm.Open(postgres.Open(ofig), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func Migrate(db *gorm.DB) {
	if err := db.AutoMigrate(&User{}, &Profile{}, &SocialMedia{}); err != nil {
		panic("failed to migrate database")
	}
}
