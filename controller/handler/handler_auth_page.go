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

func GenerateAuthPageSessionURL(c *gin.Context) {
	var request modelController.RequestGenerateAuthPageSessionURL
	var response modelController.ResponseGenerateAuthPageSessionURL

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

	fetchData, err := ctrlLogic.FetchOneClientService(modelController.ParamLogicFetchOneClientService{ServiceUUID: request.ServiceUUID})
	if err != nil {
		logger.Error("[Handler][CreateClientService()]->Error FetchOneClientService", logger.Field("error", err.Error()))
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: http.StatusInternalServerError,
		})
		return
	}

	if !fetchData.IsFound {
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: http.StatusNotFound,
			Message:      "Service Not Found",
		})
		return
	}

	createData, err := ctrlLogic.CreateAuthPageSession(modelController.ParamLogicCreateAuthPageSession{
		ServiceUUID: fetchData.ClientServiceUUID,
	})
	if err != nil {
		logger.Error("[Handler][CreateClientService()]->Error CreateAuthPageSession", logger.Field("error", err.Error()))
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: http.StatusInternalServerError,
		})
		return
	}

	url, err := ctrlLogic.GenerateAuthPageSessionURL(modelController.ParamLogicGenerateAuthPageSessionURL{
		SessionUUID:        createData.SessionUUID,
		ExpireAt:           createData.ExpireAt,
	})
	if err != nil {
		logger.Error("[Handler][CreateClientService()]->Error GenerateAuthPageSessionURL", logger.Field("error", err.Error()))
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: http.StatusInternalServerError,
		})
		return
	}

	response.AuthPageSessionURL = url.AuthURL

	utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
		ResponseCode: http.StatusOK,
		Message:      "Create Auth Page Session Success",
		Data:         response,
	})
}

func CheckAuthPageSession(c *gin.Context) {
	var request modelController.RequestCheckAuthPageSessionToken
	var response modelController.ResponseCheckAuthPageSessionToken

	if err := c.ShouldBindJSON(&request); err != nil {
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: http.StatusBadRequest,
		})
		return
	}

	isValidatePass, errorFieldList, validatorError := PTGUvalidator.Validate(request)
	if validatorError != nil {
		logger.Error("[Handler][CheckAuthPageSession()]->Error Validate", logger.Field("error", validatorError.Error()))
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

	checkTokenData, err := ctrlLogic.CheckAuthPageSessionToken(modelController.ParamLogicCheckAuthPageSessionToken{
		SessionToken: request.SessionToken,
	})
	if err != nil {
		logger.Warning("[Handler][CheckAuthPageSession()]->Error CheckAuthPageSessionToken", logger.Field("error", err.Error()))
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: http.StatusBadRequest,
		})
		return
	}
	if !checkTokenData.IsValid {
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: http.StatusNotFound,
			Message:      "Session Not Found",
		})
		return
	}

	fetchDataSession, err := ctrlLogic.FetchOneAuthPageSession(modelController.ParamLogicFetchOneAuthPageSession{SessionUUID: checkTokenData.SessionUUID})
	if err != nil {
		logger.Error("[Handler][CheckAuthPageSession()]->Error FetchOneAuthPageSession", logger.Field("error", err.Error()))
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: http.StatusInternalServerError,
		})
		return
	}
	if !fetchDataSession.IsValid {
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: http.StatusNotFound,
			Message:      "Session Not Found",
		})
		return
	}

	fetchDataService, err := ctrlLogic.FetchOneClientService(modelController.ParamLogicFetchOneClientService{ServiceUUID: fetchDataSession.ClientServiceUUID})
	if err != nil {
		logger.Error("[Handler][CreateClientService()]->Error FetchOneClientService", logger.Field("error", err.Error()))
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: http.StatusInternalServerError,
		})
		return
	}
	if !fetchDataService.IsFound {
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: http.StatusNotFound,
			Message:      "Service Not Found",
		})
		return
	}

	response.ClientName = fetchDataService.ClientName
	response.ClientServiceName = fetchDataService.ClientServiceName
	response.ClientServiceCallbackURL = fetchDataService.ClientServiceCallbackURL

	utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
		ResponseCode: http.StatusOK,
		Message:      "Session Found",
		Data:         response,
	})
}
