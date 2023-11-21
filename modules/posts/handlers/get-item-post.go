package handlers

import (
	"net/http"
	"tgr-posts-api/modules/shared/dto"

	"github.com/labstack/echo/v4"
)

type getPostItemResponse struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

func (h *Handler) GetItemPostHandler(c echo.Context) error {
	id := c.Param("id")

	// Get data
	p, err := h.store.GetById(id)
	if err != nil {
		res := dto.DataNotFound()
		return c.JSON(http.StatusNotFound, res)
	}

	// Save data to cache
	// data, _ := json.Marshal(book)
	// t.cache.SetShortCache(cacheKey, data)

	// Response
	res := getPostItemResponse{}
	res.Id = p.Id.Hex()
	res.Title = p.Title
	res.Detail = p.Detail

	resp := dto.Success(res)

	return c.JSON(http.StatusOK, resp)
}
