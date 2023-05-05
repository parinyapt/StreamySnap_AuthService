package routes

import (
	"github.com/gin-gonic/gin"

	ctrlHandler "github.com/parinyapt/StreamySnap_AuthService/controller/handler"
)

func InitAccountAPI(router *gin.RouterGroup) {
	r := router.Group("/account")
	{
		r.POST("/", ctrlHandler.CreateAccount)
	}
}