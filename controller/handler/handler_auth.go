package ctrlHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	PTGUvalidator "github.com/parinyapt/golang_utils/validator/v1"

	utilsResponse "github.com/parinyapt/StreamySnap_AuthService/utils/response"
	utilsHeader "github.com/parinyapt/StreamySnap_AuthService/utils/header"
	
	ctrlLogic "github.com/parinyapt/StreamySnap_AuthService/controller/logic"
	"github.com/parinyapt/StreamySnap_AuthService/logger"
	modelController "github.com/parinyapt/StreamySnap_AuthService/model/controller"
	modelUtils "github.com/parinyapt/StreamySnap_AuthService/model/utils"
)

func AuthLoginWithUsernamePassword(c *gin.Context) {
	var request modelController.RequestAuthLoginWithUsernamePassword
	var response modelController.ResponseAuthLoginWithUsernamePassword

	if err := c.ShouldBindJSON(&request); err != nil {
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: http.StatusBadRequest,
		})
		return
	}

	isValidatePass, errorFieldList, validatorError := PTGUvalidator.Validate(request)
	if validatorError != nil {
		logger.Error("[Handler][AuthLoginWithUsernamePassword()]->Error Validate", logger.Field("error", validatorError.Error()))
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

	authToken, err := utilsHeader.GetHeaderAuthorizationValue(c)
	if err != nil {
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: http.StatusUnauthorized,
			Message:      "Authorization Header Not Found",
		})
		return
	}

	checkTokenData, err := ctrlLogic.CheckAuthPageSessionToken(modelController.ParamLogicCheckAuthPageSessionToken{
		SessionToken: authToken,
	})
	if err != nil {
		logger.Warning("[Handler][AuthLoginWithUsernamePassword()]->Error CheckAuthPageSessionToken", logger.Field("error", err.Error()))
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: http.StatusBadRequest,
		})
		return
	}
	if !checkTokenData.IsValid {
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: http.StatusUnauthorized,
			Message:      "Session Not Found",
		})
		return
	}

	fetchDataSession, err := ctrlLogic.FetchOneAuthPageSession(modelController.ParamLogicFetchOneAuthPageSession{SessionUUID: checkTokenData.SessionUUID})
	if err != nil {
		logger.Error("[Handler][AuthLoginWithUsernamePassword()]->Error FetchOneAuthPageSession", logger.Field("error", err.Error()))
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: http.StatusInternalServerError,
		})
		return
	}
	if !fetchDataSession.IsValid {
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: http.StatusUnauthorized,
			Message:      "Session Not Found",
		})
		return
	}

	dataFetchAccount, err := ctrlLogic.FetchOneAccount(modelController.ParamLogicFetchOneAccount{
		AccountEmail: request.AccountEmail,
	})
	if err != nil {
		logger.Error("[Handler][AuthLoginWithUsernamePassword()]->Error FetchOneAccount", logger.Field("error", err.Error()))
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: http.StatusInternalServerError,
		})
		return
	}

	if !dataFetchAccount.IsFound {
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: http.StatusUnauthorized,
			Message:      "Email or Password is invalid",
		})
		return
	}

	checkPwdMatch, err := ctrlLogic.CheckAuthLoginMatchPassword(modelController.ParamLogicCheckAuthLoginMatchPassword{
		AccountPassword: request.AccountPassword,
		AccountPasswordHash: dataFetchAccount.AccountPassword,
	})
	if err != nil {
		logger.Error("[Handler][AuthLoginWithUsernamePassword()]->Error CheckAuthLoginMatchPassword", logger.Field("error", err.Error()))
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: http.StatusInternalServerError,
		})
		return
	}
	if !checkPwdMatch.IsMatch {
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: http.StatusUnauthorized,
			Message:      "Email or Password is invalid",
		})
		return
	}

	dataCreateTempToken, err := ctrlLogic.CreateAuthTemporaryToken(modelController.ParamLogicCreateAuthTemporaryToken{
		AccountUUID: dataFetchAccount.AccountUUID,
		ServiceUUID: fetchDataSession.ClientServiceUUID,
	})
	if err != nil {
		logger.Error("[Handler][AuthLoginWithUsernamePassword()]->Error CreateAuthTemporaryToken", logger.Field("error", err.Error()))
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: http.StatusInternalServerError,
		})
		return
	}

	dataGenTempToken, err := ctrlLogic.GenerateAuthTemporaryToken(modelController.ParamLogicGenerateAuthTemporaryToken{
		TempTokenUUID: dataCreateTempToken.TempTokenUUID,
		ExpireAt: 		dataCreateTempToken.ExpireAt,
	})
	if err != nil {
		logger.Error("[Handler][AuthLoginWithUsernamePassword()]->Error GenerateAuthTemporaryToken", logger.Field("error", err.Error()))
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: http.StatusInternalServerError,
		})
		return
	}

	dataFetchClientService, err := ctrlLogic.FetchOneClientService(modelController.ParamLogicFetchOneClientService{
		ServiceUUID: fetchDataSession.ClientServiceUUID,
	})
	if err != nil {
		logger.Error("[Handler][AuthLoginWithUsernamePassword()]->Error FetchOneClientService", logger.Field("error", err.Error()))
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: http.StatusInternalServerError,
		})
		return
	}

	response.AuthCallbackURL = dataFetchClientService.ClientServiceCallbackURL + "?code=" + dataGenTempToken.TempToken
	response.Code = dataGenTempToken.TempToken

	utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
		ResponseCode: http.StatusOK,
		Message:      "Login Success",
		Data:         response,
	})
}

func AuthConvertTempTokenToRealToken(c *gin.Context) {
	var request modelController.RequestAuthConvertTempTokenToRealToken
	var response modelController.ResponseAuthConvertTempTokenToRealToken

	if err := c.ShouldBindJSON(&request); err != nil {
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: http.StatusBadRequest,
		})
		return
	}

	isValidatePass, errorFieldList, validatorError := PTGUvalidator.Validate(request)
	if validatorError != nil {
		logger.Error("[Handler][AuthConvertTempTokenToRealToken()]->Error Validate", logger.Field("error", validatorError.Error()))
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

	checkTempToken, err := ctrlLogic.CheckAuthTemporaryToken(modelController.ParamLogicCheckAuthTemporaryToken{
		TempToken: request.TempToken,
	})
	if err != nil {
		logger.Warning("[Handler][AuthConvertTempTokenToRealToken()]->Error CheckAuthTemporaryToken", logger.Field("error", err.Error()))
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: http.StatusBadRequest,
		})
		return
	}
	if !checkTempToken.IsValid {
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: http.StatusUnauthorized,
		})
		return
	}

	fetchDataTempToken, err := ctrlLogic.FetchOneAuthTemporaryToken(modelController.ParamLogicFetchOneAuthTemporaryToken{
		TempTokenUUID: checkTempToken.TempTokenUUID,
	})
	if err != nil {
		logger.Error("[Handler][AuthConvertTempTokenToRealToken()]->Error FetchOneAuthTemporaryToken", logger.Field("error", err.Error()))
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: http.StatusInternalServerError,
		})
		return
	}
	if !fetchDataTempToken.IsValid {
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: http.StatusUnauthorized,
		})
		return
	}

	err = ctrlLogic.DeleteAuthTemporaryToken(modelController.ParamLogicDeleteAuthTemporaryToken{
		TempTokenUUID: checkTempToken.TempTokenUUID,
	})
	if err != nil {
		logger.Error("[Handler][AuthConvertTempTokenToRealToken()]->Error DeleteAuthTemporaryToken", logger.Field("error", err.Error()))
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: http.StatusInternalServerError,
		})
		return
	}

	dataCreateRealToken, err := ctrlLogic.CreateAuthRealToken(modelController.ParamLogicCreateAuthRealToken{
		AccountUUID: fetchDataTempToken.AccountUUID,
		ServiceUUID: fetchDataTempToken.ServiceUUID,
	})
	if err != nil {
		logger.Error("[Handler][AuthConvertTempTokenToRealToken()]->Error CreateAuthRealToken", logger.Field("error", err.Error()))
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: http.StatusInternalServerError,
		})
		return
	}

	dataGenAccessToken, err := ctrlLogic.GenerateAuthRealToken(modelController.ParamLogicGenerateAuthRealToken{
		TokenUUID: dataCreateRealToken.AccessToken,
		TokenType: "ACCESS",
		ExpireAt:  dataCreateRealToken.ExpireAt,
	})
	if err != nil {
		logger.Error("[Handler][AuthConvertTempTokenToRealToken()]->Error GenerateAuthRealToken", logger.Field("error", err.Error()))
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: http.StatusInternalServerError,
		})
		return
	}

	dataGenRefreshToken, err := ctrlLogic.GenerateAuthRealToken(modelController.ParamLogicGenerateAuthRealToken{
		TokenUUID: dataCreateRealToken.RefreshToken,
		TokenType: "REFRESH",
		ExpireAt:  dataCreateRealToken.RefreshTokenExpireAt,
	})
	if err != nil {
		logger.Error("[Handler][AuthConvertTempTokenToRealToken()]->Error GenerateAuthRealToken", logger.Field("error", err.Error()))
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: http.StatusInternalServerError,
		})
		return
	}

	response.AccessToken = dataGenAccessToken.Token
	response.TokenType = "Bearer"
	response.ExpriesIn = 1440
	response.RefreshToken = dataGenRefreshToken.Token
	response.RefreshTokenExpriesIn = 43200

	utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
		ResponseCode: http.StatusOK,
		Message:      "Get Token Success",
		Data:         response,
	})

}

func AuthLogout(c *gin.Context) {
	err := ctrlLogic.DeleteAuthRealTokenWithAccessToken(modelController.ParamLogicDeleteAuthRealToken{
		TokenUUID: c.GetString("auth_access_token"),
	})
	if err != nil {
		logger.Error("[Handler][AuthLogout()]->Error DeleteAuthRealTokenWithAccessToken", logger.Field("error", err.Error()))
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: http.StatusInternalServerError,
		})
		return
	}

	utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
		ResponseCode: http.StatusOK,
		Message:      "Logout Success",
	})

}

func AuthRefreshToken(c *gin.Context) {
	var response modelController.ResponseAuthConvertTempTokenToRealToken

	dataCreateRealToken, err := ctrlLogic.CreateAuthRealToken(modelController.ParamLogicCreateAuthRealToken{
		AccountUUID: c.GetString("auth_account_uuid"),
		ServiceUUID: c.GetString("auth_service_uuid"),
	})
	if err != nil {
		logger.Error("[Handler][AuthRefreshToken()]->Error CreateAuthRealToken", logger.Field("error", err.Error()))
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: http.StatusInternalServerError,
		})
		return
	}

	dataGenAccessToken, err := ctrlLogic.GenerateAuthRealToken(modelController.ParamLogicGenerateAuthRealToken{
		TokenUUID: dataCreateRealToken.AccessToken,
		TokenType: "ACCESS",
		ExpireAt:  dataCreateRealToken.ExpireAt,
	})
	if err != nil {
		logger.Error("[Handler][AuthRefreshToken()]->Error GenerateAuthRealToken", logger.Field("error", err.Error()))
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: http.StatusInternalServerError,
		})
		return
	}

	dataGenRefreshToken, err := ctrlLogic.GenerateAuthRealToken(modelController.ParamLogicGenerateAuthRealToken{
		TokenUUID: dataCreateRealToken.RefreshToken,
		TokenType: "REFRESH",
		ExpireAt:  dataCreateRealToken.RefreshTokenExpireAt,
	})
	if err != nil {
		logger.Error("[Handler][AuthRefreshToken()]->Error GenerateAuthRealToken", logger.Field("error", err.Error()))
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: http.StatusInternalServerError,
		})
		return
	}

	err = ctrlLogic.DeleteAuthRealTokenWithRefreshToken(modelController.ParamLogicDeleteAuthRealToken{
		TokenUUID: c.GetString("auth_refresh_token"),
	})
	if err != nil {
		logger.Error("[Handler][AuthRefreshToken()]->Error DeleteAuthRealTokenWithRefreshToken", logger.Field("error", err.Error()))
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: http.StatusInternalServerError,
		})
		return
	}

	response.AccessToken = dataGenAccessToken.Token
	response.TokenType = "Bearer"
	response.ExpriesIn = 1440
	response.RefreshToken = dataGenRefreshToken.Token
	response.RefreshTokenExpriesIn = 43200

	utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
		ResponseCode: http.StatusOK,
		Message:      "Refresh Token Success",
		Data:         response,
	})

}