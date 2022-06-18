package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"golang.api/docs"
	"golang.api/models"
	"golang.api/routes"
)

func LoadEnvy() {
	err := godotenv.Load()

	if err != nil {
		panic("Error loading .env file")
	}
}

// @contact.name   Fadlun Anaturdasa
// @contact.url    https://qaiserlab.github.io
// @contact.email  f.anaturdasa@gmail.com
func main() {
	LoadEnvy()

	APP_NAME := os.Getenv("APP_NAME")
	APP_DESCRIPTION := os.Getenv("APP_DESCRIPTION")
	APP_VERSION := os.Getenv("APP_VERSION")
	SERVER_URL := os.Getenv("SERVER_URL")
	SERVER_HOST := strings.Split(SERVER_URL, "://")[1]
	CLIENT_URL := os.Getenv("CLIENT_URL")

	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{CLIENT_URL}
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE"}

	r.Use(cors.New(config))

	db := models.SetupModels()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	routes.LoadV1Router(r)

	docs.SwaggerInfo.Title = APP_NAME
	docs.SwaggerInfo.Description = APP_DESCRIPTION
	docs.SwaggerInfo.Version = APP_VERSION
	docs.SwaggerInfo.Host = SERVER_HOST
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"APP_NAME":    APP_NAME,
			"APP_VERSION": APP_VERSION,
			"APP_SWAG":    SERVER_URL + "/swagger/index.html",
		})
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	fmt.Println("Server run on:")
	fmt.Println(SERVER_URL)

	r.Run(SERVER_HOST)
}
