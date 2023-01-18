package main

import (
	"net/http"

	"github.com/adibpraditya/go-restapi-gin/controllers"
	"github.com/adibpraditya/go-restapi-gin/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	r.GET("/", oke)
	r.GET("api/products", controllers.Index)
	r.GET("/api/product/:id", controllers.Show)
	r.POST("/api/product", controllers.Create)
	r.PUT("/api/product/:id", controllers.Update)
	r.DELETE("/api/product", controllers.Delete)

	r.Run()
}

func oke(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
