package db

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string  `grom:"type:varchar(255);not null"`
	LastName  string  `grom:"type:varchar(255);not null"`
	Email     string  `grom:"type:varchar(255);not null;unique"`
	Password  string  `grom:"type:varchar(255);not null"`
	Profile   Profile `gorm:"foreignkey:UserID"`
}
