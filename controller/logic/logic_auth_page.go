package ctrlLogic

import (
	"os"
	"time"

	"github.com/pkg/errors"

	PTGUjwt "github.com/parinyapt/golang_utils/jwt/v1"

	"github.com/parinyapt/StreamySnap_AuthService/database"
	modelController "github.com/parinyapt/StreamySnap_AuthService/model/controller"
	modelDatabase "github.com/parinyapt/StreamySnap_AuthService/model/database"
	utilsUUID "github.com/parinyapt/StreamySnap_AuthService/utils/uuid"
)

func CreateAuthPageSession(param modelController.ParamLogicCreateAuthPageSession) (data modelController.ReturnLogicCreateAuthPageSession, err error) {
	var returnData modelController.ReturnLogicCreateAuthPageSession

	authPageSessionData := modelDatabase.AuthPageSession{
		UUID:              utilsUUID.GenerateUUIDv4(),
		ClientServiceUUID: param.ServiceUUID,
		ExpiredAt:         time.Now().Add(time.Minute * 5),
	}
	DBresult := database.DB.Create(&authPageSessionData)
	if DBresult.Error != nil {
		return returnData, errors.Wrap(DBresult.Error, "[Logic][CreateAuthPageSession()]->Fail to DB query create auth page session")
	}

	returnData.SessionUUID = authPageSessionData.UUID
	returnData.ExpireAt = authPageSessionData.ExpiredAt

	return returnData, nil
}

func GenerateAuthPageSessionURL(param modelController.ParamLogicGenerateAuthPageSessionURL) (data modelController.ReturnLogicGenerateAuthPageSessionURL, err error) {
	var returnData modelController.ReturnLogicGenerateAuthPageSessionURL

	claim := modelController.ClaimAuthPageSessionURL{
		SessionUUID: param.SessionUUID,
	}
	token, err := PTGUjwt.Sign(PTGUjwt.JwtSignConfig{
		SignKey:       os.Getenv("JWT_SIGN_KEY_AUTHSESSION"),
		AppName:       os.Getenv("APP_NAME"),
		ExpireTime:    param.ExpireAt,
		IssuedTime:    time.Now(),
		NotBeforeTime: time.Now(),
	}, claim)
	if err != nil {
		return returnData, errors.Wrap(err, "[Logic][GenerateAuthPageSessionURL()]->Fail to generate jwt token")
	}

	returnData.AuthURL = os.Getenv("APP_BASE_URL") + "?session=" + token

	return returnData, nil
}

func CheckAuthPageSessionToken(param modelController.ParamLogicCheckAuthPageSessionToken) (data modelController.ReturnLogicCheckAuthPageSessionToken, err error) {
	var returnData modelController.ReturnLogicCheckAuthPageSessionToken

	claims, isExpireOrNotValidYet, err := PTGUjwt.Validate(param.SessionToken, PTGUjwt.JwtValidateConfig{
		SignKey:    os.Getenv("JWT_SIGN_KEY_AUTHSESSION"),
	})
	if err != nil {
		returnData.IsValid = false
		return returnData, errors.Wrap(err, "[Logic][CheckAuthPageSession()]->Fail to validate jwt token")
	}
	if isExpireOrNotValidYet {
		returnData.IsValid = false
		return returnData, nil
	}

	returnData.IsValid = true
	returnData.SessionUUID = claims.(map[string]interface{})["SessionUUID"].(string)

	return returnData, nil
}

func FetchOneAuthPageSession(param modelController.ParamLogicFetchOneAuthPageSession) (data modelController.ReturnLogicFetchOneAuthPageSession, err error) {
	var returnData modelController.ReturnLogicFetchOneAuthPageSession

	var DBfetchData modelController.DBFetchLogicFetchOneAuthPageSession
	DBresult := database.DB.Model(&modelDatabase.AuthPageSession{}).Select("auth_page_session_client_service_uuid").Where("auth_page_session_uuid = ? AND auth_page_session_expired_at >= ?", param.SessionUUID, time.Now()).Scan(&DBfetchData)
	if DBresult.Error != nil {
		return returnData, errors.Wrap(DBresult.Error, "[Logic][FetchOneAuthPageSession()]->Fail to DB query fetch auth page session")
	}

	if DBresult.RowsAffected == 1 {
		returnData.IsValid = true
	} else {
		returnData.IsValid = false
		return returnData, nil
	}

	returnData.ClientServiceUUID = DBfetchData.AuthPageSessionClientServiceUUID

	return returnData, nil
}
