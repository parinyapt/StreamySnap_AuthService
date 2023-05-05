package modelController

type RequestCreateAccount struct {
	AccountName     string `json:"name" validate:"required,max=200"`
	AccountEmail    string `json:"email" validate:"required,email"`
	AccountPassword string `json:"password" validate:"required,min=8,max=72"`
}

type ResponseCreateAccount struct {
	AccountUUID  string `json:"account_id"`
	AccountName  string `json:"name"`
	AccountEmail string `json:"email"`
}
