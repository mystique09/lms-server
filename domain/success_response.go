package domain

const OK_DEFAULT = "query success"
const OK_ONE = "successfully fetched one data"
const OK_ALL = "successfully fetched all data"
const OK_CREATE = "successfully created"
const OK_UPDATE = "successfully updated"
const OK_DELETE = "successfully deleted"

type SuccessResponse[T any] struct {
	Message string `json:"message"`
	Data    T      `json:"data"`
}

func OkResponse[T any](m string, t T) *SuccessResponse[T] {
	return &SuccessResponse[T]{
		Message: m,
		Data:    t,
	}
}
