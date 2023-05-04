package ctrlHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	PTGUvalidator "github.com/parinyapt/golang_utils/validator/v1"

	utilsResponse "github.com/parinyapt/StreamySnap_AuthService/utils/response"

	ctrlLogic "github.com/parinyapt/StreamySnap_AuthService/controller/logic"
	modelController "github.com/parinyapt/StreamySnap_AuthService/model/controller"
	modelUtils "github.com/parinyapt/StreamySnap_AuthService/model/utils"
)

func CreateClient(c *gin.Context) {
	var request modelController.RequestCreateClient

	if err := c.ShouldBindJSON(&request); err != nil {
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: http.StatusBadRequest,
		})
		return
	}

	isValidatePass, errorFieldList, validatorError := PTGUvalidator.Validate(request)
	if validatorError != nil {
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

	dataCheckExist, err := ctrlLogic.CheckExistClient(modelController.ParamLogicCheckExistClient{ClientName: request.ClientName})
	if err != nil {
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: http.StatusInternalServerError,
		})
		return
	}
	if dataCheckExist.IsExist {
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: http.StatusBadRequest,
			Message:      "Client Name Already Exist",
		})
		return
	}

	dataCreate, err := ctrlLogic.CreateClient(modelController.ParamLogicCreateClient{ClientName: request.ClientName})
	if err != nil {
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: http.StatusInternalServerError,
		})
		return
	}

	utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
		ResponseCode: http.StatusOK,
		Message:      "Create Client Success",
		Data:         dataCreate,
	})

}

func DeleteClient(c *gin.Context) {
	var request modelController.RequestDeleteClient

	if err := c.ShouldBindJSON(&request); err != nil {
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: http.StatusBadRequest,
		})
		return
	}

	isValidatePass, errorFieldList, validatorError := PTGUvalidator.Validate(request)
	if validatorError != nil {
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

	err = ctrlLogic.DeleteClient(modelController.ParamLogicDeleteClient{ClientUUID: request.ClientUUID})
	if err != nil {
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: http.StatusInternalServerError,
		})
		return
	}

	utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
		ResponseCode: http.StatusOK,
		Message:      "Delete Client Success",
	})
}
