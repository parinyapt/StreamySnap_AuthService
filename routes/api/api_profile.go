package routes

import (
	"github.com/gin-gonic/gin"

	ctrlHandler "github.com/parinyapt/StreamySnap_AuthService/controller/handler"
	"github.com/parinyapt/StreamySnap_AuthService/middleware"
)

func InitProfileAPI(router *gin.RouterGroup) {
	r := router.Group("/profile")
	{
		r.GET("/", middleware.CheckAuthAccessToken, ctrlHandler.GetProfile)
	}
}