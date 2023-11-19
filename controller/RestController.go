package controller

import "github.com/gin-gonic/gin"

type RestController interface {
	Create(*gin.Context)
	Update(*gin.Context)
	Show(*gin.Context)
	Delete(*gin.Context)
}
