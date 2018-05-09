package empresa

import (
	"strings"

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

// GetNome retorna nome da empresa
// Return string
func GetNome(razaoSocial string, fantasia string) string {
	nome := fantasia
	if len(nome) == 0 {
		nome = razaoSocial
	}

	return ClearNome(nome)
}

// ClearNome ...
func ClearNome(nome string) string {
	nome = strings.Replace(nome, ".", "", -1)
	nome = strings.Replace(nome, "- ME", "", -1)
	nome = strings.Replace(nome, "LTDA", "", -1)
	nome = strings.Replace(nome, "EPP", "", -1)
	nome = strings.Replace(nome, "S/S", "", -1)
	nome = strings.Replace(nome, "S/C", "", -1)
	nome = strings.Replace(nome, "EIRELI", "", -1)
	nome = strings.Replace(nome, "-", "", -1)
	nome = strings.Replace(nome, `"`, "", -1)
	nome = strings.Replace(nome, `\`, "", -1)
	nome = strings.Replace(nome, `/`, "", -1)
	nome = strings.Replace(nome, `\\`, "", -1)
	nome = strings.Replace(nome, `(`, "", -1)
	nome = strings.Replace(nome, `)`, "", -1)

	nome = removeTipoEmpresa(nome, "ME")
	nome = removeTipoEmpresa(nome, "LTDA")
	nome = removeTipoEmpresa(nome, "EPP")

	return strings.Trim(nome, " ")
}

func removeTipoEmpresa(nome string, tipo string) string {
	nomes := strings.Split(nome, " ")
	if nomes[len(nomes)-1] == tipo {
		return strings.Join(nomes[:len(nomes)-1], " ")
	}

	return strings.Join(nomes, " ")
}
