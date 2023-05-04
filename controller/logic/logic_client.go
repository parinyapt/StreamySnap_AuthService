package ctrlLogic

import (
	utilsUUID "github.com/parinyapt/StreamySnap_AuthService/utils/uuid"
	"github.com/pkg/errors"

	"github.com/parinyapt/StreamySnap_AuthService/database"
	modelController "github.com/parinyapt/StreamySnap_AuthService/model/controller"
	modelDatabase "github.com/parinyapt/StreamySnap_AuthService/model/database"
)

func CreateClient(param modelController.ParamLogicCreateClient) (data modelController.ReturnLogicCreateClient, err error) {
	var returnData modelController.ReturnLogicCreateClient

	clientData := modelDatabase.Client{
		UUID: utilsUUID.GenerateUUIDv4(),
		Name: param.ClientName,
	}
	DBresult := database.DB.Create(&clientData)
	if DBresult.Error != nil {
		return returnData, errors.Wrap(DBresult.Error, "[Logic][CreateClient()]->Fail to DB query create client")
	}

	returnData.UUID = clientData.UUID
	returnData.ClientName = clientData.Name

	return returnData, nil
}

func CheckExistClient(param modelController.ParamLogicCheckExistClient) (data modelController.ReturnLogicCheckExistClient, err error) {
	var returnData modelController.ReturnLogicCheckExistClient

	DBresult := database.DB.Where(&modelDatabase.Client{Name: param.ClientName}).Or(&modelDatabase.Client{UUID: param.ClientUUID}).Find(&modelDatabase.Client{})
	if DBresult.Error != nil {
		return returnData, errors.Wrap(DBresult.Error, "[Logic][CheckExistClient()]->Fail to DB query check exist client")
	}

	if DBresult.RowsAffected == 0 {
		returnData.IsExist = false
	} else {
		returnData.IsExist = true
	}

	return returnData, DBresult.Error
}

func DeleteClient(param modelController.ParamLogicDeleteClient) (err error) {
	DBresult := database.DB.Where(&modelDatabase.Client{UUID: param.ClientUUID}).Delete(&modelDatabase.Client{})
	if DBresult.Error != nil {
		return errors.Wrap(DBresult.Error, "[Logic][DeleteClient()]->Fail to DB query delete client")
	}

	return nil
}
