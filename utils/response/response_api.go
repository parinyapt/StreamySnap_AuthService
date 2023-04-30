package utilsResponse

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/parinyapt/StreamySnap_AuthService/model/utils"
)

var ApiResponseConfigData = map[int]modelUtils.ApiResponseConfigStruct{
	http.StatusOK: {
		Message: "Success",
		ErrorCode: "0",
	},
	http.StatusBadRequest: {
		Message: "Bad Request",
		ErrorCode: "DF400",
	},
	http.StatusNotFound: {
		Message: "Not Found",
		ErrorCode: "DF404",
	},
	http.StatusInternalServerError: {
		Message: "Internal Server Error",
		ErrorCode: "DF500",
	},
}

func ApiResponse(c *gin.Context, param modelUtils.ApiResponseStruct) {
	if param.ResponseCode == 0 {
		param.ResponseCode = 500
	}

	var success_status bool
	if param.ResponseCode >= 200 && param.ResponseCode <= 299 {
		success_status = true
	} else {
		success_status = false
	}

	if param.Message == "" {
		param.Message = ApiResponseConfigData[param.ResponseCode].Message
	}

	if param.ErrorCode == "" {
		param.ErrorCode = ApiResponseConfigData[param.ResponseCode].ErrorCode
	}

	JsonResponse(c, modelUtils.JsonResponseStruct{
		ResponseCode: param.ResponseCode,
		Detail: modelUtils.JsonResponseStructDetail{
			Success:   success_status,
			Message:   param.Message,
			ErrorCode: param.ErrorCode,
			Data:      param.Data,
		},
	})
}
