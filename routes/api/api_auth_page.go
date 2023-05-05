package routes

import (
	"github.com/gin-gonic/gin"

	ctrlHandler "github.com/parinyapt/StreamySnap_AuthService/controller/handler"
)

func InitAuthPageAPI(router *gin.RouterGroup) {
	r := router.Group("/auth-page")
	{
		r.GET("/auth-url", ctrlHandler.GenerateAuthPageSessionURL)
		r.POST("/auth-url", ctrlHandler.CheckAuthPageSession)
	}
}