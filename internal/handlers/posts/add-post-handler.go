package posts

import (
	"net/http"
	"tgr-posts-api/internal/responses"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *Handler) AddPostHandler(c echo.Context) error {
	p := Post{}

	// Binding
	if err := c.Bind(&p); err != nil {
		res := responses.ResponseError()
		return c.JSON(http.StatusBadRequest, res)
	}

	// Bind object
	post := &Post{
		Id:        primitive.NewObjectID(),
		Title:     "Title",
		Caption:   "Caption",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := h.store.Add(post)
	if err != nil {
		// res := responses.ResponseError()
		// c.JSON(http.StatusInternalServerError, common.ResponseFailed())

		return err
	}

	// return c.JSON(http.StatusOK, common.Response(book))

	res := responses.ResponseSuccess(nil)

	return c.JSON(http.StatusOK, res)
}
