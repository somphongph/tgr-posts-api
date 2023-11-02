package posts

import (
	"net/http"
	"tgr-posts-api/internal/responses"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type postRequest struct {
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

type postResponse struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

func (h *Handler) AddPostHandler(c echo.Context) error {
	req := postRequest{}

	// Binding
	if err := c.Bind(&req); err != nil {
		res := responses.ResponseError()
		return c.JSON(http.StatusBadRequest, res)
	}

	// Bind object
	post := &Post{
		Id:        primitive.NewObjectID(),
		Title:     req.Title,
		Detail:    req.Detail,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Insert
	err := h.store.Add(post)
	if err != nil {
		res := responses.ResponseOperationFailed()
		return c.JSON(http.StatusInternalServerError, res)
	}

	// Response
	res := postResponse{}
	res.Id = post.Id.Hex()
	res.Title = post.Title
	res.Detail = post.Detail

	resp := responses.ResponseSuccess(res)

	return c.JSON(http.StatusOK, resp)
}
