package repositories

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetAccount(acc string) (interface{}, error) {
	url := fmt.Sprintf("%v/v1/users/%v", "https://sit-auth.tripgator.com/auth-api", acc)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var data any
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return data, err
}
