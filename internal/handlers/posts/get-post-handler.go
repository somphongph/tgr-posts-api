package posts

import (
	"net/http"
	"tgr-posts-api/internal/responses"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetPostHandler(c echo.Context) error {
	p := Post{}

	// Request binding
	if err := c.Bind(&p); err != nil {
		res := responses.ResponseError()
		return c.JSON(http.StatusBadRequest, res)
	}

	// Pass
	res := responses.ResponseSuccess(nil)

	return c.JSON(http.StatusOK, res)
}
