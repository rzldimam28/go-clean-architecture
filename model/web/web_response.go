package web

type WebResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func CreateWebResponse(code int, msg string, data interface{}) WebResponse {
	return WebResponse{
		Code:   code,
		Status: msg,
		Data:   data,
	}
}