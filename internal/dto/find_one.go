package dto

type GetOne[T any] struct {
	Data T `json:"data"`
}

func NewGetOne[T any](data T) GetOne[T] {
	return GetOne[T]{
		Data: data,
	}
}
