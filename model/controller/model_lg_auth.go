package modelController

import "time"

type ParamLogicCheckAuthLoginMatchPassword struct {
	AccountPassword     string
	AccountPasswordHash string
}

type ReturnLogicCheckAuthLoginMatchPassword struct {
	IsMatch bool
}

type ParamLogicCreateAuthTemporaryToken struct {
	AccountUUID string
	ServiceUUID string
}

type ReturnLogicCreateAuthTemporaryToken struct {
	TempTokenUUID string
	ExpireAt      time.Time
}

type ParamLogicGenerateAuthTemporaryToken struct {
	TempTokenUUID string
	ExpireAt      time.Time
}

type ClaimAuthTemporaryToken struct {
	TempTokenUUID string
}

type ReturnLogicGenerateAuthTemporaryToken struct {
	TempToken string
}

type ParamLogicCheckAuthTemporaryToken struct {
	TempToken string
}

type ReturnLogicCheckAuthTemporaryToken struct {
	IsValid       bool
	TempTokenUUID string
}

type ParamLogicFetchOneAuthTemporaryToken struct {
	TempTokenUUID string
}

type DBFetchLogicFetchOneAuthTemporaryToken struct {
	AuthTemporaryTokenAccountUUID       string
	AuthTemporaryTokenClientServiceUUID string
}

type ReturnLogicFetchOneAuthTemporaryToken struct {
	IsValid     bool
	AccountUUID string
	ServiceUUID string
}

type ParamLogicDeleteAuthTemporaryToken struct {
	TempTokenUUID string
}

type ParamLogicCreateAuthRealToken struct {
	AccountUUID string
	ServiceUUID string
}

type ReturnLogicCreateAuthRealToken struct {
	AccessToken          string
	ExpireAt             time.Time
	RefreshToken         string
	RefreshTokenExpireAt time.Time
}

type ParamLogicGenerateAuthRealToken struct {
	TokenUUID string
	TokenType string
	ExpireAt  time.Time
}

type ClaimAuthRealToken struct {
	TokenUUID string
}

type ReturnLogicGenerateAuthRealToken struct {
	Token string
}

type ParamLogicCheckAuthRealToken struct {
	Token     string
	TokenType string
}

type ReturnLogicCheckAuthRealToken struct {
	IsValid   bool
	TokenUUID string
}

type ParamLogicFetchOneAuthRealToken struct {
	TokenUUID string
}

type DBFetchLogicFetchOneAuthRealToken struct {
	AuthHistoryAccountUUID       string
	AuthHistoryClientServiceUUID string
}

type ReturnLogicFetchOneAuthRealToken struct {
	IsValid     bool
	AccountUUID string
	ServiceUUID string
}

type ParamLogicDeleteAuthRealToken struct {
	TokenUUID string
}

