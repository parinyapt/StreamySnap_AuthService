package modelController

type RequestCreateClientService struct {
	ClientUUID         string `json:"client_id" validate:"required,uuid"`
	ServiceName        string `json:"name" validate:"required,max=200"`
	ServiceCallBackURL string `json:"callback_url" validate:"required,url"`
}

type ResponseCreateClientService struct {
	ServiceName            string `json:"name"`
	ServiceUUID      string `json:"service_id"`
	ServiceSecretKey string `json:"service_secret_key"`
}
