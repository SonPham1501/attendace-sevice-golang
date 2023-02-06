package base

import (
	"com.son.server.attendaceservice/db/mongodb"
	"github.com/gin-gonic/gin"
)

type BaseControllerInterface interface{
	GetAll(*gin.Context, interface{})
	Fetch(*gin.Context, interface{})
	CreateOne(*gin.Context, interface{})
	Update(*gin.Context, interface{})
	Remove(*gin.Context, interface{})
	Prepare(*gin.Context, interface{})
}

type BaseController struct {
	mongodb.MongoCrudService
}