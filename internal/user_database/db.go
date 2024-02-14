package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Define the connection string
	dsn := "host=localhost port=5432 user=rituparn password=rituparn28 dbname=video_conversion sslmode=disable"

	// Open a connection to the database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error opening connection:", err)
		return
	}

	// Check if the database exists
	var count int64
	db.Raw("SELECT COUNT(*) FROM pg_database WHERE datname = ?", "video_conversion").Scan(&count)
	if count == 0 {
		fmt.Println("Database 'my_database' does not exist.")
	} else {
		fmt.Println("Database 'my_database' exists.")
	}
}
