package handlers

import (
	"net/http"
	"tgr-posts-api/modules/posts/entities"
	"tgr-posts-api/modules/shared/dto"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type addPostRequest struct {
	Title    string `json:"title"`
	Detail   string `json:"detail"`
	ImageUrl string `json:"imageUrl"`
	PlaceTag string `json:"placeTag"`
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
	p := entities.Post{}
	p.Id = primitive.NewObjectID()
	p.Title = req.Title
	p.Detail = req.Detail
	p.ImageUrl = req.ImageUrl
	p.PlaceTag = req.PlaceTag

	// Insert
	if err := h.store.Add(&p); err != nil {
		res := dto.ResponseOperationFailed()
		return c.JSON(http.StatusInternalServerError, res)
	}

	// Response
	res := addPostResponse{}
	res.Id = p.Id.Hex()
	res.Title = p.Title
	res.Detail = p.Detail

	resp := dto.ResponseSuccess(res)

	return c.JSON(http.StatusOK, resp)
}
