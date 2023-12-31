package router

import (
	"blog/controller"
	"blog/middleware"

	"github.com/gin-gonic/gin"
)

func RouterInit(r *gin.Engine) {
	r.Use(middleware.CORSMiddlewqre(), middleware.RecoveryMiddleware())
	apiR := r.Group("/api")
	{
		apiR.POST("/auth/register", controller.Register)
		apiR.POST("/auth/login", controller.Login)
		apiR.GET("/auth/info", middleware.AuthMiddleware(), controller.Info)
	}

	categoryC := controller.NewCategoryController()
	categoryR := r.Group("/categories")
	{
		categoryR.POST("", categoryC.Create)
		categoryR.PUT("/:id", categoryC.Update)
		categoryR.GET("/:id", categoryC.Show)
		categoryR.DELETE("/:id", categoryC.Delete)
	}

	postC := controller.NewPostController()
	postR := r.Group("/posts")
	{
		postR.Use(middleware.AuthMiddleware())
		postR.POST("", postC.Create)
		postR.PUT("/:id", postC.Update)
		postR.GET("/:id", postC.Show)
		postR.DELETE("/:id", postC.Delete)
		postR.POST("page/list", postC.PageList)
	}

}
