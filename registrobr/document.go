package registrobr

// Document ...
type Document struct {
	EmpresaID    int
	CNPJ         string
	Domains      []string `json:"domains"`
	Telefone     string   `json:"telefone"`
	ContatoID    string   `json:"contato_id"`
	ContatoEmail string   `json:"contato_email"`
	ContatoNome  string   `json:"contato_nome"`
	Responsavel  string   `json:"responsavel"`
	Endereco     string   `json:"endereco"`
}
