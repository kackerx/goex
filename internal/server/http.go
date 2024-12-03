package server

import (
	"github.com/gin-gonic/gin"

	"goex1/internal/conf"
	"goex1/internal/handler"
)

func NewHTTPServer(
	cfg *conf.Conf,
	userHandler *handler.UserHandler,
) *gin.Engine {
	g := gin.Default()
	gin.SetMode(gin.DebugMode)

	v1 := g.Group("/v1")
	userRouter := v1.Group("/user")
	{
		noAuthRouter := userRouter.Group("/")
		{
			noAuthRouter.GET("/hello", userHandler.Hello)
			noAuthRouter.POST("/register", userHandler.RegisterUser)
		}

		authRouter := userRouter.Group("/").Use()
		{
			authRouter.GET("/:id", userHandler.GetUserInfo)
		}

	}

	return g
}
