package models

type Paging struct {
	Page  int   `json:"page"`
	Limit int   `json:"limit"`
	Total int64 `json:"total"`
}

type Response struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseItem struct {
	IsCached bool `json:"isCached"`
	Response
}

type ResponseList struct {
	Response
}

type ResponsePaging struct {
	Response
	Paging
}
