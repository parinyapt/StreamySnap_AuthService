package modelController

import "time"

type ParamLogicCreateAuthPageSession struct {
	ServiceUUID string
}

type ReturnLogicCreateAuthPageSession struct {
	SessionUUID string
	ExpireAt    time.Time
}

type ParamLogicGenerateAuthPageSessionURL struct {
	SessionUUID        string
	ExpireAt           time.Time
}

type ClaimAuthPageSessionURL struct {
	SessionUUID string
}

type ReturnLogicGenerateAuthPageSessionURL struct {
	AuthURL string
}

type ParamLogicCheckAuthPageSessionToken struct {
	SessionToken string
}

type ReturnLogicCheckAuthPageSessionToken struct {
	IsValid     bool
	SessionUUID string
}

type ParamLogicFetchOneAuthPageSession struct {
	SessionUUID string
}

type DBFetchLogicFetchOneAuthPageSession struct {
	AuthPageSessionClientServiceUUID string
}

type ReturnLogicFetchOneAuthPageSession struct {
	IsValid           bool
	ClientServiceUUID string
}
