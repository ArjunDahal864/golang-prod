package db

import "gorm.io/gorm"

type SocialMedia struct {
	gorm.Model
	ProfileID uint `gorm:"not null"`
	Link      string
	Name      string
}
