package modelController

type ResponseGetProfile struct {
	Account         ResponseGetProfileAccount `json:"account"`
	Service         ResponseGetProfileService `json:"service"`
	AuthCallbackURL string                    `json:"callback_url"`
}

type ResponseGetProfileAccount struct {
	AccountUUID  string `json:"account_id"`
	AccountName  string `json:"account_name"`
	AccountEmail string `json:"account_email"`
}

type ResponseGetProfileService struct {
	ServiceUUID string `json:"service_id"`
	ServiceName string `json:"service_name"`
}
