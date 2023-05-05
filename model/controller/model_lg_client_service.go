package modelController

type ParamLogicCreateClientService struct {
	ClientUUID         string
	ServiceName        string
	ServiceCallbackURL string
}

type ReturnLogicCreateClientService struct {
	ServiceName      string
	ServiceUUID      string
	ServiceSecretKey string
}

type ParamLogicFetchOneClientService struct {
	ServiceUUID string
}

type DBFetchLogicFetchOneClientService struct {
	ClientServiceUUID        string
	ClientServiceName        string
	ClientServiceStatus      string
	ClientServiceCallbackURL string
	ClientUUID               string
	ClientName               string
}

type ReturnLogicFetchOneClientService struct {
	IsFound                  bool
	ClientServiceUUID        string
	ClientServiceName        string
	ClientServiceStatus      string
	ClientServiceCallbackURL string
	ClientUUID               string
	ClientName               string
}
