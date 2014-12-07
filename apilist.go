package steam

import (
	"encoding/json"
	"net/http"
)

type supportedAPIListResponse struct {
	APIList *APIList `json:"apilist"`
}

type APIList struct {
	Interfaces []Interface `json:"interfaces"`
}

type Interface struct {
	Name    string   `json:"name"`
	Methods []Method `json:"methods"`
}

type Method struct {
	Name       string      `json:"name"`
	Version    int         `json:"version"`
	HTTPMethod string      `json:"httpmethod"`
	Parameters []Parameter `json:"parameters"`
}

type Parameter struct {
	Name        string `json:"parameter"`
	Type        string `json:"type"`
	Optional    bool   `json:"optional"`
	Description string `json:"description"`
}

func GetSupportedAPIList(accessKey string) (*APIList, error) {
	url := "http://api.steampowered.com/ISteamWebAPIUtil/GetSupportedAPIList/v0001/"
	if len(accessKey) > 0 {
		url += "?key=" + accessKey
	}
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	var apiListResp supportedAPIListResponse
	dec := json.NewDecoder(resp.Body)
	dec.Decode(&apiListResp)
	return apiListResp.APIList, nil
}
