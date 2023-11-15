package router

import (
	"blog/controller"
	"blog/middleware"

	"github.com/gin-gonic/gin"
)

func ApiRouterInit(r *gin.Engine) {

	apiR := r.Group("/api")
	{
		apiR.POST("/auth/register", controller.Register)
		apiR.POST("/auth/login", controller.Login)
		apiR.GET("/auth/info", middleware.AuthMiddleware(), controller.Info)
	}

}
