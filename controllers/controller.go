package controllers

import (
	"consume-api-go-gin/database"
	"consume-api-go-gin/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type ProductRequest struct {
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Image       string  `json:"image"`
}

func GetData(c *gin.Context) {
	client := resty.New()
	resp, err := client.R().
		EnableTrace().
		Get("https://fakestoreapi.com/products")
	if err != nil {
		c.JSON(500, gin.H{
			"message": "failed fetch data",
		})
		return
	}

	var data []models.DataTemp
	err = json.Unmarshal(resp.Body(), &data)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "failed to unmarshal data",
		})
		return
	}

	for _, req := range data {
		product := models.Product{
			ID:          "products" + "-" + strconv.Itoa(req.ID),
			Title:       req.Title,
			Price:       req.Price,
			Description: req.Description,
			Category:    req.Category,
			Image:       req.Image,
		}

		err := database.DB.Create(&product).Error
		if err != nil {
			c.JSON(500, gin.H{
				"message": "failed to insert data to database",
			})
			return
		}
	}

	c.JSON(200, gin.H{
		"message": "data successfully inserted to database",
	})
}

func Index(c *gin.Context) {
	var products []models.Product
	database.DB.Find(&products)

	c.JSON(200, gin.H{
		"data": products,
	})
}

func Show(c *gin.Context) {
	id := c.Param("id")

	var product models.Product

	err := database.DB.First(&product, "id = ?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(404, gin.H{
				"message": "Product not found",
			})
			return
		}
		c.JSON(500, gin.H{
			"message": "Failed to get product",
		})
		return
	}

	c.JSON(200, gin.H{
		"data": product,
	})
}

func Create(c *gin.Context) {
	var req ProductRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request data",
			"errors":  err.Error(),
		})
		return
	}

	if req.Title == "" || req.Description == "" || req.Category == "" || req.Image == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "All fields are required",
		})
		return
	}

	if req.Price <= 0 {
		c.JSON(400, gin.H{
			"message": "Invalid price value",
		})
		return
	}

	id, _ := gonanoid.New(16)

	product := models.Product{
		ID:          "product" + "-" + id,
		Title:       req.Title,
		Price:       req.Price,
		Description: req.Description,
		Category:    req.Category,
		Image:       req.Image,
	}

	database.DB.Create(&product)

	c.JSON(201, gin.H{
		"data": product,
	})
}

func Update(c *gin.Context) {
	id := c.Param("id")

	var req ProductRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request data",
			"errors":  err.Error(),
		})
		return
	}

	var product models.Product

	err := database.DB.First(&product, "id = ?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(404, gin.H{
				"message": "Product not found",
			})
			return
		}
		c.JSON(500, gin.H{
			"message": "Failed to get product",
		})
		return
	}

	if req.Title == "" || req.Description == "" || req.Category == "" || req.Image == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "All fields are required",
		})
		return
	}

	if req.Price <= 0 {
		c.JSON(400, gin.H{
			"message": "Invalid price value",
		})
		return
	}

	product.Title = req.Title
	product.Price = req.Price
	product.Category = req.Category
	product.Description = req.Description
	product.Image = req.Image

	database.DB.Save(&product)

	c.JSON(201, gin.H{
		"data": product,
	})
}

func Delete(c *gin.Context) {
	id := c.Param("id")

	var product models.Product

	err := database.DB.First(&product, "id = ?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(404, gin.H{
				"message": "Product not found",
			})
			return
		}
		c.JSON(500, gin.H{
			"message": "Failed to get product",
		})
		return
	}

	database.DB.Delete(&product)

	c.JSON(200, gin.H{
		"message": "Success deleted product",
	})
}
