package main

import (
	"byrindly/internal/user-service/model"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
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

func getUserById(c *gin.Context) {
	id := c.Param("id")
	userID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	url := fmt.Sprintf("http://localhost:8081/user/%d", userID)
	response, err := http.Get(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		return
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse data from User Service."})
	}
	c.Data(response.StatusCode, "application/json", body)
}

func updateUser(c *gin.Context) {
	id := c.Param("id")
	userID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	url := "http://localhost:8081/update/" + strconv.FormatUint(userID, 10)

	userData, err := json.Marshal(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal user data"})
		return
	}

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(userData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		c.JSON(response.StatusCode, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func deleteUser(c *gin.Context) {
	id := c.Param("id")
	userID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	url := "http://localhost:8081/delete/" + strconv.FormatUint(userID, 10)

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		c.JSON(response.StatusCode, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func main() {
	router := gin.Default()

	router.GET("/api/users", getUsers)
	router.GET("/api/user/:id", getUserById)
	router.POST("/api/user/create", createUser)
	router.PUT("/api/update/:id", updateUser)
	router.DELETE("/api/delete/:id", deleteUser)

	router.Run(":8080")
}
