package handlers

import (
	"net/http"
	"tgr-posts-api/modules/shared/dto"
	"tgr-posts-api/modules/shared/models"

	"github.com/labstack/echo/v4"
)

type getListItemResponse struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

func (h *Handler) GetListPostHandler(c echo.Context) error {

	// Get data
	p, err := h.store.Fetch(1, 1)
	if err != nil {
		res := dto.ResponseDataNotFound()
		return c.JSON(http.StatusNotFound, res)
	}

	// fmt.Printf("%v", p)
	// Save data to cache
	// data, _ := json.Marshal(book)
	// t.cache.SetShortCache(cacheKey, data)

	// Response
	l := getListItemResponse{}
	res := []getListItemResponse{}
	for _, v := range p {
		l.Id = v.Id.Hex()
		l.Title = v.Title
		l.Detail = v.Detail

		res = append(res, l)
	}

	page := models.Paging{}
	page.Page = 1
	page.Limit = 20
	page.Total = 200

	resp := dto.ResponsePagingSuccess(res, page)

	return c.JSON(http.StatusOK, resp)
}
