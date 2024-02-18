package main

import (
	"github.com/shuklarituparn/Conversion-Microservice/internal/models"
	"github.com/shuklarituparn/Conversion-Microservice/internal/user_database"
	"gorm.io/gorm"
)

func GetAllUsers(db *gorm.DB) ([]models.User, error) {
	var users []models.User
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func main() {
	db := user_database.ReturnDbInstance()
	GetAllUsers(db)
}
