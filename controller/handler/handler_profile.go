package ctrlHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ctrlLogic "github.com/parinyapt/StreamySnap_AuthService/controller/logic"
	"github.com/parinyapt/StreamySnap_AuthService/logger"
	modelController "github.com/parinyapt/StreamySnap_AuthService/model/controller"
	modelUtils "github.com/parinyapt/StreamySnap_AuthService/model/utils"
	utilsResponse "github.com/parinyapt/StreamySnap_AuthService/utils/response"
)

func GetProfile(c *gin.Context) {
	var response modelController.ResponseGetProfile

	dataFetchAccount, err := ctrlLogic.FetchOneAccount(modelController.ParamLogicFetchOneAccount{
		AccountUUID: c.GetString("auth_account_uuid"),
	})
	if err != nil {
		logger.Error("[Handler][GetProfile()]->Error FetchOneAccount", logger.Field("error", err.Error()))
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: http.StatusInternalServerError,
		})
		return
	}
	if !dataFetchAccount.IsFound {
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: http.StatusNotFound,
		})
		return
	}

	dataFetchService, err := ctrlLogic.FetchOneClientService(modelController.ParamLogicFetchOneClientService{
		ServiceUUID: c.GetString("auth_service_uuid"),
	})
	if err != nil {
		logger.Error("[Handler][GetProfile()]->Error FetchOneClientService", logger.Field("error", err.Error()))
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: http.StatusInternalServerError,
		})
		return
	}
	if !dataFetchService.IsFound {
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: http.StatusNotFound,
		})
		return
	}

	response.Account.AccountUUID = dataFetchAccount.AccountUUID
	response.Account.AccountEmail = dataFetchAccount.AccountEmail
	response.Account.AccountName = dataFetchAccount.AccountName
	response.Service.ServiceUUID = dataFetchService.ClientServiceUUID
	response.Service.ServiceName = dataFetchService.ClientServiceName

	utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
		ResponseCode: http.StatusOK,
		Message:      "Get Profile Success",
		Data:         response,
	})

}
