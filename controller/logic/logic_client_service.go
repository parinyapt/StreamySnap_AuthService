package ctrlLogic

import (
	PTGUpassword "github.com/parinyapt/golang_utils/password/v1"
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
		UUID:        utilsUUID.GenerateUUIDv4(),
		SecretKey:   secretkeyHash,
		ClientUUID:  param.ClientUUID,
		Name:        param.ServiceName,
		CallbackURL: param.ServiceCallbackURL,
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

func FetchOneClientService(param modelController.ParamLogicFetchOneClientService) (data modelController.ReturnLogicFetchOneClientService, err error) {
	var returnData modelController.ReturnLogicFetchOneClientService

	var DBfetchData modelController.DBFetchLogicFetchOneClientService
	DBresult := database.DB.Model(&modelDatabase.ClientService{}).Select("authservice_client_service.client_service_uuid, authservice_client_service.client_service_name, authservice_client_service.client_service_status, authservice_client_service.client_service_callback_url, authservice_client.client_uuid, authservice_client.client_name").Joins("INNER JOIN authservice_client ON authservice_client_service.client_service_client_uuid = authservice_client.client_uuid").Where("client_service_uuid = ?", param.ServiceUUID).Scan(&DBfetchData)
	if DBresult.Error != nil {
		return returnData, errors.Wrap(DBresult.Error, "[Logic][FetchOneClientService()]->Fail to DB query fetch one client service")
	}

	if DBresult.RowsAffected == 1 {
		returnData.IsFound = true
	} else {
		returnData.IsFound = false
		return returnData, nil
	}

	returnData.ClientServiceUUID = DBfetchData.ClientServiceUUID
	returnData.ClientServiceName = DBfetchData.ClientServiceName
	returnData.ClientServiceStatus = DBfetchData.ClientServiceStatus
	returnData.ClientServiceCallbackURL = DBfetchData.ClientServiceCallbackURL
	returnData.ClientUUID = DBfetchData.ClientUUID
	returnData.ClientName = DBfetchData.ClientName

	return returnData, nil
}
