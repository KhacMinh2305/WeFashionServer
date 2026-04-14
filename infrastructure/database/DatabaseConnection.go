package database

import (
	"WeFashionServer/infrastructure/model"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func initTable[T interface{}](table *T, err *error) {
	var tableErr error
	if tableErr = DB.AutoMigrate(table); tableErr != nil {
		*err = tableErr
	}
}

func initTables() error {
	var tableErr error
	initTable(&model.Account{}, &tableErr)
	initTable(&model.Address{}, &tableErr)
	initTable(&model.Category{}, &tableErr)
	initTable(&model.Color{}, &tableErr)
	initTable(&model.Coupon{}, &tableErr)
	initTable(&model.FavoriteProduct{}, &tableErr)
	initTable(&model.FavoriteSku{}, &tableErr)
	initTable(&model.Order{}, &tableErr)
	initTable(&model.OrderProductVariant{}, &tableErr)
	initTable(&model.Payment{}, &tableErr)
	initTable(&model.Product{}, &tableErr)
	initTable(&model.ProductVariant{}, &tableErr)
	initTable(&model.Shipper{}, &tableErr)
	initTable(&model.ShippingState{}, &tableErr)
	initTable(&model.Shop{}, &tableErr)
	initTable(&model.Size{}, &tableErr)
	initTable(&model.User{}, &tableErr)
	return tableErr
}

func Connect() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return err
	}

	envMode := os.Getenv("ENV")

	var (
		host     string
		port     string
		user     string
		password string
		dbName   string
	)

	if envMode == "prod" {
		host = os.Getenv("DB_HOST_PROD")
		port = os.Getenv("DB_PORT_PROD")
		user = os.Getenv("DB_USER_PROD")
		password = os.Getenv("DB_PASSWORD_PROD")
		dbName = os.Getenv("DB_NAME_PROD")
	} else {
		host = os.Getenv("DB_HOST_DEV")
		port = os.Getenv("DB_PORT_DEV")
		user = os.Getenv("DB_USER_DEV")
		password = os.Getenv("DB_PASSWORD_DEV")
		dbName = os.Getenv("DB_NAME_DEV")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbName, port)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	return initTables()
}
