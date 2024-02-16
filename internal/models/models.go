package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID                int    `gorm:"primaryKey; uniqueIndex"`
	Username          string `gorm:"uniqueIndex; not null"`
	UserPicture       string `gorm:"uniqueIndex; not null"`
	UserEmail         string
	Verified          bool
	VerificationToken string
	Videos            []Video `gorm:"foreignKey:UserID"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	Deleted           gorm.DeletedAt `gorm:"index"`
}

type Video struct {
	ID         int `gorm:"primaryKey"`
	UserID     int
	Title      string `gorm:"not null"`
	FilePath   string `gorm:"not null"`
	MongoDBOID string
	CreatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
