package response

type Response struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Remark string      `json:"remark"`
	Data   interface{} `json:"data"`
}
