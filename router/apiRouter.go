package router

import (
	"blog/controller"

	"github.com/gin-gonic/gin"
)

func ApiRouterInit(r *gin.Engine) {

	apiR := r.Group("/api")
	{
		apiR.POST("/auth/register", controller.Register)
		apiR.POST("/auth/login", controller.Login)
	}

}
