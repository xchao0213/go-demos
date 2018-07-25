package common

import (
	"github.com/joho/godotenv"
	"os"
)

func GetDataSource() (string) {
	err := godotenv.Load()
	
	CheckErr(err)
	
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	database := os.Getenv("DATABASE")

    datasource := username + ":" + password  + "@tcp(" + host + ":" + port + ")/" + database
	 
    return datasource
}