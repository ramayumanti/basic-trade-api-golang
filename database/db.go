package database

import (
	"basictrade-api/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func StartDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbport := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		host, user, password, dbname, dbport)
	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database: ", err)
	}

	db.Debug().AutoMigrate(models.Admin{}, models.Product{}, models.Variant{})

}

func GetDB() *gorm.DB {
	return db
}
