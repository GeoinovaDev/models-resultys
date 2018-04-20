package token

import "strings"

// Token dados
type Token struct {
	CNPJ     string
	Nome     string
	Fantasia string
	Cidade   string
	Estado   string
	Webhook  string
}

// GetNome retorna nome da empresa
// Return string
func (e *Token) GetNome() string {
	nome := e.Fantasia
	if len(nome) == 0 {
		nome = e.Nome
	}

	return clearNome(nome)
}

func clearNome(nome string) string {
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
