package posts

import (
	"net/http"
	"tgr-posts-api/internal/responses"

	"github.com/labstack/echo/v4"
)

type PostItemResponse struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Caption string `json:"caption"`
}

func (h *Handler) GetPostItemHandler(c echo.Context) error {
	id := c.Param("id")
	// post := Post{}

	// Get data
	post, err := h.store.GetById(id)
	if err != nil {
		res := responses.ResponseDataNotFound()
		return c.JSON(http.StatusNotFound, res)
	}

	// Save data to cache
	// data, _ := json.Marshal(book)
	// t.cache.SetShortCache(cacheKey, data)

	// Response
	res := PostItemResponse{}
	res.Id = post.Id.Hex()
	res.Title = post.Title
	res.Caption = post.Caption

	resp := responses.ResponseSuccess(res)

	return c.JSON(http.StatusOK, resp)
}
