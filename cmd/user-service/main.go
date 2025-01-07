package main

import (
	"byrindly/internal/db"
	"byrindly/internal/user-service/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func getUsers(c *gin.Context) {
	users, err := db.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Users not found."})
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, users)
}

func createUser(c *gin.Context) {
	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.CreateUser(user)
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully.", "user": user})
}

func main() {

	db.Connect()
	db.CreateTable()

	defer db.Close()

	router := gin.Default()

	router.GET("/users", getUsers)
	router.POST("/create", createUser)

	err := router.Run(":8081")
	if err != nil {
		log.Fatal(err)
	}
}
