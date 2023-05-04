package modelController

type RequestCreateClient struct {
	ClientName string `json:"client_name" validate:"required,max=200"`
}

type RequestDeleteClient struct {
	ClientUUID string `json:"client_uuid" validate:"required,uuid"`
}