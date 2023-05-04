package routes

import (
	"github.com/gin-gonic/gin"

	ctrlHandler "github.com/parinyapt/StreamySnap_AuthService/controller/handler"
)

func InitClientServiceAPI(router *gin.RouterGroup) {
	r := router.Group("/client-service")
	{
		r.POST("/", ctrlHandler.CreateClientService)
	}
}