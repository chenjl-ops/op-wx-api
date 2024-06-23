package response

type Response struct {
	Code      int         `json:"code"`
	RequestId string      `json:"request_id"`
	Data      interface{} `json:"data"`
	Message   string      `json:"message"`
}
