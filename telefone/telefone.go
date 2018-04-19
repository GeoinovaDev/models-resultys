package telefone

import (
	"strings"

	"git.resultys.com.br/prospecta/models/telefone/tipo"
)

// Telefone dados
type Telefone struct {
	Numero   string `json:"numero"`
	Pais     string `json:"pais"`
	DDD      string `json:"ddd"`
	CreateAt string `json:"create_at"`
	Fonte    int    `json:"fonte"`
	Tipo     int    `json:"tipo"`
}

// New Cria telefone a partir do numero
// Return Telefone
func New(numero string) Telefone {
	telefone := Telefone{}

	numero = clear(numero)

	// TELEFONE LOCAL
	if strings.Index(numero, "4003") == 0 {
		telefone.Tipo = tipo.LOCAL
		telefone.Pais = ""
		telefone.DDD = ""
		telefone.Numero = numero

		return telefone
	}

	// TELEFONE LOCAL
	if strings.Index(numero, "4004") == 0 {
		telefone.Tipo = tipo.LOCAL
		telefone.Pais = ""
		telefone.DDD = ""
		telefone.Numero = numero

		return telefone
	}

	// TELEFONE GRATUITO
	if strings.Index(numero, "0800") == 0 {
		telefone.Tipo = tipo.GRATUITO
		telefone.Pais = ""
		telefone.DDD = ""
		telefone.Numero = numero

		return telefone
	}

	// CELULAR COM DDI
	if len(numero) == 13 {
		telefone.Tipo = tipo.CELULAR
		telefone.Pais = numero[:2]
		telefone.DDD = numero[2:4]
		telefone.Numero = numero[4:]

		return telefone
	}

	// CELULAR COM DDD
	if len(numero) == 11 {
		telefone.Tipo = tipo.CELULAR
		telefone.Pais = ""
		telefone.DDD = numero[:2]
		telefone.Numero = numero[2:]

		return telefone
	}

	// FIXO COM DDI
	if len(numero) == 12 {
		telefone.Tipo = tipo.FIXO
		telefone.Pais = numero[:2]
		telefone.DDD = numero[2:4]
		telefone.Numero = numero[4:]

		return telefone
	}

	// FIXO COM DDD
	if len(numero) == 10 {
		telefone.Tipo = tipo.FIXO
		telefone.Pais = ""
		telefone.DDD = numero[:2]
		telefone.Numero = numero[2:]

		return telefone
	}

	return telefone
}

func clear(numero string) string {
	numero = strings.Replace(numero, "-", "", -1)
	numero = strings.Replace(numero, "/", "", -1)
	numero = strings.Replace(numero, ")", "", -1)
	numero = strings.Replace(numero, "(", "", -1)

	return strings.Trim(numero, " ")
}

// EXEMPLOS DE NUMERO DE TELEFONES VALIDOS
// CELULAR: 5562982334440 - 62982334440
// FIXO: 556232025484 - 6232025484
// LOCAL: 40041010 - 0800123456
