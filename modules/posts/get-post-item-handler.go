package posts

import (
	"net/http"
	"tgr-posts-api/modules/shared/dto"

	"github.com/labstack/echo/v4"
)

type postItemResponse struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

func (h *Handler) GetPostItemHandler(c echo.Context) error {
	id := c.Param("id")
	// post := Post{}

	// Get data
	post, err := h.store.GetById(id)
	if err != nil {
		res := dto.ResponseDataNotFound()
		return c.JSON(http.StatusNotFound, res)
	}

	// Save data to cache
	// data, _ := json.Marshal(book)
	// t.cache.SetShortCache(cacheKey, data)

	// Response
	res := postItemResponse{}
	res.Id = post.Id.Hex()
	res.Title = post.Title
	res.Detail = post.Detail

	resp := dto.ResponseSuccess(res)

	return c.JSON(http.StatusOK, resp)
}
