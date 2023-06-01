package services

import (
	"log"
	"os"
	"zappin/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Database struct {
	DB *gorm.DB
}

func (db *Database) ConnectDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file", err)
	}
	dburi := os.Getenv("DB_URI") //used a cockroachdb database but postgres is fine
	db.DB, err = gorm.Open(postgres.Open(dburi), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database", err)
	} else {
		log.Println("connected to database")
	}

	db.DB.AutoMigrate(&models.User{}, &models.Post{})
}
