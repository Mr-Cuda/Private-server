package database

import (
	"fmt"
	"os"
	"strconv"
)

// DB gorm connector
var DB *gorm.DB

// connection with database
func ConnectDB() {
	var err error
	p := os.Getenv("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	DB, err = gorm.Open("postgres", fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		port, os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME")))

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")

	DB.AutoMigrate(&models.Employee{})
	fmt.Println("Database Migrated")
}
