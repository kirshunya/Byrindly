package main

import (
	"byrindly/internal/db"
	"byrindly/internal/user-service/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
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
	err := db.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully.", "user": user})
}

func updateUserById(c *gin.Context) {
	var newUser model.User

	id := c.Param("id")
	userID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedUser, err := db.UpdateUserById(userID, newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully.", "user": updatedUser})
}

func getUserById(c *gin.Context) {
	id := c.Param("id")
	userID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	user, err := db.GetUserById(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User successfully found.", "user": user})
}

func main() {

	db.Connect()
	db.CreateTable()

	defer db.Close()

	router := gin.Default()

	//TODO: Delete by ID,
	router.GET("/users", getUsers)
	router.GET("/user/:id", getUserById)
	router.POST("/create", createUser)
	router.PUT("/update/:id", updateUserById)

	err := router.Run(":8081")
	if err != nil {
		log.Fatal(err)
	}
}
