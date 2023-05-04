package ctrlLogic

import (
	"github.com/parinyapt/golang_utils/password/v1"
	"github.com/pkg/errors"

	"github.com/parinyapt/StreamySnap_AuthService/database"
	modelController "github.com/parinyapt/StreamySnap_AuthService/model/controller"
	modelDatabase "github.com/parinyapt/StreamySnap_AuthService/model/database"
	utilsUUID "github.com/parinyapt/StreamySnap_AuthService/utils/uuid"
)

func CreateClientService(param modelController.ParamLogicCreateClientService) (data modelController.ReturnLogicCreateClientService, err error) {
	var returnData modelController.ReturnLogicCreateClientService

	secretkey := utilsUUID.GenerateUUIDv4() + utilsUUID.GenerateUUIDv4()
	secretkeyHash, err := PTGUpassword.HashPassword(secretkey, 14)
	if err != nil {
		return returnData, errors.Wrap(err, "[Logic][CreateClientService()]->Fail to hash secret key")
	}

	clientServiceData := modelDatabase.ClientService{
		UUID: utilsUUID.GenerateUUIDv4(),
		SecretKey: secretkeyHash,
		ClientUUID: param.ClientUUID,
		Name: param.ServiceName,
		CallBackURL: param.ServiceCallBackURL,
	}
	DBresult := database.DB.Create(&clientServiceData)
	if DBresult.Error != nil {
		return returnData, errors.Wrap(DBresult.Error, "[Logic][CreateClientService()]->Fail to DB query create client service")
	}

	returnData.ServiceName = clientServiceData.Name
	returnData.ServiceUUID = clientServiceData.UUID
	returnData.ServiceSecretKey = secretkey

	return returnData, nil
}