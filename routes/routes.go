package routes

import (
	"consume-api-go-gin/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	r.GET("/get-data", controllers.GetData)
	r.GET("/products", controllers.Index)
	r.POST("/products", controllers.Create)
	r.GET("/products/:id", controllers.Show)
	r.PUT("/products/:id", controllers.Update)
	r.DELETE("/products/:id", controllers.Delete)
}
