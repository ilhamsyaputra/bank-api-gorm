package response

type Response struct {
	Code   int
	Status string
	Remark string
	Data   interface{}
}
