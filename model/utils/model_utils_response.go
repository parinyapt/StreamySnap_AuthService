package modelUtils

type JsonResponseStruct struct {
	ResponseCode int
	Detail       JsonResponseStructDetail
}

type JsonResponseStructDetail struct {
	Timestamp int64       `json:"timestamp"`
	Success   bool        `json:"success"`
	Message   string      `json:"message"`
	ErrorCode string      `json:"error_code"`
	Data      interface{} `json:"data"`
}

type ApiResponseStruct struct {
	ResponseCode int
	Message      string
	ErrorCode    string
	Data         interface{}
}

type ApiResponseConfigStruct struct {
	Message   string
	ErrorCode string
}
