package handler

const (
	SuccessMessage = "success!"
)

type defaultResponse struct {
	Code     int         `json:"code"`
	Message  string      `json:"message"`
	Error    string      `json:"error,omitempty"`
	Response interface{} `json:"response,omitempty"`
}
