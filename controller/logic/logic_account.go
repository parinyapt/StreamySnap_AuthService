package ctrlLogic

import (
	"github.com/parinyapt/StreamySnap_AuthService/database"
	modelController "github.com/parinyapt/StreamySnap_AuthService/model/controller"
	modelDatabase "github.com/parinyapt/StreamySnap_AuthService/model/database"
	utilsUUID "github.com/parinyapt/StreamySnap_AuthService/utils/uuid"
	PTGUpassword "github.com/parinyapt/golang_utils/password/v1"
	"github.com/pkg/errors"
)

func CreateAccount(param modelController.ParamLogicCreateAccount) (data modelController.ReturnLogicCreateAccount, err error) {
	var returnData modelController.ReturnLogicCreateAccount

	passwordHash, err := PTGUpassword.HashPassword(param.AccountPassword, 14)
	if err != nil {
		return returnData, errors.Wrap(err, "[Logic][CreateAccount()]->Fail to hash password")
	}

	accountData := modelDatabase.Account{
		UUID: utilsUUID.GenerateUUIDv4(),
		Name: param.AccountName,
		Email: param.AccountEmail,
		Password: passwordHash,
	}
	DBresult := database.DB.Create(&accountData)
	if DBresult.Error != nil {
		return returnData, errors.Wrap(DBresult.Error, "[Logic][CreateAccount()]->Fail to DB query create account")
	}

	returnData.AccountUUID = accountData.UUID
	returnData.AccountName = accountData.Name
	returnData.AccountEmail = accountData.Email

	return returnData, nil
}

func FetchOneAccount(param modelController.ParamLogicFetchOneAccount) (data modelController.ReturnLogicFetchOneAccount, err error) {
	var returnData modelController.ReturnLogicFetchOneAccount

	var DBfetchData modelController.DBFetchLogicFetchOneAccount
	DBresult := database.DB.Model(&modelDatabase.Account{}).Select("account_uuid, account_name, account_email, account_image, account_status").Where(&modelDatabase.Account{UUID: param.AccountUUID}).Or(&modelDatabase.Account{Email: param.AccountEmail}).Scan(&DBfetchData)
	if DBresult.Error != nil {
		return returnData, errors.Wrap(DBresult.Error, "[Logic][FetchOneAccount()]->Fail to DB query fetch one account")
	}
	if DBresult.RowsAffected == 1 {
		returnData.IsFound = true
	} else {
		returnData.IsFound = false
		return returnData, nil
	}

	returnData.AccountUUID = DBfetchData.AccountUUID
	returnData.AccountName = DBfetchData.AccountName
	returnData.AccountEmail = DBfetchData.AccountEmail
	returnData.AccountImage = DBfetchData.AccountImage
	returnData.AccountStatus = DBfetchData.AccountStatus

	return returnData, nil
}