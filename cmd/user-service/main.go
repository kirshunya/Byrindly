package main

import (
	"byrindly/internal/db"
	model "byrindly/internal/user-service/model"
	"fmt"
	"time"
)

func main() {
	user := model.User{
		ID:               1,
		Name:             "Test User",
		Age:              25,
		Gender:           1, // Например, 1 для мужчины
		Latitude:         40.7128,
		Longitude:        -74.0060,
		RegistrationDate: time.Now(),
		About:            "This is a test user.",
		Username:         "testuser",
		Photo:            "link_to_test_photo",
		TgId:             123456789,
		Password:         "securepassword",
	}

	db.Connect()
	db.CreateTable()
	db.CreateUser(user)
	users, err := db.GetAllUsers()
	fmt.Println(err, users)
	db.Close()
}
