package utils

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Error  string      `json:"error"`
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
