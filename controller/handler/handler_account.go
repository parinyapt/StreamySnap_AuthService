package ctrlHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	ctrlLogic "github.com/parinyapt/StreamySnap_AuthService/controller/logic"
	"github.com/parinyapt/StreamySnap_AuthService/logger"
	modelController "github.com/parinyapt/StreamySnap_AuthService/model/controller"
	modelUtils "github.com/parinyapt/StreamySnap_AuthService/model/utils"
	utilsResponse "github.com/parinyapt/StreamySnap_AuthService/utils/response"
	PTGUvalidator "github.com/parinyapt/golang_utils/validator/v1"
)

func CreateAccount(c *gin.Context) {
	var request modelController.RequestCreateAccount
	var response modelController.ResponseCreateAccount

	if err := c.ShouldBindJSON(&request); err != nil {
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: http.StatusBadRequest,
		})
		return
	}

	isValidatePass, errorFieldList, validatorError := PTGUvalidator.Validate(request)
	if validatorError != nil {
		logger.Error("[Handler][CreateAccount()]->Error Validate", logger.Field("error", validatorError.Error()))
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: http.StatusInternalServerError,
		})
		return
	}
	if !isValidatePass {
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: http.StatusBadRequest,
			Data:         errorFieldList,
		})
		return
	}

	dataFetchAccount, err := ctrlLogic.FetchOneAccount(modelController.ParamLogicFetchOneAccount{
		AccountEmail: request.AccountEmail,
	})
	if err != nil {
		logger.Error("[Handler][CreateAccount()]->Error FetchOneAccount", logger.Field("error", err.Error()))
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: http.StatusInternalServerError,
		})
		return
	}

	if dataFetchAccount.IsFound {
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: http.StatusConflict,
			Message:      "This Email is Already Exist",
		})
		return
	}

	dataCreateAccount, err := ctrlLogic.CreateAccount(modelController.ParamLogicCreateAccount{
		AccountName:     request.AccountName,
		AccountEmail:    request.AccountEmail,
		AccountPassword: request.AccountPassword,
	})
	if err != nil {
		logger.Error("[Handler][CreateAccount()]->Error CreateAccount", logger.Field("error", err.Error()))
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: http.StatusInternalServerError,
		})
		return
	}

	response.AccountUUID = dataCreateAccount.AccountUUID
	response.AccountName = dataCreateAccount.AccountName
	response.AccountEmail = dataCreateAccount.AccountEmail

	utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
		ResponseCode: http.StatusOK,
		Message:      "Create Account Success",
		Data:         response,
	})
}
