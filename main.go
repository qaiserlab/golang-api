package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.api/controllers"
	"golang.api/models"
)

func LoadEnvy() {
	err := godotenv.Load()

	if err != nil {
		panic("Error loading .env file")
	}
}

func main() {
	LoadEnvy()

	r := gin.Default()

	db := models.SetupModels()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data":    "GOLANG API",
			"version": "1.0.0",
		})
	})

	r.GET("/users", controllers.GetAllData)
	r.POST("/users", controllers.CreateData)
	r.PUT("/users/:id", controllers.UpdateData)
	r.DELETE("/users/:id", controllers.DeleteData)

	// host := "localhost"
	r.Run("localhost:" + os.Getenv("PORT"))
}
