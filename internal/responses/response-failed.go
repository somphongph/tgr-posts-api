package responses

import "tgr-posts-api/internal/models"

func ResponseError() (res models.Response) {
	res = models.Response{}
	res.Code = "cannot_bind_data"
	res.Message = "Cannot bind data"
	res.Data = nil

	return res
}

func ResponseDataNotFound() (res models.Response) {
	res = models.Response{}
	res.Code = "data_not_found"
	res.Message = "Data not found."
	res.Data = nil

	return res
}
