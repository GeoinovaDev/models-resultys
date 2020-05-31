package gmaps

// Comment ...
type Comment struct {
	Author string
	Text   string
}

// Company ...
type Company struct {
	Name    string   `json:"name"`
	LatLng  []string `json:"latlng"`
	Address string   `json:"address"`
	Phones  []string `json:"phone"`
	Sites   []string `json:"site"`
	Status  string   `json:"status"`
	Nota    float32  `json:"nota"`
}
