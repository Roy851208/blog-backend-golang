package main

import (
	"blog/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	router.ApiRouterInit(r)

	panic(r.Run())
}
