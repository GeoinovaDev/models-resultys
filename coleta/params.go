package coleta

// Params struct
type Params struct {
	IsLoadFromCache bool `json:"load_from_cache"`
}

// ToMap ...
func (p Params) ToMap() map[string]string {
	m := map[string]string{}

	if p.IsLoadFromCache {
		m["cache"] = "true"
	} else {
		m["cache"] = "false"
	}

	return m
}
