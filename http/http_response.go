package http

// HandlerResponse
// @Description: 响应结构体
type HandlerResponse struct {
	Status int         `json:"status"    dc:"Error code"`
	Msg    string      `json:"msg" dc:"Error message"`
	Remark string      `json:"remark" dc:"client tip message"`
	Data   interface{} `json:"data"    dc:"Result data for certain request according API definition"`
}
