package routes

import (
	"github.com/gin-gonic/gin"

	ctrlHandler "github.com/parinyapt/StreamySnap_AuthService/controller/handler"
)

func InitClientAPI(router *gin.RouterGroup) {
	r := router.Group("/client")
	{
		r.POST("/", ctrlHandler.CreateClient)
		r.DELETE("/", ctrlHandler.DeleteClient)
	}
}