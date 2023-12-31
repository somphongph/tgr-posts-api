package responses

import "tgr-posts-api/internal/models"

func ResponseSuccess(i interface{}) (res models.Response) {
	res = models.Response{}
	res.Code = "success"
	res.Message = "Success"
	res.Data = i

	return res
}

func ResponseItemSuccess(i interface{}, c bool) (res models.ResponseItem) {
	res = models.ResponseItem{}
	res.Code = "success"
	res.Message = "Success"
	res.IsCached = c
	res.Data = i

	return res
}

func ResponsePagingSuccess(i interface{}, p models.Paging) (res models.ResponsePaging) {
	res = models.ResponsePaging{}
	res.Code = "success"
	res.Message = "Success"
	res.Data = i
	res.Page = p.Page
	res.Limit = p.Limit
	res.Total = p.Total

	return res
}
