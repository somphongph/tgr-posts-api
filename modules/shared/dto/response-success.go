package dto

import "tgr-posts-api/modules/shared/domains"

func ResponseSuccess(i interface{}) (res domains.Response) {
	res = domains.Response{}
	res.Code = "success"
	res.Message = "Success"
	res.Data = i

	return res
}

func ResponseItemSuccess(i interface{}, c bool) (res domains.ResponseItem) {
	res = domains.ResponseItem{}
	res.Code = "success"
	res.Message = "Success"
	res.IsCached = c
	res.Data = i

	return res
}

func ResponsePagingSuccess(i interface{}, p domains.Paging) (res domains.ResponsePaging) {
	res = domains.ResponsePaging{}
	res.Code = "success"
	res.Message = "Success"
	res.Data = i
	res.Page = p.Page
	res.Limit = p.Limit
	res.Total = p.Total

	return res
}
