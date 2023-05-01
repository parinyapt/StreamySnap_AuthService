package utilsResponse

import (
	"time"

	"github.com/gin-gonic/gin"

	"github.com/parinyapt/StreamySnap_AuthService/model/utils"
)

func JsonResponse(c *gin.Context, param modelUtils.JsonResponseStruct) {
	c.JSON(param.ResponseCode, modelUtils.JsonResponseStructDetail{
		Timestamp: time.Now().Unix(),
		Success:   param.Detail.Success,
		Message:   param.Detail.Message,
		ErrorCode: param.Detail.ErrorCode,
		Data:      param.Detail.Data,
	})
}