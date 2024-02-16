package user_database

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/shuklarituparn/Conversion-Microservice/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var Database *gorm.DB

func init() {
	errorLoadingenv := godotenv.Load("../../.env")
	if errorLoadingenv != nil {
		log.Fatalf("Error loading the env variables")
	}

	password := os.Getenv("POSTGRES_PASSWORD")
	database := os.Getenv("DB_NAME")
	host := os.Getenv("POSTGRES_HOST")
	username := os.Getenv("POSTGRES_USERNAME")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai", host, username, password, database)
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err := db.AutoMigrate(&models.User{}, &models.Video{}); err != nil {
		log.Fatalf("Error performing migration: %v", err)
	}

	Database = db

}

func ReturnDbInstance() *gorm.DB {
	return Database
}

func UserWithID(db *gorm.DB, UserID int) (bool, error) {
	var user models.User

	result := db.First(&user, UserID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, result.Error
	}
	return true, nil
}

func EmailExits(db *gorm.DB, UserId int) (bool, error) {
	var user models.User

	result := db.Select("id, email").Where("id=? AND email<>''", UserId).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, result.Error
	}
	return true, nil

}

func IsVerified(db *gorm.DB, UserId int) (bool, error) {
	var user models.User
	result := db.Select("verified").Where("id = ?", UserId).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, result.Error
	}
	return user.Verified, nil

}

func GetUserWithID(db *gorm.DB, UserID int) (models.User, error) {
	var user models.User

	result := db.First(&user, UserID)
	return user, result.Error
}
