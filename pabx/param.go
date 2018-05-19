package pabx

// Params struct
type Params struct {
	IsConfirmaTelefone bool `json:"is_confirma_telefone"`
}

// ToMap ...
func (p Params) ToMap() map[string]string {
	m := map[string]string{}

	if p.IsConfirmaTelefone {
		m["is_confirma_telefone"] = "true"
	} else {
		m["is_confirma_telefone"] = "false"
	}

	return m
}
