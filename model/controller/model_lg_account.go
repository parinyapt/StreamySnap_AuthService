package modelController

type ParamLogicCreateAccount struct {
	AccountName     string
	AccountEmail    string
	AccountPassword string
}

type ReturnLogicCreateAccount struct {
	AccountUUID  string
	AccountName  string
	AccountEmail string
}

type ParamLogicFetchOneAccount struct {
	AccountUUID  string
	AccountEmail string
}

type DBFetchLogicFetchOneAccount struct {
	AccountUUID     string
	AccountName     string
	AccountEmail    string
	AccountImage    string
	AccountStatus   string
	AccountPassword string
}

type ReturnLogicFetchOneAccount struct {
	IsFound         bool
	AccountUUID     string
	AccountName     string
	AccountEmail    string
	AccountImage    string
	AccountStatus   string
	AccountPassword string
}
