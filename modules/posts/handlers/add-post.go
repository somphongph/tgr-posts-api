package handlers

import (
	"net/http"
	"tgr-posts-api/modules/posts/constants"
	"tgr-posts-api/modules/posts/entities"
	"tgr-posts-api/modules/shared/dto"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type addPostRequest struct {
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

type addPostResponse struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

func (h *Handler) AddPostHandler(c echo.Context) error {
	req := addPostRequest{}

	// Binding
	if err := c.Bind(&req); err != nil {
		res := dto.ResponseError()
		return c.JSON(http.StatusBadRequest, res)
	}

	// Bind object
	post := entities.Post{}
	post.Id = primitive.NewObjectID()
	post.Title = req.Title
	post.Detail = req.Detail
	post.ImageUrl = "abc"
	// post.PlaceTag = "sdfgsdfg"
	post.Status = constants.Active
	post.CreatedBy = "12345"
	post.CreatedOn = time.Time{}
	post.UpdatedBy = "12345"
	post.UpdatedOn = time.Time{}

	// Insert
	err := h.store.Add(&post)
	if err != nil {
		res := dto.ResponseOperationFailed()
		return c.JSON(http.StatusInternalServerError, res)
	}

	// Response
	res := addPostResponse{}
	res.Id = post.Id.Hex()
	res.Title = post.Title
	res.Detail = post.Detail

	resp := dto.ResponseSuccess(res)

	return c.JSON(http.StatusOK, resp)
}
