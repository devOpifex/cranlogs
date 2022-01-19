package data

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Daily struct {
	Downloads []DailyDownload `json:"downloads"`
	Start     string          `json:"start"`
	End       string          `json:"end"`
	Package   string          `json:"package"`
}

type DailyDownload struct {
	Day       string `json:"day"`
	Downloads int    `json:"downloads"`
}

func GetDaily(period, pkg string) (Daily, error) {
	// weirdly the API returns an array of length 1
	var daily []Daily

	path := URL + "downloads/daily/" + period + "/" + pkg

	resp, err := http.Get(path)

	if err != nil {
		return Daily{}, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return Daily{}, err
	}

	err = json.Unmarshal(body, &daily)

	if err != nil {
		return Daily{}, err
	}

	return daily[0], nil
}
