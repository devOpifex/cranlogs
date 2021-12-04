package data

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Total struct {
	Start     string `json:"start"`
	End       string `json:"end"`
	Downloads int    `json:"downloads"`
	Package   string `json:"package"`
}

func (api *API) Total(period, pkg string) (Total, error) {
	// weirdly the API returns an array of length 1
	var total []Total

	path := api.URL + "downloads/total/" + period + "/" + pkg

	resp, err := http.Get(path)

	if err != nil {
		return total[0], err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return total[0], err
	}

	json.Unmarshal(body, &total)

	return total[0], nil
}
