package dtowrapper

type normalWrapperResponse[T any] struct {
	Data T `json:"data"`
}

func NewNormalWrapperResponse[T any](data T) normalWrapperResponse[T] {
	return normalWrapperResponse[T]{
		Data: data,
	}
}

// Pagination Response Wrapper
type meta struct {
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	TotalItems int64 `json:"total_items"`
	TotalPages int   `json:"total_pages"`
}

type paginationWrapperResponse struct {
	Data []interface{} `json:"data"`
	Meta meta          `json:"meta"`
}

func NewPaginationWrapperResponse(data []interface{}, page, limit int, totalItems int64) paginationWrapperResponse {
	return paginationWrapperResponse{
		Data: data,
		Meta: meta{
			Page:       page,
			Limit:      limit,
			TotalItems: totalItems,
			TotalPages: int((totalItems + int64(limit) - 1) / int64(limit)),
		},
	}
}
