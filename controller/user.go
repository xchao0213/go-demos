package controller

import (
	"fmt"
	"time"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"github.com/xchao0213/go-demos/model"
	"github.com/xchao0213/go-demos/common"
)

func AddUser(c *gin.Context) {
	name := c.PostForm("name")
	birthday := c.PostForm("birthday")
	sex := c.PostForm("sex")
	identify_card := c.PostForm("identify_card")
	mobile := c.PostForm("mobile")

	photo, formerr := c.FormFile("photo")
	common.CheckErr(formerr)

    fileerr := c.SaveUploadedFile(photo, "./uploads/" + photo.Filename)
	common.CheckErr(fileerr)
    uid := uuid.Must(uuid.NewV4()).String()
	u := model.User{
		Id: uid,
		Name: name,
		Birthday: birthday,
		Sex: sex,
		IdentifyCard: identify_card,
		Mobile: mobile,
		Photo: "http://localhost:9090/" + photo.Filename,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix() }

	res,err := model.UserInsert(&u)

	common.CheckErr(err)

	fmt.Println(res)
	
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