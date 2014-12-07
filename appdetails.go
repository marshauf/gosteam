package steam

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type AppDetail struct {
	Success bool     `json:"success"`
	Data    *AppData `json:"data"`
}

type AppData struct {
	Type         string `json:"type"`
	Name         string `json:"name"`
	SteamAppID   int    `json:"steam_appid"`
	RequiredAgee int    `json:"required_age"`
	IsFree       bool   `json:"is_free"`
	Dlc          []int  `json:"dlc"`

	DetailedDescription string `json:"detailed_description"`
	AboutTheGame        string `json:"about_the_game"`
	SupportedLanguages  string `json:"supported_languages"`
	Reviews             string `json:"reviews"`
	HeaderImage         string `json:"header_image"`
	Website             string `json:"website"`

	WindowsRequirements *Requirements `json:"pc_requirements"`
	MacRequirements     *Requirements `json:"mac_requirements"`
	Linux               *Requirements `json:"linux_requirements"` // TODO Tropico has no Linux support, field is an array instead of an object

	DrmNotice     string         `json:"drm_notice"`
	Developers    []string       `json:"developers"`
	Publishers    []string       `json:"publishers"`
	Demos         []App          `json:"demos"`
	PriceOverview *PriceOverview `json:"price_overview"`
	Packages      []string       `json:"packages"`
	PackageGroups []PackageGroup `json:"package_groups"`
}

type Requirements struct {
	Minimum     string `json:"minimum"`
	Recommended string `json:"recommended"`
	unmarshaled bool
}

func (r *Requirements) UnmarshalJSON(b []byte) error {
	if r.unmarshaled {
		return nil
	}
	r.unmarshaled = true
	if len(b) == 0 {
		return nil
	}
	if b[0] == '[' {
		return nil
	}
	return json.Unmarshal(b, r)
}

type PriceOverview struct {
	Currency        string `json:"currency"`
	Initial         int    `json:"initial"`
	Final           int    `json:"final"`
	DiscountPercent int    `json:"discount_percent"`
}

type PackageGroup struct {
	Name                    string           `json:"name"`
	Title                   string           `json:"title"`
	Description             string           `json:"description"`
	SelectionText           string           `json:"selection_text"`
	SaveText                string           `json:"save_text"`
	DisplayType             string           `json:"display_type"`
	IsRecurringSubscription string           `json:"is_recurring_subscription"`
	Subs                    []SubPackage     `json:"subs"`
	Platforms               map[string]bool  `json:"platforms"`
	Metacritic              *Metacritic      `json:"metacritic"`
	Categories              []Category       `json:"categories"`
	Genres                  []Genre          `json:"genres"`
	Screenshots             []Screenshot     `json:"screenshots"`
	Movies                  []Movie          `json:"movies"`
	Recommendations         *Recommendations `json:"recommendations"`
	Achievements            *Achievements    `json:"achievements"`
	ReleaseDate             *ReleaseDate     `json:"release_date"`
	SupportInfo             *SupportInfo     `json:"support_info"`
	Background              string           `json:"background"`
}

type SubPackage struct {
	PackageID                string `json:"packageid"`
	PercentSavingsText       string `json:"percent_savings_text"`
	PercentSavings           int    `json:"percent_savings"`
	OptionText               string `json:"option_text"`
	OptionDescription        string `json:"option_description"`
	CanGetFreeLicense        string `json:"can_get_free_license"`
	IsFreeLicense            bool   `json:"is_free_license"`
	PriceInCentsWithDiscount int    `json:"price_in_cents_with_discount"`
}

type Metacritic struct {
	Score int    `json:"score"`
	URL   string `json:"url"`
}

type Category struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}

type Genre struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}

type Screenshot struct {
	ID            int    `json:"id"`
	PathThumbnail string `json:"path_thumbnail"`
	PathFull      string `json:"path_full"`
}

type Movie struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Thumbnail string `json:"thumbnail"`
	Webm      *Webm  `json:"webm"`
	Highlight bool   `json:"highlight"`
}

type Webm struct {
	Hd480 string `json:"480"`
	Max   string `json:"max"`
}

type Recommendations struct {
	Total int `json:"total"`
}

type Achievements struct {
	Total       int           `json:"total"`
	Highlighted []Achievement `json:"highlighted"`
}

type Achievement struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type ReleaseDate struct {
	ComingSoon bool   `json:"coming_soon"`
	Date       string `json:"date"`
}

type SupportInfo struct {
	URL   string `json:"url"`
	Email string `json:"email"`
}

func GetAppDetails(appID int) (map[string]AppDetail, error) {
	//http://store.steampowered.com/api/appdetails/?appids=57690
	resp, err := http.Get("http://store.steampowered.com/api/appdetails/?appids=" + strconv.Itoa(appID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	appdetails := make(map[string]AppDetail)
	dec := json.NewDecoder(resp.Body)
	dec.Decode(&appdetails)
	return appdetails, err
}
