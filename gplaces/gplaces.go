package gplaces

import (
	"strings"
)

// Result struct
type Result struct {
	Results []Company `json:"results"`
	Result  Company   `json:"result"`
	Status  string    `json:"status"`
}

// Company struct
type Company struct {
	ID        string   `json:"id"`
	URL       string   `json:"url"`
	Website   string   `json:"website"`
	Name      string   `json:"name"`
	Scope     string   `json:"scope"`
	PlaceID   string   `json:"place_id"`
	Rating    float64  `json:"rating"`
	Reference string   `json:"reference"`
	Geometry  Geometry `json:"geometry"`
	Icon      string   `json:"icon"`
	Types     []string `json:"types"`
	OpenHours OpenHour `json:"opening_hours"`

	Address          string             `json:"adr_address"`
	AddressFormatted string             `json:"formatted_address"`
	AddressComponent []AddressComponent `json:"address_components"`

	PhoneNumber       string    `json:"formatted_phone_number"`
	GlobalPhoneNumber string    `json:"international_phone_number"`
	Comments          []Comment `json:"reviews"`
}

// AddressComponent struct
type AddressComponent struct {
	LongName  string   `json:"long_name"`
	ShortName string   `json:"short_name"`
	Types     []string `json:"types"`
}

// ViewPort struct
type ViewPort struct {
	Northeast Location `json:"northeast"`
	Southwest Location `json:"southwest"`
}

// OpenHour struct
type OpenHour struct {
	IsOpenNow bool                   `json:"open_now"`
	Weekday   []string               `json:"weekday_text"`
	Periods   map[string]interface{} `json:"periods"`
}

// Location strutc
type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

// Geometry struct
type Geometry struct {
	Location Location `json:"location"`
	ViewPort ViewPort `json:"viewport"`
}

// Comment struct
type Comment struct {
	Name     string  `json:"author_name"`
	URL      string  `json:"author_url"`
	Language string  `json:"language"`
	PhotoURL string  `json:"profile_photo_url"`
	Rating   float64 `json:"rating"`
	Text     string  `json:"text"`
	Time     int     `json:"time"`
}

// GetPhone ...
func (r *Company) GetPhone() string {
	p := strings.Replace(r.GlobalPhoneNumber, "-", "", -1)
	p = strings.Replace(r.GlobalPhoneNumber, "(", "", -1)
	p = strings.Replace(r.GlobalPhoneNumber, ")", "", -1)
	p = strings.Replace(r.GlobalPhoneNumber, " ", "", -1)
	p = strings.Replace(r.GlobalPhoneNumber, "+", "", -1)

	return p
}
