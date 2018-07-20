package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func AddUser(c *gin.Context) {
  c.JSON(http.StatusOK,gin.H{"method":"addUser"})
}

func GetUser(c *gin.Context) {
	c.JSON(http.StatusOK,gin.H{"method":"getUser"})
}

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK,gin.H{"method":"getUsers"})
}

func UpdateUser(c *gin.Context) {
	c.JSON(http.StatusOK,gin.H{"method":"updateUser"})
}

func DeleteUser(c *gin.Context) {
	c.JSON(http.StatusOK,gin.H{"method":"deleteUser"})
}