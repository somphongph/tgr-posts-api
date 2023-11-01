package posts

import (
	"fmt"
	"net/http"
	"tgr-posts-api/internal/responses"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *Handler) AddPostHandler(c echo.Context) error {
	req := PostRequest{}

	// Binding
	if err := c.Bind(&req); err != nil {
		res := responses.ResponseError()
		return c.JSON(http.StatusBadRequest, res)
	}

	// Bind object
	post := &Post{
		Id:        primitive.NewObjectID(),
		Title:     req.Title,
		Caption:   req.Caption,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	fmt.Println(post)

	err := h.store.Add(post)
	if err != nil {
		res := responses.ResponseOperationFailed()
		return c.JSON(http.StatusInternalServerError, res)
	}

	res := PostResponse{}
	res.Id = post.Id.Hex()
	res.Title = post.Title
	res.Caption = post.Caption

	resp := responses.ResponseSuccess(res)

	return c.JSON(http.StatusOK, resp)
}
