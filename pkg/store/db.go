package store

import (
	"fmt"
	types "s21/surdo/pkg/model"

	"log"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	// "music-library/pkg/model"
	"os"
)

var DB *gorm.DB

func InitDB() {
	// .env fayldan ma'lumotlarni o'qish
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	
	// DB.AutoMigrate(&types.TUser{})

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
    os.Getenv("DB_HOST"),
    os.Getenv("DB_USER"),
    os.Getenv("DB_PASSWORD"),
    os.Getenv("DB_NAME"),
    os.Getenv("DB_PORT"),
    os.Getenv("DB_SSLMODE"),
)

	// connect data base
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %s", err)
	}
	DB.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)

	// Jadval yaratish
	err = DB.AutoMigrate(&types.TUser{})
	if err != nil {
		log.Fatalf("Error migrating database: %s", err)
	}
}

func GetDB() *gorm.DB {
	return DB
}
