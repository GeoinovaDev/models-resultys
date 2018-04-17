package models

// Telefone dados
type Telefone struct {
	Numero   string `json:"numero"`
	Pais     string `json:"pais"`
	DDD      string `json:"ddd"`
	Fonte    string `json:"fonte"`
	CreateAt string `json:"create_at"`
}
