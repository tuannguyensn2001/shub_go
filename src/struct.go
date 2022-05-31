package _struct

type Paginate struct {
	Page  int
	Limit int
}

type PaginateOutput struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
	Total int `json:"total"`
	Data  interface{}
}
