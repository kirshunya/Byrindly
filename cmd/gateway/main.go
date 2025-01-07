package main

import (
	"byrindly/internal/user-service/model"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func getUsers(c *gin.Context) {
	response, err := http.Get("http://localhost:8081/users")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users."})
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse data from User Service."})
	}
	c.Data(response.StatusCode, "application/json", body)
}

func createUser(c *gin.Context) {
	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userData, err := json.Marshal(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response, err := http.Post("http://localhost:8081/create", "application/json", bytes.NewBuffer(userData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send request to user service."})
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusCreated {
		var errorResponse gin.H
		if err := json.NewDecoder(response.Body).Decode(&errorResponse); err == nil {
			c.JSON(response.StatusCode, gin.H{"error": errorResponse["error"]})
		} else {
			c.JSON(response.StatusCode, gin.H{"error": "Error from user service."})
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User data forwarded successfully."})
}

func main() {
	router := gin.Default()

	router.GET("/api/users", getUsers)
	router.POST("/api/user/create", createUser)

	router.Run(":8080")
}
