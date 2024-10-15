package gorm

import (
	"Pelter_backend/internal/config"
	"errors"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string `gorm:"primaryKey"`
	Price uint
}

func DbConn(cfg *config.Db) (*gorm.DB, error) {
	// NOTE: check the docs how to pass in db url cause i forgot
	db, err := gorm.Open(postgres.Open("test.db"), &gorm.Config{})
	if err != nil {
		return nil, errors.New("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	// db.Create(&Product{Code: "D42", Price: 100})

	// Read
	// var product Product
	// db.First(&product, 1)                 // find product with integer primary key
	// db.First(&product, "code = ?", "D42") // find product with code D42

	// // Update - update product's price to 200
	// db.Model(&product).Update("Price", 200)
	// // Update - update multiple fields
	// db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	// db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// // Delete - delete product
	// db.Delete(&product, 1)
	return db, nil
}
