package repositories

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"tgr-posts-api/configs"
)

type accountApi struct {
	url string
}

type response struct {
	Code    string         `json:"code"`
	Message string         `json:"message"`
	Data    accountProfile `json:"data"`
}

type accountProfile struct {
	Id          string `json:"id"`
	Username    string `json:"username"`
	DisplayName string `json:"displayName"`
	PhotoUrl    string `json:"photoUrl"`
}

func InitAccountApi(c *configs.Tgr) *accountApi {
	return &accountApi{
		url: c.AuthApi,
	}
}

func (a *accountApi) GetAccount(acc string) (accountProfile, error) {
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
