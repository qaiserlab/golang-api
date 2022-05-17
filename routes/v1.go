package routes

import (
	"github.com/gin-gonic/gin"
	"golang.api/controllers/auth"
	"golang.api/controllers/user"
	a "golang.api/middlewares/auth"
)

const basePath = "/v1"

func LoadV1Router(r *gin.Engine) {
	v1 := r.Group(basePath)

	{
		authRouter := v1.Group("/auth")
		{
			authRouter.POST("/login", auth.Login)
		}

		userRouter := v1.Group("/users", a.AuthMiddleware())
		{
			userRouter.GET("/", user.GetRecords)
			userRouter.GET("/:id", user.GetRecordById)
			userRouter.POST("/", user.CreateRecord)
			userRouter.PUT("/:id", user.UpdateRecordById)
			userRouter.DELETE("/:id", user.DeleteRecordById)
		}
	}

}
