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
			"APP_NAME":    os.Getenv("APP_NAME"),
			"APP_VERSION": os.Getenv("APP_VERSION"),
		})
	})

	r.GET("/users", controllers.GetAllData)
	r.GET("/users/:id", controllers.GetOneData)
	r.POST("/users", controllers.CreateData)
	r.PUT("/users/:id", controllers.UpdateData)
	r.DELETE("/users/:id", controllers.DeleteData)

	// host := "localhost"
	r.Run("localhost:" + os.Getenv("PORT"))
}
