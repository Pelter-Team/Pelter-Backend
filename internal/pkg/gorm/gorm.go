package gorm

import (
	"Pelter_backend/internal/config"

	"errors"

	"fmt"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"

	"Pelter_backend/internal/entity"
)

func DbConn(cfg *config.Db) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		cfg.Host,
		cfg.User,
		cfg.Password,
		cfg.DBName,
		cfg.Port,
		cfg.SSLMode,
		cfg.TimeZone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, errors.New("failed to connect database")
	}
	if err := db.Exec("DO $$ BEGIN IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'role_type') THEN CREATE TYPE role_type AS ENUM ('admin', 'customer', 'seller', 'foundation'); END IF; END $$;").Error; err != nil {
		return nil, errors.New("failed to create ENUM before automigrate")
	}

	// prevent errors when creating tables
	err = db.AutoMigrate(&entity.User{})
	if err != nil {
		return nil, errors.New("failed to migrate schema")
	}

	err = db.AutoMigrate(&entity.Product{}, &entity.Review{}, &entity.Transaction{})
	if err != nil {
		return nil, errors.New("failed to migrate schema")
	}

	// err = db.AutoMigrate(&entity.Product{}, &entity.User{}, &entity.Review{}, &entity.Transaction{})
	// if err != nil {
	// 	return nil, errors.New("failed to migrate schema")
	// }

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
