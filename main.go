package main

import (
	"fmt"
	"net/http"
	"os"

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

	r := gin.Default()

	db := models.SetupModels()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	routes.LoadV1Router(r)

	basePath := os.Getenv("HOST") + ":" + os.Getenv("PORT")
	baseUrl := "http://" + basePath

	fmt.Println(baseUrl)

	docs.SwaggerInfo.Title = os.Getenv("APP_NAME")
	docs.SwaggerInfo.Description = os.Getenv("APP_DESCRIPTION")
	docs.SwaggerInfo.Version = os.Getenv("APP_VERSION")
	docs.SwaggerInfo.Host = basePath
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"APP_NAME":    os.Getenv("APP_NAME"),
			"APP_VERSION": os.Getenv("APP_VERSION"),
			"APP_SWAG":    baseUrl + "/swagger/index.html",
		})
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run("localhost:" + os.Getenv("PORT"))
}
