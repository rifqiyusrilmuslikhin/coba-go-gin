package main

import (
	"consume-api-go-gin/config"
	"consume-api-go-gin/database"
	"consume-api-go-gin/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnv()
	database.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	routes.Routes(r)
	r.Run()
}
