package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/parinyapt/StreamySnap_AuthService/logger"
	modelUtils "github.com/parinyapt/StreamySnap_AuthService/model/utils"
	utilsHeader "github.com/parinyapt/StreamySnap_AuthService/utils/header"
	utilsResponse "github.com/parinyapt/StreamySnap_AuthService/utils/response"
	PTGUvalidator "github.com/parinyapt/golang_utils/validator/v1"
	ctrlLogic "github.com/parinyapt/StreamySnap_AuthService/controller/logic"
	modelController "github.com/parinyapt/StreamySnap_AuthService/model/controller"
)

type AuthorizationToken struct {
	BearerToken string `validate:"required,jwt" json:"bearer_token"`
}

func CheckAuthAccessToken(c *gin.Context) {
	authToken, err := utilsHeader.GetHeaderAuthorizationValue(c)
	if err != nil {
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: http.StatusUnauthorized,
			Message:      "Authorization Header Not Found",
		})
		c.Abort()
		return
	}

	var authorizationToken AuthorizationToken
	authorizationToken.BearerToken = authToken

	isValidatePass, _, validatorError := PTGUvalidator.Validate(authorizationToken)
	if validatorError != nil {
		logger.Error("[Middleware][CheckAuthAccessToken()]->Error Validate", logger.Field("error", validatorError.Error()))
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: http.StatusInternalServerError,
		})
		c.Abort()
		return
	}
	if !isValidatePass {
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: http.StatusBadRequest,
			Message: "Authorization Header Not Valid",
		})
		c.Abort()
		return
	}

	checkTokenData, err := ctrlLogic.CheckAuthRealToken(modelController.ParamLogicCheckAuthRealToken{
		Token: authToken,
		TokenType: "ACCESS",
	})
	if err != nil {
		logger.Warning("[Middleware][CheckAuthAccessToken()]->Error CheckAuthRealToken", logger.Field("error", err.Error()))
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: http.StatusBadRequest,
		})
		c.Abort()
		return
	}
	if !checkTokenData.IsValid {
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: http.StatusUnauthorized,
			Message:      "Session Not Found",
		})
		c.Abort()
		return
	}

	fetchDataToken, err := ctrlLogic.FetchOneAuthAccessToken(modelController.ParamLogicFetchOneAuthRealToken{TokenUUID: checkTokenData.TokenUUID})
	if err != nil {
		logger.Error("[Middleware][CheckAuthAccessToken()]->Error FetchOneAuthAccessToken", logger.Field("error", err.Error()))
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: http.StatusInternalServerError,
		})
		c.Abort()
		return
	}
	if !fetchDataToken.IsValid {
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: http.StatusUnauthorized,
		})
		c.Abort()
		return
	}

	c.Set("auth_access_token", checkTokenData.TokenUUID)
	c.Set("auth_account_uuid", fetchDataToken.AccountUUID)
	c.Set("auth_service_uuid", fetchDataToken.ServiceUUID)
	c.Next()
}

func CheckAuthRefreshToken(c *gin.Context) {
	authToken, err := utilsHeader.GetHeaderAuthorizationValue(c)
	if err != nil {
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: http.StatusUnauthorized,
			Message:      "Authorization Header Not Found",
		})
		c.Abort()
		return
	}

	var authorizationToken AuthorizationToken
	authorizationToken.BearerToken = authToken

	isValidatePass, _, validatorError := PTGUvalidator.Validate(authorizationToken)
	if validatorError != nil {
		logger.Error("[Middleware][CheckAuthRefreshToken()]->Error Validate", logger.Field("error", validatorError.Error()))
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: http.StatusInternalServerError,
		})
		c.Abort()
		return
	}
	if !isValidatePass {
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: http.StatusBadRequest,
			Message: "Authorization Header Not Valid",
		})
		c.Abort()
		return
	}

	checkTokenData, err := ctrlLogic.CheckAuthRealToken(modelController.ParamLogicCheckAuthRealToken{
		Token: authToken,
		TokenType: "REFRESH",
	})
	if err != nil {
		logger.Warning("[Middleware][CheckAuthRefreshToken()]->Error CheckAuthRealToken", logger.Field("error", err.Error()))
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: http.StatusBadRequest,
		})
		c.Abort()
		return
	}
	if !checkTokenData.IsValid {
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: http.StatusUnauthorized,
			Message:      "Session Not Found",
		})
		c.Abort()
		return
	}

	fetchDataToken, err := ctrlLogic.FetchOneAuthRefreshToken(modelController.ParamLogicFetchOneAuthRealToken{TokenUUID: checkTokenData.TokenUUID})
	if err != nil {
		logger.Error("[Middleware][CheckAuthRefreshToken()]->Error FetchOneAuthRefreshToken", logger.Field("error", err.Error()))
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: http.StatusInternalServerError,
		})
		c.Abort()
		return
	}
	if !fetchDataToken.IsValid {
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: http.StatusUnauthorized,
		})
		c.Abort()
		return
	}

	c.Set("auth_refresh_token", checkTokenData.TokenUUID)
	c.Set("auth_account_uuid", fetchDataToken.AccountUUID)
	c.Set("auth_service_uuid", fetchDataToken.ServiceUUID)
	c.Next()
}