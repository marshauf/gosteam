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
	resp, err := http.Get("http://api.steampowered.com/ISteamWebAPIUtil/GetServerInfo/v0001/")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var info ServerInfo
	dec := json.NewDecoder(resp.Body)
	dec.Decode(&info)
	return &info, nil
}
