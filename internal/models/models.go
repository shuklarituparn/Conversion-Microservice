package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID          uint64 `gorm:"primaryKey; uniqueIndex"`
	Username    string `gorm:"uniqueIndex; not null"`
	UserPicture string `gorm:"uniqueIndex; not null"`
	UserEmail   string
	Videos      []Video `gorm:"foreignKey:UserID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Deleted     gorm.DeletedAt `gorm:"index"`
}

type Video struct {
	ID         uint64 `gorm:"primaryKey"`
	UserID     uint64
	Title      string `gorm:"not null"`
	FilePath   string `gorm:"not null"`
	MongoDBOID string
	CreatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
