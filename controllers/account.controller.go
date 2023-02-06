package controllers

import (
	"net/http"

	"com.son.server.attendaceservice/base"
	"com.son.server.attendaceservice/models"
	"github.com/gin-gonic/gin"
)

type AccountController struct {
	base.BaseController
}

func (this *AccountController) Prepare() {
	this.BindCollection("myFirstDatabase", "acounts")
}

func (this *AccountController) GetAll(c *gin.Context) {
	var result []models.Account
	if status := this.BaseController.GetAll(c, &result); status {
		response := models.AccountListResponse{
			Data:  result,
			Total: len(result),
		}
		c.JSON(http.StatusOK, response)
	}
}

func (this *AccountController) Get(c *gin.Context) {
	var result models.Account
	if status := this.BaseController.Fetch(c, &result); status {
		response := models.AccountResponse{
			Data: result,
		}
		c.JSON(http.StatusOK, response)
	}
}

func (this *AccountController) Post(c *gin.Context) {
	var result models.Account
	if uid := this.BaseController.CreateOne(c, &result); uid != nil {
		c.JSON(http.StatusOK, uid)
	}
}

func (this *AccountController) Put(c *gin.Context) {
	var result models.Account
	if status := this.BaseController.Update(c, &result); status {
		response := models.AccountResponse{
			Data: result,
		}
		c.JSON(http.StatusOK, response)
	}
}

func (this *AccountController) Delete(c *gin.Context) {
	var result models.Account
	if status := this.BaseController.Remove(c, &result); status {
		c.JSON(http.StatusOK, "Delete successfully")
	}
}
