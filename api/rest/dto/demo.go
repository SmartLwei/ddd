package dto

type JSONResult struct {
	Code    int64 `json:"code"`
	Data    interface{}
	Message string `json:"message"`
}

type DemoReq struct {
	ID int64 `json:"id" form:"id" example:"1"`
}

type DemoResp struct {
	ID int64 `json:"id"`
}
