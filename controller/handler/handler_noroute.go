package ctrlHandler

import (
	"github.com/gin-gonic/gin"

	"github.com/parinyapt/StreamySnap_AuthService/model/utils"
	"github.com/parinyapt/StreamySnap_AuthService/utils/response"
)

func NoRouteHandler(c *gin.Context) {
	utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
		ResponseCode: 404,
		Message: "Route Not Found",
	})
}