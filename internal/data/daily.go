package data

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
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

func GetDaily(period Period, pkgs []string) ([]Daily, error) {
	// weirdly the API returns an array of length 1
	var daily []Daily
	path := URL + "downloads/daily/" + string(period) + "/" + strings.Join(pkgs, ",")
	resp, err := http.Get(path)

	if err != nil {
		return daily, err
	}
	// This is currently kind of meaningless until
	// https://github.com/r-hub/cranlogs.app/issues/41 is resolved but at least
	// its there
	if resp.StatusCode != 200 {
		return daily, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return daily, err
	}

	err = json.Unmarshal(body, &daily)
	// if this errors given the current 200 status code returning an error response
	// such as { "error": "Invalid query",   "info": "https://github.com/metacran/cranlogs.app" }
	// its still not the best error code
	if err != nil {
		return daily, err
	}

	return daily, nil
}
