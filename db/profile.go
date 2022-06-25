package db

import "gorm.io/gorm"

type Profile struct {
	gorm.Model
	UserID      uint `gorm:"not null"`
	Bio         string
	Image       string
	SocialMedia []SocialMedia `gorm:"foreignkey:ProfileID"`
}
