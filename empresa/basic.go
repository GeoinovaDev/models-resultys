package empresa

import "strings"

// Basic struct
type Basic struct {
	RazaoSocial string
	Fantasia    string
	CNPJ        string
}

// GetNome ...
func (basic *Basic) GetNome() string {
	nome := ""

	if len(basic.Fantasia) == 0 {
		if len(basic.RazaoSocial) > 0 {
			partes := strings.Split(basic.RazaoSocial, " ")
			if len(partes) >= 2 {
				nome = partes[0] + " " + partes[1]
			} else {
				nome = partes[0]
			}
		}
	} else {
		nome = basic.Fantasia
	}

	return nome
}
