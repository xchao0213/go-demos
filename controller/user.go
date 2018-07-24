package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	// "github.com/xchao0213/go-demos/model"
)


func AddUser(c *gin.Context) {
	// model.DoInsert()
  c.JSON(http.StatusOK,gin.H{"method":"addUser"})
}

func GetUser(c *gin.Context) {
	c.JSON(http.StatusOK,gin.H{"method":"getUser"})
}

func GetUsers(c *gin.Context) {
	// model.DoQuery()
	c.JSON(http.StatusOK,gin.H{"method":"getUsers"})
}

func UpdateUser(c *gin.Context) {
	c.JSON(http.StatusOK,gin.H{"method":"updateUser"})
}

func DeleteUser(c *gin.Context) {
	c.JSON(http.StatusOK,gin.H{"method":"deleteUser"})
}