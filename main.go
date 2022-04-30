package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"golang.api/controllers"
	"golang.api/docs"
	"golang.api/models"
)

func LoadEnvy() {
	err := godotenv.Load()

	if err != nil {
		panic("Error loading .env file")
	}
}

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
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

	v1 := r.Group("/api/v1")
	{
		users := v1.Group("/users")
		{
			users.GET("/", controllers.GetAllData)
			users.GET("/:id", controllers.GetOneData)
			users.POST("/", controllers.CreateData)
			users.PUT("/:id", controllers.UpdateData)
			users.DELETE("/:id", controllers.DeleteData)
		}
	}

	docs.SwaggerInfo.Title = "Golang API"
	docs.SwaggerInfo.Description = "This is a Golang API server."
	docs.SwaggerInfo.Version = "1.0.0"
	docs.SwaggerInfo.Host = "localhost:1234/swagger/index.html"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run("localhost:" + os.Getenv("PORT"))
}
