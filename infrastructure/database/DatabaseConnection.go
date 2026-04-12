package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return err
	}

	envMode := os.Getenv("ENV")

	host := ""
	port := ""
	user := ""
	password := ""
	dbName := ""

	if envMode == "prod" {
		host = os.Getenv("DB_HOST_PROD")
	} else {
		host = os.Getenv("DB_HOST_DEV")
	}

	if envMode == "prod" {
		port = os.Getenv("DB_PORT_PROD")
	} else {
		port = os.Getenv("DB_PORT_DEV")
	}

	if envMode == "prod" {
		user = os.Getenv("DB_USER_PROD")
	} else {
		user = os.Getenv("DB_USER_DEV")
	}

	if envMode == "prod" {
		password = os.Getenv("DB_PASSWORD_PROD")
	} else {
		password = os.Getenv("DB_PASSWORD_DEV")
	}

	if envMode == "prod" {
		dbName = os.Getenv("DB_NAME_PROD")
	} else {
		dbName = os.Getenv("DB_NAME_DEV")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbName, port)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	// var tableErr error
	// tableErr = DB.AutoMigrate(&model.Player{})
	// if tableErr != nil {
	// 	return tableErr
	// }
	return nil
}
