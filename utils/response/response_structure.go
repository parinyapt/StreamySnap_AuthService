package utilsResponse

import (
	"github.com/gin-gonic/gin"

	"github.com/parinyapt/StreamySnap_AuthService/model/utils"
)

func JsonResponse(c *gin.Context, param modelUtils.JsonResponseStruct) {
	c.JSON(param.ResponseCode, modelUtils.JsonResponseStructDetail{
		Success:   param.Detail.Success,
		Message:   param.Detail.Message,
		ErrorCode: param.Detail.ErrorCode,
		Data:      param.Detail.Data,
	})
}