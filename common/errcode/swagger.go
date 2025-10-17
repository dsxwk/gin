package errcode

// SuccessResponse swagger 文档
type SuccessResponse struct {
	Code int64    `json:"code" example:"0"`
	Msg  string   `json:"msg" example:"Success"`
	Data []string `json:"data"`
}

// ArgsErrorResponse swagger 文档
type ArgsErrorResponse struct {
	Code int64    `json:"code" example:"400"`
	Msg  string   `json:"msg" example:"Invalid arguments"`
	Data []string `json:"data"`
}

// AuthErrorResponse swagger 文档
type AuthErrorResponse struct {
	Code int64    `json:"code" example:"401"`
	Msg  string   `json:"msg" example:"Unauthorized"`
	Data []string `json:"data"`
}

// SystemErrorResponse swagger 文档
type SystemErrorResponse struct {
	Code int64    `json:"code" example:"500"`
	Msg  string   `json:"msg" example:"Internal server error"`
	Data []string `json:"data"`
}
