package posts

import (
	"net/http"
	"tgr-posts-api/internal/responses"

	"github.com/labstack/echo/v4"
)

func (h *Handler) AddPostHandler(c echo.Context) error {
	p := Post{}

	// Binding
	if err := c.Bind(&p); err != nil {
		res := responses.ResponseError()
		return c.JSON(http.StatusBadRequest, res)
	}

	res := responses.ResponseSuccess(nil)

	return c.JSON(http.StatusOK, res)
}
