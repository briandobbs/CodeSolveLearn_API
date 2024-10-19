package db

import (
	"CodeSolveLearn_API/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

// InitDB initializes the database and runs migrations
func InitDB(user, password, dbname, host string, port int) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Run migrations for Author and Article models
	err = db.AutoMigrate(&models.Author{}, &models.Article{})
	if err != nil {
		return nil, err
	}

	log.Println("Database connection established and models migrated")
	return db, nil
}
