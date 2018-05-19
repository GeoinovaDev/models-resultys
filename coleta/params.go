package coleta

// Params struct
type Params struct {
	IsLoadFromCache bool `json:"load_from_cache"`
}

// ToMap ...
func (p Params) ToMap() map[string]string {
	m := map[string]string{}

	if p.IsLoadFromCache {
		m["is_load_cache"] = "true"
	} else {
		m["is_load_cache"] = "false"
	}

	return m
}
