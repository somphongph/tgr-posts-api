package handlers

import (
	"net/http"
	"tgr-posts-api/modules/posts/constants"
	"tgr-posts-api/modules/posts/domains"
	"tgr-posts-api/modules/shared/dto"
	"tgr-posts-api/modules/shared/models"

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
		res := dto.ResponseError()
		return c.JSON(http.StatusBadRequest, res)
	}

	// Bind object
	post := domains.Post{
		Id:       primitive.NewObjectID(),
		Title:    req.Title,
		Detail:   req.Detail,
		ImageUrl: "abc",
		Entity: models.Entity{
			Status:    constants.Active,
			CreatedBy: "12345",
			CreatedOn: time.Time{},
			UpdatedBy: "12345",
			UpdatedOn: time.Time{},
		},
	}

	// Insert
	err := h.store.Add(&post)
	if err != nil {
		res := dto.ResponseOperationFailed()
		return c.JSON(http.StatusInternalServerError, res)
	}

	// Response
	res := postResponse{}
	res.Id = post.Id.Hex()
	res.Title = post.Title
	res.Detail = post.Detail

	resp := dto.ResponseSuccess(res)

	return c.JSON(http.StatusOK, resp)
}
