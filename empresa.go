package models

// Empresa dados
type Empresa struct {
	CNPJ        string `json:"cnpj"`
	RazaoSocial string `json:"razao_social"`
	Fantasia    string `json:"fantasia"`
	Municipio   string `json:"municipio"`
	UF          string `json:"uf"`
}
