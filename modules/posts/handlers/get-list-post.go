package handlers

import (
	"net/http"
	"strconv"
	"tgr-posts-api/modules/shared/dto"
	"tgr-posts-api/modules/shared/models"
	"tgr-posts-api/modules/shared/repositories/services"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

type getPostListResponse struct {
	Id       string                `json:"id"`
	Title    string                `json:"title"`
	Detail   string                `json:"detail"`
	ImageUrl string                `json:"imageUrl"`
	PlaceTag string                `json:"placeTag"`
	Account  models.AccountProfile `json:"account"`
}

func (h *handler) GetListPostHandler(c echo.Context) error {
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
		res := dto.OperationFailed()
		return c.JSON(http.StatusInternalServerError, res)
	}
	account := services.InitAccountApi(&h.cfg.Tgr)

	// Response
	l := getPostListResponse{}
	res := []getPostListResponse{}
	for _, v := range posts {
		l.Id = v.Id.Hex()
		l.Title = v.Title
		l.Detail = v.Detail
		l.ImageUrl = v.ImageUrl
		l.PlaceTag = v.PlaceTag

		// Get Account
		acc, _ := account.GetAccount(v.CreatedBy)
		l.Account = acc

		res = append(res, l)
	}

	resp := dto.ListSuccess(res)

	return c.JSON(http.StatusOK, resp)
}
