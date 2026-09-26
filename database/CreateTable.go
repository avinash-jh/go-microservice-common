package database

import (
	"fmt"

	"github.com/avinash-jh/go-microservice-common/models"

	"gorm.io/gorm"
)

func CreateTable(db *gorm.DB) error {
	// AutoMigrate creates the table if it does not exist
	err := db.AutoMigrate(&models.Product{})
	if err != nil {
		return fmt.Errorf("auto migration failed: %w", err)
	}
	fmt.Println("Product table created successfully")

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return fmt.Errorf("auto migration failed: %w", err)
	}
	fmt.Println("user table created successfully")
	return nil
}
