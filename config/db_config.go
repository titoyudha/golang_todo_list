package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Cant Load .env files")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=true&loc=Local", dbUser, dbPass, dbHost, dbName)
	db, errorDB := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if errorDB != nil {
		panic("Failed to connect mysql database")
	}

	return db

}

func DisconnectDB(db *gorm.DB) {
	dbSQl, err := db.DB()
	if err != nil {
		panic("Failed to terminated database")
	}
	dbSQl.Close()
}
