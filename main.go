package main

import (
	"consume-api-go-gin/config"
	"consume-api-go-gin/database"
	"consume-api-go-gin/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnv()
	database.ConnectToDB()
}

func main() {
	r := gin.Default()
	routes.Routes(r)
	r.Run()
}
