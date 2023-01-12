package domain

type SuccessResponse[T any] struct {
	Message string `json:"message"`
	Data    T      `json:"data"`
}
