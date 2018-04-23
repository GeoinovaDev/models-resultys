package empresa

import (
	"git.resultys.com.br/motor/models/email"
	"git.resultys.com.br/motor/models/socio"
	"git.resultys.com.br/motor/models/telefone"
)

// Empresa dados
type Empresa struct {
	CNPJ                 string              `json:"cnpj"`
	RazaoSocial          string              `json:"razao_social"`
	Fantasia             string              `json:"fantasia"`
	Abertura             string              `json:"abertura"`
	NaturezaJuridica     string              `json:"natureza_juridica"`
	Logradouro           string              `json:"logradouro"`
	Numero               string              `json:"numero"`
	Complemento          string              `json:"complemento"`
	CEP                  string              `json:"cep"`
	Bairro               string              `json:"bairro"`
	Municipio            string              `json:"municipio"`
	UF                   string              `json:"uf"`
	Emails               []email.Email       `json:"emails"`
	Telefones            []telefone.Telefone `json:"telefones"`
	EFR                  string              `json:"efr"`
	Situacao             string              `json:"situacao"`
	DataSituacao         string              `json:"data_situacao"`
	MotivoSituacao       string              `json:"motivo_situacao"`
	SituacaoEspecial     string              `json:"situacao_especial"`
	DataSituacaoEspecial string              `json:"data_situacao_especial"`
	CapitalSocial        string              `json:capital_social""`

	AtividadesPrincipais  []Atividade   `json: "atividades_principais"`
	AtividadesSecundarias []Atividade   `json: "atividades_secundarias"`
	Socios                []socio.Socio `json: "socios"`
}
