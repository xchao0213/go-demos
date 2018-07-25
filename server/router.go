package server

import (
	"github.com/gin-gonic/gin"
	"github.com/xchao0213/go-demos/controller"
)

func SetupRouter() *gin.Engine {
	r:= gin.Default();

	r.Static("/images","./uploads")
    //get all
	r.GET("/user",controller.GetUsers)
	//get single
	r.GET("/user/:name",controller.GetUser)
    //add
	r.POST("/user",controller.AddUser)
	//update
	r.PUT("/user",controller.UpdateUser)
	//delete
	r.DELETE("/user",controller.DeleteUser)
	
	return r
}