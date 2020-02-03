package defs

type Err struct {
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"`
}

type ErroResponse struct {
	HTTPSC int
	Error  Err
}

var (
	ErrorRequestBodyParseFailed = ErroResponse{HTTPSC: 400, Error: Err{Error: "Requests body is not correct", ErrorCode: "001"}}
	ErrorNotAuthUser            = ErroResponse{HTTPSC: 401, Error: Err{Error:"User authentication failed.", ErrorCode: "002"}}
)
