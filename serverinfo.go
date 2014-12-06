package steam

import (
	"encoding/json"
	"net/http"
)

type ServerInfo struct {
	ServerTime       int    `json:"servertime"`
	ServerTimeString string `json:"servertimestring"`
}

func GetServerInfo() (*ServerInfo, error) {
	req, err := http.NewRequest("GET", "http://api.steampowered.com/ISteamWebAPIUtil/GetServerInfo/v0001/", nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	var info ServerInfo
	dec := json.NewDecoder(resp.Body)
	dec.Decode(&info)
	return &info, nil
}
