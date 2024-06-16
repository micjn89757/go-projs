package utils 


// 定义响应返回的json格式
type Ecode struct {
	Status int `json:"status"`
	Msg string `json:"msg"`
	Data any	`json:"data"`
	ID int 		`json:"id"`
	TOKEN string `json:"token"`
}

