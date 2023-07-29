package main

import (
	"consume-api-go-gin/config"
	"consume-api-go-gin/database"
	"consume-api-go-gin/models"
	"fmt"
)

func init() {
	config.LoadEnv()
	database.ConnectToDB()
}

func main() {
	database.DB.AutoMigrate(&models.Product{})
	fmt.Println("Migrations success!")
}
