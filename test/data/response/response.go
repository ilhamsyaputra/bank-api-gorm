package response

type LoginResponseData struct {
	Token string `json:"Token"`
}

type LoginResponse struct {
	Code   int               `json:"Code"`
	Status string            `json:"Status"`
	Remark string            `json:"Remark"`
	Data   LoginResponseData `json:"Data"`
}

type TabungResponseData struct {
	Saldo float64 `json:"Saldo"`
}

type TabungResponse struct {
	Code   int    `json:"code"`
	Remark string `json:"remark"`
	Status string `json:"status"`
}

type TabungResponseSuccess struct {
	Code   int                `json:"Code"`
	Status string             `json:"Status"`
	Remark string             `json:"Remark"`
	Data   TabungResponseData `json:"Data"`
}
