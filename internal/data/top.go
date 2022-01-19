package data

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

type TopDownload struct {
	Package   string `json:"package"`
	Downloads string `json:"downloads"`
}

type TopDownloads []TopDownload

type Top struct {
	Start     string        `json:"start"`
	End       string        `json:"end"`
	Downloads []TopDownload `json:"downloads"`
}

func GetTop(period string, count int) (Top, error) {
	var data Top

	acount := strconv.Itoa(count)
	path := URL + "top/" + period + "/" + acount

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
