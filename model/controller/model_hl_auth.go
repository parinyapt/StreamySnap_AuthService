package modelController

type RequestAuthLoginWithUsernamePassword struct {
	AccountEmail    string `json:"email" validate:"required,email"`
	AccountPassword string `json:"password" validate:"required"`
}

type ResponseAuthLoginWithUsernamePassword struct {
	AuthCallbackURL string `json:"callback_url"`
}

type RequestAuthConvertTempTokenToRealToken struct {
	TempToken string `json:"token" validate:"required,jwt"`
}

type ResponseAuthConvertTempTokenToRealToken struct {
	AccessToken           string `json:"access_token"`
	TokenType             string `json:"token_type"`
	ExpriesIn             int    `json:"expries_in"`
	RefreshToken          string `json:"refresh_token"`
	RefreshTokenExpriesIn int    `json:"refresh_token_expries_in"`
}
