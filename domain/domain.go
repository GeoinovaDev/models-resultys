package domain

import (
	"net/url"
	"strings"
)

// Domain struct
type Domain struct {
	URL     string `json:"url" bson:"url"`
	Ranking int    `json:"ranking" bson:"ranking"`
	Repeat  int    `json:"repeat" bson:"repeat"`
	Status  string `json:"status" bson:"status"`
}

// Raw ...
func (d Domain) Raw() string {
	u := d.URL

	if len(u) == 0 {
		return ""
	}

	if strings.Index(u, "http") == -1 {
		u = "http://" + u
	}

	r, err := url.Parse(u)
	if err != nil {
		return ""
	}

	h := r.Host
	h = strings.Replace(h, "www.", "", -1)
	h = strings.Replace(h, "www1.", "", -1)
	h = strings.Replace(h, "www3.", "", -1)
	h = strings.Replace(h, "www3.", "", -1)

	return h
}

// GetDomain ...
func (d *Domain) GetDomain() string {
	return d.Raw()
}

// ExtractDomain ...
func ExtractDomain(url string) string {
	return Domain{URL: url}.Raw()
}
