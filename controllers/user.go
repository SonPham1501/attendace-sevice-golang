package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IUserController interface {
	Retrieve()
}

type UserController struct{}

func (this *UserController) Retrieve(c *gin.Context) {
	if id := c.Param("id"); id != "" {
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
	c.Abort()
	return
}
