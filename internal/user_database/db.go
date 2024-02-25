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
		log.Fatalf("Error loading the env variablesz: %v", errorLoadingenv)
	}

	password := os.Getenv("POSTGRES_PASSWORD")
	database := os.Getenv("DB_NAME")
	host := os.Getenv("POSTGRES_HOST")
	username := os.Getenv("POSTGRES_USERNAME")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai", host, username, password, database)
	var err error
	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	if err = Database.AutoMigrate(&models.User{}, &models.Video{}); err != nil {
		log.Fatalf("Error performing migration: %v", err)
	}
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

func DeletedUserWithID(db *gorm.DB, UserID int) (bool, error) {
	var user models.User

	result := db.Unscoped().First(&user, UserID)
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

func GetUserWithID(db *gorm.DB, UserID int) (*models.User, error) {
	var user models.User

	result := db.First(&user, UserID)
	return &user, result.Error
}

func GetDeletedUserWithID(db *gorm.DB, UserID int) (*models.User, error) {
	var user models.User

	result := db.Unscoped().Where("id=?", UserID).First(&user)
	return &user, result.Error
}

func GetLatestVideo(db *gorm.DB, userId int) (*models.Video, error) {
	var Video models.Video
	if result := db.Where("user_id=?", userId).Order("created_at desc").First(&Video).Error; result != nil {
		if result == gorm.ErrRecordNotFound {
			return nil, result
		}
		return nil, nil
	}
	return &Video, nil
}

func GetAllVideo(db *gorm.DB, userId int) ([]models.Video, error) {
	var Video []models.Video
	if result := db.Where("user_id=?", userId).Order("created_at desc").Find(&Video).Error; result != nil {
		if result == gorm.ErrRecordNotFound {
			return nil, result
		}
		return nil, nil
	}
	return Video, nil
}
func GetVideoByID(db *gorm.DB, videoID string) (*models.Video, error) {
	var video models.Video
	if err := db.Where("video_key = ?", videoID).First(&video).Error; err != nil {
		return nil, err
	}
	return &video, nil
}
