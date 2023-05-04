package modelController

type ParamLogicCreateClient struct {
	ClientName string
}

type ReturnLogicCreateClient struct {
	UUID       string
	ClientName string
}

type ParamLogicCheckExistClient struct {
	ClientName string
	ClientUUID string
}

type ReturnLogicCheckExistClient struct {
	IsExist bool
}

type ParamLogicDeleteClient struct {
	ClientUUID string
}