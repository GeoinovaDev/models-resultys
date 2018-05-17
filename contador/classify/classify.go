package classify

// Tipo
const (
	CONTADOR     = 1
	NAO_CONTADOR = 2
	NAO_VALIDADO = 3
)

// Classify struct
type Classify struct {
	Contador    []string `json:"contador"`
	NaoContador []string `json:"nao_contador"`
	NaoValidado []string `json:"nao_validado"`
}

// New ...
func New() *Classify {
	return &Classify{}
}

// Add ...
func (c *Classify) Add(telefone string, tipo int) {
	if tipo == CONTADOR {
		c.Contador = append(c.Contador, telefone)
	}

	if tipo == NAO_CONTADOR {
		c.NaoContador = append(c.NaoContador, telefone)
	}

	if tipo == NAO_VALIDADO {
		c.NaoValidado = append(c.NaoValidado, telefone)
	}
}

// Remove ...
func (c *Classify) Remove(telefone string, tipo int) {
	if tipo == CONTADOR {
		c.Contador = remove(c.Contador, telefone)
	}

	if tipo == NAO_CONTADOR {
		c.NaoContador = remove(c.NaoContador, telefone)
	}

	if tipo == NAO_VALIDADO {
		c.NaoValidado = remove(c.NaoValidado, telefone)
	}
}

func remove(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}
