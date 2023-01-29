package domain

const ERROR_DEFAULT = "failed to perform action"
const NO_PAYLOAD = "payload not found in headers"
const RESOURCE_NOT_FOUND = "resource not found"
const INTERNAL_ERROR = "internal server error"
const UNAUTHORIZED = "not allowed to perform this action"
const FORBIDDEN = "you are forbidden to perform this action"

type ErrorResponse struct {
	Message string `json:"message"`
}

func NewError(m string) *ErrorResponse {
	return &ErrorResponse{m}
}
