package main

import (
	"fmt"
	// "fmt"
	// "net/http"

	// "github.com/gin-gonic/gin"
	// "github.com/gin-contrib/cors"
	"github.com/xchao0213/go-demos/server"
	"github.com/xchao0213/go-demos/common"
)

func main() {
	// r := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	// r.MaxMultipartMemory = 8 << 20 // 8 MiB
	// router.Static("/home", "./public")
	// r.Use(cors.Default())

	r := server.SetupRouter()

	var ds = common.GetDataSource()
	fmt.Println(ds)
	// router.POST("/user", func(c *gin.Context) {
	// 	name := c.PostForm("name")
	// 	email := c.PostForm("email")

	// 	// Source
	// 	file, err := c.FormFile("file")
	// 	if err != nil {
	// 		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
	// 		return
	// 	}

	// 	if err := c.SaveUploadedFile(file, file.Filename); err != nil {
	// 		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
	// 		return
	// 	}

	// 	c.JSON(http.StatusOK, getjson(file.Filename,name,email))
	// })
    // router.GET("/user", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, getjson("Teid.png","alpha","alpha@gmail.com"))
	// })
	// router.GET("/user/:id",func(c *gin.Context){

	// })

	// router.DELETE("/user",func(c *gin.Context){

	// })

	r.Run(":9090")
}

// func getjson(filename,name,email string) gin.H {
// 	return gin.H{"ImgUrl":filename,"name":name,"email":email}
// }