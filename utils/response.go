package utils

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data,omitempty"`
	Error  string      `json:"error,omitempty"`
}

/*
A function to create a new response.
*/
func NewResponse(status int, data interface{}, error string) Response {
	return Response{
		Status: status,
		Data:   data,
		Error:  error,
	}
}
