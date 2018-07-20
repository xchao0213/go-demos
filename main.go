package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {
	router := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.Static("/home", "./public")
	router.Use(cors.Default())
	router.POST("/upload", func(c *gin.Context) {
		name := c.PostForm("name")
		email := c.PostForm("email")

		// Source
		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
			return
		}

		if err := c.SaveUploadedFile(file, file.Filename); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}

		c.JSON(http.StatusOK, getjson(file.Filename,name,email))
	})
    router.GET("/get", func(c *gin.Context) {
		c.JSON(http.StatusOK, getjson("Teid.png","alpha","alpha@gmail.com"))
	})

	router.Run(":9090")
}

func getjson(filename,name,email string) gin.H {
	return gin.H{"ImgUrl":filename,"name":name,"email":email}
}