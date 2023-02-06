package routers

import (
	"com.son.server.attendaceservice/controllers"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Recovery())
	// router.Use(middlewares.AuthMiddleware())

	// Service V1 Api
	addV1Services(router)

	return router
}

func addV1Services(router *gin.Engine) {
	v1 := router.Group("v1")
	{
		accountRoute(v1, new(controllers.AccountController))
	}
}
