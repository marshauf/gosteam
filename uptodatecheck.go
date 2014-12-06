package steam

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type upToDateCheckResponse struct {
	Response *UpToDateCheckResponse `json:"response"`
}

type UpToDateCheckResponse struct {
	Success           bool   `json:"success"`
	UpToDate          bool   `json:"up_to_date"`
	VersionIsListable bool   `json:"version_is_listable"`
	Message           string `json:"message"`
}

func UpToDateCheck(appid, version int) (*UpToDateCheckResponse, error) {
	req, err := http.NewRequest("GET", "http://api.steampowered.com/ISteamApps/UpToDateCheck/v0001/", nil)
	if err != nil {
		return nil, err
	}
	req.Form = map[string][]string{
		"appid":   []string{strconv.Itoa(appid)},
		"version": []string{strconv.Itoa(version)},
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	var utdcResp upToDateCheckResponse
	dec := json.NewDecoder(resp.Body)
	dec.Decode(&utdcResp)
	return utdcResp.Response, nil
}
