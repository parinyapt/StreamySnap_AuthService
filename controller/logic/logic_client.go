package ctrlLogic

import (
	"github.com/parinyapt/StreamySnap_AuthService/logger"
	utilsUUID "github.com/parinyapt/StreamySnap_AuthService/utils/uuid"

	"github.com/parinyapt/StreamySnap_AuthService/database"
	modelDatabase "github.com/parinyapt/StreamySnap_AuthService/model/database"
	modelController "github.com/parinyapt/StreamySnap_AuthService/model/controller"
)

func CreateClient(param modelController.ParamLogicCreateClient) (data modelController.ReturnLogicCreateClient, err error) {
	var returnData modelController.ReturnLogicCreateClient

	clientData := modelDatabase.Client{
		UUID: utilsUUID.GenerateUUIDv4(),
		Name: param.ClientName,
	}
	DBresult := database.DB.Create(&clientData)
	if DBresult.Error != nil {
		logger.Error("Fail to create client", logger.Field("error", DBresult.Error.Error()))
		return returnData, DBresult.Error
	}

	returnData.UUID = clientData.UUID
	returnData.ClientName = clientData.Name

	return returnData, nil
}

func CheckExistClient(param modelController.ParamLogicCheckExistClient) (data modelController.ReturnLogicCheckExistClient, err error) {
	var returnData modelController.ReturnLogicCheckExistClient

	DBresult := database.DB.Where(&modelDatabase.Client{Name: param.ClientName}).Or(&modelDatabase.Client{UUID: param.ClientUUID}).Find(&modelDatabase.Client{})
	if DBresult.Error != nil {
		logger.Error("Fail to check exist client", logger.Field("error", DBresult.Error.Error()))
		return returnData, DBresult.Error
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
		logger.Error("Fail to delete client", logger.Field("error", DBresult.Error.Error()))
		return DBresult.Error
	}

	return nil
}
