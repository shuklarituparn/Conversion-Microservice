package main

import (
	"fmt"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100})
	fmt.Println("Product created")

	// Read
	var product Product
	db.First(&product, 1) // find product with integer primary key
	fmt.Printf("Product found by ID: %+v\n", product)

	db.First(&product, "code = ?", "D42") // find product with code D42
	fmt.Printf("Product found by code: %+v\n", product)

	// Update - update product's price to 200
	db.Model(&product).Update("Price", 200)
	fmt.Println("Product price updated to 200")

	// Update - update multiple fields
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	fmt.Println("Product fields updated")

	// Read updated product
	db.First(&product, 1)
	fmt.Printf("Updated product: %+v\n", product)

	// Delete - delete product
	db.Delete(&product, 1)
	fmt.Println("Product deleted")
}
