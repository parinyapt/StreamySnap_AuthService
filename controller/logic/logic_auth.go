package ctrlLogic

import (
	"os"
	"time"

	PTGUjwt "github.com/parinyapt/golang_utils/jwt/v1"
	PTGUpassword "github.com/parinyapt/golang_utils/password/v1"
	"github.com/pkg/errors"

	"github.com/parinyapt/StreamySnap_AuthService/database"
	"github.com/parinyapt/StreamySnap_AuthService/logger"
	modelController "github.com/parinyapt/StreamySnap_AuthService/model/controller"
	modelDatabase "github.com/parinyapt/StreamySnap_AuthService/model/database"
	utilsUUID "github.com/parinyapt/StreamySnap_AuthService/utils/uuid"
)

func CheckAuthLoginMatchPassword(param modelController.ParamLogicCheckAuthLoginMatchPassword) (data modelController.ReturnLogicCheckAuthLoginMatchPassword, err error) {
	var returnData modelController.ReturnLogicCheckAuthLoginMatchPassword

	returnData.IsMatch = PTGUpassword.VerifyHashPassword(param.AccountPassword, param.AccountPasswordHash)

	return returnData, nil
}

func CreateAuthTemporaryToken(param modelController.ParamLogicCreateAuthTemporaryToken) (data modelController.ReturnLogicCreateAuthTemporaryToken, err error) {
	var returnData modelController.ReturnLogicCreateAuthTemporaryToken

	tempTokenData := modelDatabase.AuthTemporaryToken{
		UUID:              utilsUUID.GenerateUUIDv4(),
		AccountUUID:       param.AccountUUID,
		ClientServiceUUID: param.ServiceUUID,
		ExpiredAt:         time.Now().Add(time.Minute * 5),
	}
	DBresult := database.DB.Create(&tempTokenData)
	if DBresult.Error != nil {
		return returnData, errors.Wrap(DBresult.Error, "[Logic][CreateAuthTemporaryToken()]->Fail to DB query create auth temporary token")
	}

	returnData.TempTokenUUID = tempTokenData.UUID
	returnData.ExpireAt = tempTokenData.ExpiredAt

	return returnData, nil
}

func GenerateAuthTemporaryToken(param modelController.ParamLogicGenerateAuthTemporaryToken) (data modelController.ReturnLogicGenerateAuthTemporaryToken, err error) {
	var returnData modelController.ReturnLogicGenerateAuthTemporaryToken

	claim := modelController.ClaimAuthTemporaryToken{
		TempTokenUUID: param.TempTokenUUID,
	}
	token, err := PTGUjwt.Sign(PTGUjwt.JwtSignConfig{
		SignKey:       os.Getenv("JWT_SIGN_KEY_TEMPORARYTOKEN"),
		AppName:       os.Getenv("APP_NAME"),
		ExpireTime:    param.ExpireAt,
		IssuedTime:    time.Now(),
		NotBeforeTime: time.Now(),
	}, claim)
	if err != nil {
		return returnData, errors.Wrap(err, "[Logic][GenerateAuthTemporaryToken()]->Fail to generate jwt token")
	}

	returnData.TempToken = token

	return returnData, nil
}

func CheckAuthTemporaryToken(param modelController.ParamLogicCheckAuthTemporaryToken) (data modelController.ReturnLogicCheckAuthTemporaryToken, err error) {
	var returnData modelController.ReturnLogicCheckAuthTemporaryToken

	claims, isExpireOrNotValidYet, err := PTGUjwt.Validate(param.TempToken, PTGUjwt.JwtValidateConfig{
		SignKey: os.Getenv("JWT_SIGN_KEY_TEMPORARYTOKEN"),
	})
	if err != nil {
		returnData.IsValid = false
		return returnData, errors.Wrap(err, "[Logic][CheckAuthTemporaryToken()]->Fail to validate jwt token")
	}
	if isExpireOrNotValidYet {
		returnData.IsValid = false
		return returnData, nil
	}

	returnData.IsValid = true
	returnData.TempTokenUUID = claims.(map[string]interface{})["TempTokenUUID"].(string)

	return returnData, nil
}

func FetchOneAuthTemporaryToken(param modelController.ParamLogicFetchOneAuthTemporaryToken) (data modelController.ReturnLogicFetchOneAuthTemporaryToken, err error) {
	var returnData modelController.ReturnLogicFetchOneAuthTemporaryToken

	var DBfetchData modelController.DBFetchLogicFetchOneAuthTemporaryToken
	DBresult := database.DB.Model(&modelDatabase.AuthTemporaryToken{}).Select("auth_temporary_token_account_uuid,auth_temporary_token_client_service_uuid").Where("auth_temporary_token_uuid = ? AND auth_temporary_token_expired_at >= ?", param.TempTokenUUID, time.Now()).Scan(&DBfetchData)
	if DBresult.Error != nil {
		return returnData, errors.Wrap(DBresult.Error, "[Logic][FetchOneAuthPageSession()]->Fail to DB query fetch auth page session")
	}

	if DBresult.RowsAffected == 1 {
		returnData.IsValid = true
	} else {
		returnData.IsValid = false
		return returnData, nil
	}

	returnData.AccountUUID = DBfetchData.AuthTemporaryTokenAccountUUID
	returnData.ServiceUUID = DBfetchData.AuthTemporaryTokenClientServiceUUID

	return returnData, nil
}

func DeleteAuthTemporaryToken(param modelController.ParamLogicDeleteAuthTemporaryToken) (err error) {
	DBresult := database.DB.Where("auth_temporary_token_uuid = ?", param.TempTokenUUID).Delete(&modelDatabase.AuthTemporaryToken{})
	if DBresult.Error != nil {
		return errors.Wrap(DBresult.Error, "[Logic][DeleteAuthTemporaryToken()]->Fail to DB query delete auth temporary token")
	}

	return nil
}

func CreateAuthRealToken(param modelController.ParamLogicCreateAuthRealToken) (data modelController.ReturnLogicCreateAuthRealToken, err error) {
	var returnData modelController.ReturnLogicCreateAuthRealToken

	realTokenData := modelDatabase.AuthHistory{
		AccountUUID:        param.AccountUUID,
		ClientServiceUUID:  param.ServiceUUID,
		UUID:               utilsUUID.GenerateUUIDv4(),
		ExpiredAt:          time.Now().Add(time.Minute * 1440),
		RefreshToken:       utilsUUID.GenerateUUIDv4(),
		RefreshTokenExpire: time.Now().Add(time.Hour * 24 * 30),
	}
	DBresult := database.DB.Create(&realTokenData)
	if DBresult.Error != nil {
		return returnData, errors.Wrap(DBresult.Error, "[Logic][CreateAuthTemporaryToken()]->Fail to DB query create auth temporary token")
	}

	returnData.AccessToken = realTokenData.UUID
	returnData.ExpireAt = realTokenData.ExpiredAt
	returnData.RefreshToken = realTokenData.RefreshToken
	returnData.RefreshTokenExpireAt = realTokenData.RefreshTokenExpire

	return returnData, nil
}

func GenerateAuthRealToken(param modelController.ParamLogicGenerateAuthRealToken) (data modelController.ReturnLogicGenerateAuthRealToken, err error) {
	var returnData modelController.ReturnLogicGenerateAuthRealToken

	claim := modelController.ClaimAuthRealToken{
		TokenUUID: param.TokenUUID,
	}
	token, err := PTGUjwt.Sign(PTGUjwt.JwtSignConfig{
		SignKey:       os.Getenv("JWT_SIGN_KEY_" + param.TokenType + "TOKEN"),
		AppName:       os.Getenv("APP_NAME"),
		ExpireTime:    param.ExpireAt,
		IssuedTime:    time.Now(),
		NotBeforeTime: time.Now(),
	}, claim)
	if err != nil {
		return returnData, errors.Wrap(err, "[Logic][GenerateAuthRealToken()]->Fail to generate jwt token")
	}

	returnData.Token = token

	return returnData, nil
}

func CheckAuthRealToken(param modelController.ParamLogicCheckAuthRealToken) (data modelController.ReturnLogicCheckAuthRealToken, err error) {
	var returnData modelController.ReturnLogicCheckAuthRealToken

	claims, isExpireOrNotValidYet, err := PTGUjwt.Validate(param.Token, PTGUjwt.JwtValidateConfig{
		SignKey: os.Getenv("JWT_SIGN_KEY_" + param.TokenType + "TOKEN"),
	})
	if err != nil {
		returnData.IsValid = false
		return returnData, errors.Wrap(err, "[Logic][CheckAuthRealToken()]->Fail to validate jwt token")
	}
	if isExpireOrNotValidYet {
		returnData.IsValid = false
		return returnData, nil
	}

	returnData.IsValid = true
	returnData.TokenUUID = claims.(map[string]interface{})["TokenUUID"].(string)

	return returnData, nil
}

func FetchOneAuthAccessToken(param modelController.ParamLogicFetchOneAuthRealToken) (data modelController.ReturnLogicFetchOneAuthRealToken, err error) {
	var returnData modelController.ReturnLogicFetchOneAuthRealToken

	var DBfetchData modelController.DBFetchLogicFetchOneAuthRealToken
	DBresult := database.DB.Model(&modelDatabase.AuthHistory{}).Select("auth_history_account_uuid,auth_history_client_service_uuid").Where("auth_history_uuid = ? AND auth_history_expired_at >= ?", param.TokenUUID, time.Now()).Scan(&DBfetchData)
	if DBresult.Error != nil {
		return returnData, errors.Wrap(DBresult.Error, "[Logic][FetchOneAuthAccessToken()]->Fail to DB query fetch auth page session")
	}

	if DBresult.RowsAffected == 1 {
		returnData.IsValid = true
	} else {
		returnData.IsValid = false
		return returnData, nil
	}

	returnData.AccountUUID = DBfetchData.AuthHistoryAccountUUID
	returnData.ServiceUUID = DBfetchData.AuthHistoryClientServiceUUID

	return returnData, nil
}

func FetchOneAuthRefreshToken(param modelController.ParamLogicFetchOneAuthRealToken) (data modelController.ReturnLogicFetchOneAuthRealToken, err error) {
	var returnData modelController.ReturnLogicFetchOneAuthRealToken

	var DBfetchData modelController.DBFetchLogicFetchOneAuthRealToken
	DBresult := database.DB.Model(&modelDatabase.AuthHistory{}).Select("auth_history_account_uuid,auth_history_client_service_uuid").Where("auth_history_refresh_token = ? AND auth_history_refresh_token_expire >= ?", param.TokenUUID, time.Now()).Scan(&DBfetchData)
	if DBresult.Error != nil {
		return returnData, errors.Wrap(DBresult.Error, "[Logic][FetchOneAuthRefreshToken()]->Fail to DB query fetch auth refresh token")
	}

	if DBresult.RowsAffected == 1 {
		returnData.IsValid = true
	} else {
		returnData.IsValid = false
		return returnData, nil
	}

	returnData.AccountUUID = DBfetchData.AuthHistoryAccountUUID
	returnData.ServiceUUID = DBfetchData.AuthHistoryClientServiceUUID

	return returnData, nil
}

func DeleteAuthRealTokenWithAccessToken(param modelController.ParamLogicDeleteAuthRealToken) (err error) {
	DBresult := database.DB.Where("auth_history_uuid = ?", param.TokenUUID).Delete(&modelDatabase.AuthHistory{})
	if DBresult.Error != nil {
		return errors.Wrap(DBresult.Error, "[Logic][DeleteAuthRealToken()]->Fail to DB query delete auth access token")
	}

	return nil
}

func DeleteAuthRealTokenWithRefreshToken(param modelController.ParamLogicDeleteAuthRealToken) (err error) {
	DBresult := database.DB.Where("auth_history_refresh_token = ?", param.TokenUUID).Delete(&modelDatabase.AuthHistory{})
	if DBresult.Error != nil {
		return errors.Wrap(DBresult.Error, "[Logic][DeleteAuthRealToken()]->Fail to DB query delete auth refresh token")
	}

	logger.Debug("RowsAffected",logger.Field("row",DBresult.RowsAffected))

	return nil
}
