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

func ResponseAddressMismatch() (res models.Response) {
	res = models.Response{}
	res.Code = "address_mismatch"
	res.Message = "Address mismatch."
	res.Data = nil

	return res
}

func ResponseTitleNameMismatch() (res models.Response) {
	res = models.Response{}
	res.Code = "title_name_mismatch"
	res.Message = "Title name mismatch."
	res.Data = nil

	return res
}

func ResponseCarBrandInfoMismatch() (res models.Response) {
	res = models.Response{}
	res.Code = "car_brand_info_mismatch"
	res.Message = "CarBrandInfo mismatch."
	res.Data = nil

	return res
}

func ResponseColorMismatch() (res models.Response) {
	res = models.Response{}
	res.Code = "color_mismatch"
	res.Message = "Color mismatch."
	res.Data = nil

	return res
}

func ResponseOccupationMismatch() (res models.Response) {
	res = models.Response{}
	res.Code = "occupation_mismatch"
	res.Message = "Occupation mismatch."
	res.Data = nil

	return res
}

func ResponseAmloMismatch() (res models.Response) {
	res = models.Response{}
	res.Code = "amlo_mismatch"
	res.Message = "AMLO mismatch."
	res.Data = nil

	return res
}
