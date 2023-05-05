package modelController

type RequestGenerateAuthPageSessionURL struct {
	ServiceUUID      string `json:"service_id" validate:"required,uuid"`
	ServiceSecretKey string `json:"service_secret_key" validate:"required,len=72"`
}

type ResponseGenerateAuthPageSessionURL struct {
	AuthPageSessionURL string `json:"url"`
}

type RequestCheckAuthPageSessionToken struct {
	SessionToken string `json:"session_token" validate:"required,jwt"`
}

type ResponseCheckAuthPageSessionToken struct {
	ClientName        string `json:"client_name"`
	ClientServiceName string `json:"service_name"`
}
