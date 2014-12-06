package steam

import (
	"encoding/json"
	"net/http"
)

type appListResponse struct {
	AppList *AppList `json:"applist"`
}

type AppList struct {
	Apps []App `json:"apps"`
}

type App struct {
	Id   int    `json:"appid"`
	Name string `json:"name"`
}

func GetAppList() (*AppList, error) {
	req, err := http.NewRequest("GET", "http://api.steampowered.com/ISteamApps/GetAppList/v0002/", nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	var appListResp appListResponse
	dec := json.NewDecoder(resp.Body)
	dec.Decode(&appListResp)
	return appListResp.AppList, nil
}
