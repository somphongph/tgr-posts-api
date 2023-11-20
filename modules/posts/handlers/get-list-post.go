package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"tgr-posts-api/modules/shared/dto"
	"tgr-posts-api/modules/shared/models"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

type getListItemResponse struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

func (h *Handler) GetListPostHandler(c echo.Context) error {
	filter := bson.M{}
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	fmt.Println(page)

	// Get data
	posts, err := h.store.Fetch(filter, page, limit)
	if err != nil {
		res := dto.ResponseDataNotFound()
		return c.JSON(http.StatusNotFound, res)
	}

	// fmt.Printf("%v", p)
	// Save data to cache
	// data, _ := json.Marshal(book)
	// t.cache.SetShortCache(cacheKey, data)

	// Response
	l := getListItemResponse{}
	res := []getListItemResponse{}
	for _, v := range posts {
		l.Id = v.Id.Hex()
		l.Title = v.Title
		l.Detail = v.Detail

		res = append(res, l)
	}

	p := models.Paging{}
	p.Page = 1
	p.Limit = 20
	p.Total = 200

	resp := dto.ResponsePagingSuccess(res, p)

	return c.JSON(http.StatusOK, resp)
}
