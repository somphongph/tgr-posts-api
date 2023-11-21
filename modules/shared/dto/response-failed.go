package dto

import "tgr-posts-api/modules/shared/models"

func CannotBindData() (res models.Response) {
	res = models.Response{}
	res.Code = "cannot_bind_data"
	res.Message = "Cannot bind data"
	res.Data = nil

	return res
}

func DataNotFound() (res models.Response) {
	res = models.Response{}
	res.Code = "data_not_found"
	res.Message = "Data not found."
	res.Data = nil

	return res
}

func OperationFailed() (res models.Response) {
	res = models.Response{}
	res.Code = "operation_failed"
	res.Message = "The operation failed."
	res.Data = nil

	return res
}
