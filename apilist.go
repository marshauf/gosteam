package steam

import (
	"net/http"
	"encoding/json"
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
	HttpMethod string      `json:"httpmethod"`
	Parameters []Parameter `json:"parameters"`
}

type Parameter struct {
	Name        string `json:"parameter"`
	Type        string `json:"type"`
	Optional    bool   `json:"optional"`
	Description string `json:"description"`
}

func GetSupportedAPIList(accessKey string) (*APIList, error) {
	req, err := http.NewRequest("GET", "http://api.steampowered.com/ISteamWebAPIUtil/GetSupportedAPIList/v0001/", nil)
	if err != nil {
		return nil, err
	}
	if len(accessKey) > 0 {
		req.Form["key"] = []string{accessKey}
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	var apiListResp supportedAPIListResponse
	dec := json.NewDecoder(resp.Body)
	dec.Decode(&apiListResp)
	return apiListResp.APIList, nil
}
