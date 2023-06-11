package utils

// Error code

type ErrCode struct {
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"`
}

type ErrorResponse struct {
	HttpSc int
	Error  ErrCode
}

var (
	ErrorRequestBodyParseFailed = ErrorResponse{HttpSc: 400, Error: ErrCode{Error: "Request body is not correct", ErrorCode: "001"}}
	ErrorNotAuthUser            = ErrorResponse{HttpSc: 401, Error: ErrCode{Error: "user authentication failed", ErrorCode: "002"}}
)
