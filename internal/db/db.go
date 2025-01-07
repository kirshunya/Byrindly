package db

import (
	model "byrindly/internal/user-service/model"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var db *gorm.DB

func Connect() {
	envFilePath := "D:\\Byrindly\\.env"
	err := godotenv.Load(envFilePath)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
}

func Close() {
	db = nil
}

func CreateTable() {
	err := db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}

func CreateUser(user model.User) error {
	result := db.Create(&user)
	if result.Error != nil {
		log.Fatalf("Failed to create user %s", result.Error)
		return result.Error
	} else {
		log.Printf("User created: \nID:%d\nName:%s\n", user.ID, user.Name)
		return nil
	}
}

func GetAllUsers() ([]model.User, error) {
	var users []model.User
	result := db.Find(&users)
	if result.Error != nil {
		log.Printf("Failed to get all users %s", result.Error)
		return nil, result.Error
	}
	return users, nil
}
