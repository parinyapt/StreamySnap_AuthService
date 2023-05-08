package routes

import (
	"github.com/gin-gonic/gin"

	ctrlHandler "github.com/parinyapt/StreamySnap_AuthService/controller/handler"
	"github.com/parinyapt/StreamySnap_AuthService/middleware"
)

func InitAuthAPI(router *gin.RouterGroup) {
	r := router.Group("/auth")
	{
		r.POST("/login", ctrlHandler.AuthLoginWithUsernamePassword)
		r.POST("/token", ctrlHandler.AuthConvertTempTokenToRealToken)
		

		r.GET("/logout", middleware.CheckAuthAccessToken, ctrlHandler.AuthLogout)
		
		r.GET("/refresh-token", middleware.CheckAuthRefreshToken, ctrlHandler.AuthRefreshToken)
		
	}
}