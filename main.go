package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.api/models"
)

func main() {

	r := gin.Default()

	db := models.SetupModels()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": "Halo Dunia",
		})
	})

	r.Run("localhost:1234")
}
