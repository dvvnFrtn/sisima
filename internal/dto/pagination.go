package dto

type Meta struct {
	Page int `json:"page"`
}

type Pagination struct {
	Data []interface{} `json:"data"`
	Meta Meta          `json:"meta"`
}
