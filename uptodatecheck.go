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
	RequiredVersion   int    `json:"required_version"`
	VersionIsListable bool   `json:"version_is_listable"`
	Message           string `json:"message"`
}

func UpToDateCheck(appid, version int) (*UpToDateCheckResponse, error) {
	// http://api.steampowered.com/ISteamApps/UpToDateCheck/v0001/?appid=570&version=36
	url := "http://api.steampowered.com/ISteamApps/UpToDateCheck/v0001/?appid=" + strconv.Itoa(appid) + "&version=" + strconv.Itoa(version)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var utdcResp upToDateCheckResponse
	dec := json.NewDecoder(resp.Body)
	dec.Decode(&utdcResp)
	return utdcResp.Response, nil
}
