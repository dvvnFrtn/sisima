package dto

type Meta struct {
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	TotalItems int64 `json:"total_items"`
	TotalPages int   `json:"total_pages"`
}

type Pagination struct {
	Data []interface{} `json:"data"`
	Meta Meta          `json:"meta"`
}

func NewPagination(data []interface{}, page, limit int, totalItems int64) Pagination {
	return Pagination{
		Data: data,
		Meta: Meta{
			Page:       page,
			Limit:      limit,
			TotalItems: totalItems,
			TotalPages: int((totalItems + int64(limit) - 1) / int64(limit)),
		},
	}
}
