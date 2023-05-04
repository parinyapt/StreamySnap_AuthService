package modelController

type RequestCreateClient struct {
	ClientName string `json:"name" validate:"required,max=200"`
}

type ResponseCreateClient struct {
	UUID       string `json:"client_id"`
	ClientName string `json:"name"`
}

type RequestDeleteClient struct {
	ClientUUID string `json:"client_id" validate:"required,uuid"`
}
