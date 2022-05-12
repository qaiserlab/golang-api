package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	c "golang.api/controllers"
	"golang.api/docs"
	"golang.api/models"
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
	const BasePath = "/api/v1"

	LoadEnvy()

	r := gin.Default()

	db := models.SetupModels()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	v1 := r.Group(BasePath)
	{
		users := v1.Group("/users")
		{
			users.GET("/", c.GetUsers)
			users.GET("/:id", c.GetUserById)
			users.POST("/", c.CreateUser)
			users.PUT("/:id", c.UpdateUserById)
			users.DELETE("/:id", c.DeleteUserById)
		}
	}

	docs.SwaggerInfo.Title = os.Getenv("APP_NAME")
	docs.SwaggerInfo.Description = os.Getenv("APP_DESCRIPTION")
	docs.SwaggerInfo.Version = os.Getenv("APP_VERSION")
	docs.SwaggerInfo.Host = "localhost:" + os.Getenv("PORT")
	docs.SwaggerInfo.BasePath = BasePath
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"APP_NAME":    os.Getenv("APP_NAME"),
			"APP_VERSION": os.Getenv("APP_VERSION"),
			"APP_SWAG":    "http://" + os.Getenv("HOST") + ":" + os.Getenv("PORT") + "/swagger/index.html",
		})
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run("localhost:" + os.Getenv("PORT"))
}
