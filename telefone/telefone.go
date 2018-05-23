package telefone

import (
	"strings"

	"git.resultys.com.br/lib/lower/str"
	"git.resultys.com.br/motor/models/telefone/tipo"
)

// Telefone dados
type Telefone struct {
	Numero   string `json:"numero" bson:"numero"`
	Pais     string `json:"pais" bson:"pais"`
	DDD      string `json:"ddd" bson:"ddd"`
	CreateAt string `json:"create_at" bson:"create_at"`
	Fonte    int    `json:"fonte" bson:"fonte"`
	Tipo     int    `json:"tipo" bson:"tipo"`
	Ranking  int    `json:"ranking" bson:"ranking"`
	Repeat   int    `json:"repeat" bson:"repeat"`
}

// Format ...
func (telefone Telefone) Format() string {
	formatted := ""

	if telefone.Tipo == tipo.LOCAL {
		formatted = telefone.Raw()
	}

	if telefone.Tipo == tipo.GRATUITO {
		formatted = telefone.Raw()
	}

	if telefone.Tipo == tipo.FIXO {
		formatted = str.Format("({0}) {1}-{2}", telefone.DDD, telefone.Numero[:4], telefone.Numero[4:])
	}

	if telefone.Tipo == tipo.CELULAR {
		formatted = str.Format("({0}) {1}-{2}", telefone.DDD, telefone.Numero[:5], telefone.Numero[5:])
	}

	return formatted
}

// Raw retorna somente o numero do telefone completo
func (telefone Telefone) Raw() string {
	return telefone.Pais + telefone.DDD + telefone.Numero
}

// ReFormat ...
func (telefone Telefone) ReFormat() Telefone {
	t := New(telefone.Raw())

	t.Fonte = telefone.Fonte
	t.Tipo = telefone.Tipo
	t.CreateAt = telefone.CreateAt
	t.Ranking = telefone.Ranking

	return t
}

// RawSemDDI ...
func (telefone Telefone) RawSemDDI() string {
	return telefone.DDD + telefone.Numero
}

// New Cria telefone a partir do numero
// Return Telefone
func New(numero string) Telefone {
	telefone := Telefone{}

	numero = Clear(numero)

	// TELEFONE LOCAL
	if strings.Index(numero, "4003") == 0 {
		telefone.Tipo = tipo.LOCAL
		telefone.Pais = "55"
		telefone.DDD = ""
		telefone.Numero = numero

		return telefone
	}

	// TELEFONE LOCAL
	if strings.Index(numero, "4004") == 0 {
		telefone.Tipo = tipo.LOCAL
		telefone.Pais = "55"
		telefone.DDD = ""
		telefone.Numero = numero

		return telefone
	}

	// TELEFONE GRATUITO
	if strings.Index(numero, "0800") == 0 {
		telefone.Tipo = tipo.GRATUITO
		telefone.Pais = "55"
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
		telefone.Pais = "55"
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

	if len(numero) == 10 {

		// CELULAR ANTIGO COM DDD
		if string(numero[2]) == "9" || string(numero[2]) == "8" {
			telefone.Tipo = tipo.CELULAR
			telefone.Pais = "55"
			telefone.DDD = numero[:2]
			telefone.Numero = "9" + numero[2:]

			return telefone
		}

		// FIXO COM DDD
		telefone.Tipo = tipo.FIXO
		telefone.Pais = "55"
		telefone.DDD = numero[:2]
		telefone.Numero = numero[2:]

		return telefone
	}

	// FIXO ANTIGO COM DDD
	if len(numero) == 9 {
		telefone.Tipo = tipo.FIXO
		telefone.Pais = "55"
		telefone.DDD = numero[:2]
		telefone.Numero = "3" + numero[2:]
	}

	return telefone
}

// Clear ...
func Clear(numero string) string {
	numero = strings.Replace(numero, "-", "", -1)
	numero = strings.Replace(numero, "/", "", -1)
	numero = strings.Replace(numero, ")", "", -1)
	numero = strings.Replace(numero, "(", "", -1)
	numero = strings.Replace(numero, "+", "", -1)
	numero = strings.Replace(numero, " ", "", -1)

	return strings.Trim(numero, " ")
}

// EXEMPLOS DE NUMERO DE TELEFONES VALIDOS
// CELULAR: 5562982334440 - 62982334440
// FIXO: 556232025484 - 6232025484
// LOCAL: 40041010 - 0800123456
