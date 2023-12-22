package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"tgr-posts-api/configs"
	"tgr-posts-api/modules/shared/models"
)

type accountApi struct {
	url string
}

type response struct {
	Code    string                `json:"code"`
	Message string                `json:"message"`
	Data    models.AccountProfile `json:"data"`
}

func InitAccountApi(c *configs.Tgr) *accountApi {
	return &accountApi{
		url: c.AuthApi,
	}
}

func (a *accountApi) GetAccount(acc string) (models.AccountProfile, error) {
	resp := response{}
	url := fmt.Sprintf("%v/v1/users/%v/profile", a.url, acc)
	res, err := http.Get(url)
	if err != nil {
		return resp.Data, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return resp.Data, err
	}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return resp.Data, err
	}

	return resp.Data, err
}
