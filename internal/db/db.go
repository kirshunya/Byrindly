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
		log.Printf("Failed to get all users : %s", result.Error)
		return nil, result.Error
	}
	return users, nil
}

func UpdateUserById(id uint64, newUser model.User) (model.User, error) {
	var existingUser model.User
	result := db.First(&existingUser, id)
	if result.Error != nil {
		log.Printf("Failed to get user with id = %d: %s", id, result.Error)
		return existingUser, result.Error
	}

	existingUser.Name = newUser.Name
	existingUser.Age = newUser.Age
	existingUser.Gender = newUser.Gender
	existingUser.Latitude = newUser.Latitude
	existingUser.Longitude = newUser.Longitude
	existingUser.RegistrationDate = newUser.RegistrationDate
	existingUser.About = newUser.About
	existingUser.Username = newUser.Username
	existingUser.Photo = newUser.Photo
	existingUser.TgId = newUser.TgId
	existingUser.Password = newUser.Password
	existingUser.UserId = newUser.UserId
	existingUser.Status = newUser.Status
	existingUser.ProfileLink = newUser.ProfileLink

	result = db.Save(&existingUser)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("User with ID %d not found.", id)
			return model.User{}, nil
		}
		log.Printf("Error retrieving user: %s", err)
		return model.User{}, err
	}

	return existingUser, nil
}

func GetUserById(id uint64) (model.User, error) {
	var user model.User
	result := db.First(&user, id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			log.Printf("User with ID %d not found.", id)
			return model.User{}, gorm.ErrRecordNotFound
		}
		log.Printf("Error retrieving user: %s", result.Error)
		return model.User{}, result.Error
	}

	return user, nil
}

func DeleteUserById(id uint64) error {
	var user model.User

	result := db.First(&user, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			log.Printf("User with ID %d not found.", id)
			return gorm.ErrRecordNotFound
		}
		log.Printf("Error retrieving user: %s", result.Error)
		return result.Error
	}

	if err := db.Delete(&user).Error; err != nil {
		log.Printf("Failed to delete user with ID %d: %s", id, err)
		return err
	}

	log.Printf("User with ID %d deleted successfully.", id)
	return nil
}
