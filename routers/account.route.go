package routers

import (
	"com.son.server.attendaceservice/controllers"
	"github.com/gin-gonic/gin"
)

func accountRoute(g *gin.RouterGroup, ctrl *controllers.AccountController) {
	// Initlization database
	ctrl.Prepare()

	accountGroup := g.Group("account")
	{
		accountGroup.GET("", ctrl.GetAll)
		accountGroup.GET("/:id", ctrl.Get)
		accountGroup.POST("", ctrl.Post)
		accountGroup.PUT("/:id", ctrl.Put)
		accountGroup.DELETE("/:id", ctrl.Delete)
	}
}
