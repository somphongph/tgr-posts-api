package posts

import (
	"net/http"
	"tgr-posts-api/internal/responses"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetPostHandler(c echo.Context) error {
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

	res := responses.ResponseSuccess(post)

	return c.JSON(http.StatusOK, res)
}
