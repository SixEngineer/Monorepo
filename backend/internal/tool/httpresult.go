package tool

type HttpResult struct {
	Code    int         `json:"code"` // 状态码
	Message interface{} `json:"msg"`  // 消息
	Data    interface{} `json:"data"` // 数据
}

func (h HttpResult) Success(data interface{}) HttpResult {
	return HttpResult{
		Code:    1,
		Message: nil,
		Data:    data,
	}
}