package ctrlHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	PTGUvalidator "github.com/parinyapt/golang_utils/validator/v1"

	utilsResponse "github.com/parinyapt/StreamySnap_AuthService/utils/response"

	ctrlLogic "github.com/parinyapt/StreamySnap_AuthService/controller/logic"
	"github.com/parinyapt/StreamySnap_AuthService/logger"
	modelController "github.com/parinyapt/StreamySnap_AuthService/model/controller"
	modelUtils "github.com/parinyapt/StreamySnap_AuthService/model/utils"
)

func CreateClientService(c *gin.Context) {
	var request modelController.RequestCreateClientService
	var response modelController.ResponseCreateClientService

	if err := c.ShouldBindJSON(&request); err != nil {
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: http.StatusBadRequest,
		})
		return
	}

	isValidatePass, errorFieldList, validatorError := PTGUvalidator.Validate(request)
	if validatorError != nil {
		logger.Error("[Handler][CreateClientService()]->Error Validate", logger.Field("error", validatorError.Error()))
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

	dataCheckExist, err := ctrlLogic.CheckExistClient(modelController.ParamLogicCheckExistClient{ClientUUID: request.ClientUUID})
	if err != nil {
		logger.Error("[Handler][CreateClientService()]->Error CheckExistClient", logger.Field("error", err.Error()))
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: http.StatusInternalServerError,
		})
		return
	}
	if !dataCheckExist.IsExist {
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: http.StatusNotFound,
			Message:      "Client Not Found",
		})
		return
	}

	dataCreate, err := ctrlLogic.CreateClientService(modelController.ParamLogicCreateClientService{
		ClientUUID:         request.ClientUUID,
		ServiceName:        request.ServiceName,
		ServiceCallbackURL: request.ServiceCallbackURL,
	})
	if err != nil {
		logger.Error("[Handler][CreateClientService()]->Error CreateClientService", logger.Field("error", err.Error()))
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: http.StatusInternalServerError,
		})
		return
	}

	response.ServiceName = dataCreate.ServiceName
	response.ServiceUUID = dataCreate.ServiceUUID
	response.ServiceSecretKey = dataCreate.ServiceSecretKey

	utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
		ResponseCode: http.StatusOK,
		Message:      "Create Client Service Success",
		Data:         response,
	})

}
