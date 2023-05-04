package modelController

type ParamLogicCreateClientService struct {
	ClientUUID         string
	ServiceName        string
	ServiceCallBackURL string
}

type ReturnLogicCreateClientService struct {
	ServiceName        string
	ServiceUUID        string
	ServiceSecretKey string
}
