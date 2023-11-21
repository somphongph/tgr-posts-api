package handlers

import (
	"net/http"
	"strconv"
	"tgr-posts-api/modules/shared/dto"
	"tgr-posts-api/modules/shared/models"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

type getPostListResponse struct {
	Id       string `json:"id"`
	Title    string `json:"title"`
	Detail   string `json:"detail"`
	ImageUrl string `json:"imageUrl"`
	PlaceTag string `json:"placeTag"`
}

func (h *Handler) GetListPostHandler(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		page = 1
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		limit = 20
	}

	filter := bson.M{}
	sort := bson.D{{Key: "createdOn", Value: -1}}

	// Get data
	posts, err := h.store.Fetch(filter, sort, page, limit)
	if err != nil {
		res := dto.ResponseOperationFailed()
		return c.JSON(http.StatusNotFound, res)
	}

	// Count
	count, err := h.store.Count(filter)
	if err != nil {
		res := dto.ResponseOperationFailed()
		return c.JSON(http.StatusNotFound, res)
	}

	// Response
	l := getPostListResponse{}
	res := []getPostListResponse{}
	for _, v := range posts {
		l.Id = v.Id.Hex()
		l.Title = v.Title
		l.Detail = v.Detail
		l.ImageUrl = v.ImageUrl
		l.PlaceTag = v.PlaceTag

		res = append(res, l)
	}

	p := models.Paging{}
	p.Page = page
	p.Limit = limit
	p.Total = count

	resp := dto.ResponsePagingSuccess(res, p)

	return c.JSON(http.StatusOK, resp)
}
