package response

type Response struct {
	Code      int    `json:"code"`
	RequestId string `json:"request_id"`
	Data      any    `json:"data"`
	Message   string `json:"message"`
}
