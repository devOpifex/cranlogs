package data

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Trend struct {
	Package  string `json:"package"`
	Increase string `json:"increase"`
}

type Trending []Trend

func (api *API) Trending() (Trending, error) {
	var trending Trending

	path := api.URL + "trending/"

	resp, err := http.Get(path)

	if err != nil {
		return trending, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return trending, err
	}

	json.Unmarshal(body, &trending)

	return trending, nil
}
