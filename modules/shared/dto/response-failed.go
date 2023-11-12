package dto

import "tgr-posts-api/modules/shared/domains"

func ResponseError() (res domains.Response) {
	res = domains.Response{}
	res.Code = "cannot_bind_data"
	res.Message = "Cannot bind data"
	res.Data = nil

	return res
}

func ResponseDataNotFound() (res domains.Response) {
	res = domains.Response{}
	res.Code = "data_not_found"
	res.Message = "Data not found."
	res.Data = nil

	return res
}

func ResponseOperationFailed() (res domains.Response) {
	res = domains.Response{}
	res.Code = "operation_failed"
	res.Message = "The operation failed."
	res.Data = nil

	return res
}
