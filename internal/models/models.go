package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID                int    `gorm:"primaryKey; uniqueIndex"`
	Username          string `gorm:"uniqueIndex; not null"`
	UserPicture       string
	UserEmail         string
	Verified          bool
	VerificationToken string
	RestoreSecureKey  string
	Videos            []Video `gorm:"foreignKey:UserID"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	Deleted           gorm.DeletedAt `gorm:"index"`
}

type Video struct {
	ID             uint `gorm:"primaryKey; autoIncrement"`
	UserID         int
	User           User   `gorm:"foreignKey:UserID"`
	Title          string `gorm:"not null"`
	FilePath       string `gorm:"not null"`
	MongoDBOID     string
	Converted      bool
	Cut            bool
	Watermarked    bool
	Screenshot     bool
	ScreenshotPath string
	CreatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}
