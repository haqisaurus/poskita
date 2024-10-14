package dto

type ResponseSuccess[T any] struct {
	ActivityRefCode string `json:"activityRefCode"`
	Code            string `json:"code"`
	Message         string `json:"message"`
	Data            T      `json:"data"`
}

type ResponseError struct {
	ActivityRefCode string `json:"activityRefCode"`
	Code            int    `json:"code"`
	Message         string `json:"message"`
	MessageError    string `json:"messageError"`
}

type InsertSuccess struct {
	ID uint64 `json:"id"`
}

type PaginationRs struct {
	Page    int         `json:"page"`
	PerPage int         `json:"perPage"`
	Total   int         `json:"totalItems"`
	Content interface{} `json:"content"`
}
