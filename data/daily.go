package data

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Daily struct {
	Start     string
	End       string
	Package   string
	Downloads []DailyDownload
}

type DailyDownload struct {
	Day       string `json:"day"`
	Downloads int    `json:"downloads"`
}

func (api *API) Daily(period, pkg string) (Daily, error) {
	var data Daily

	path := api.URL + "downloads/daily/" + period + "/" + pkg

	resp, err := http.Get(path)

	if err != nil {
		return data, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return data, err
	}

	if err = json.Unmarshal(body, &data); err != nil {
		return data, err
	}

	return data, nil
}
