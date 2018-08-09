package classify

// Tipo
const (
	VALIDO       = 1
	NAO_VALIDO   = 2
	NAO_VALIDADO = 3
)

// Classify struct
type Classify struct {
	Valido      []string `json:"valido"`
	NaoValido   []string `json:"nao_valido"`
	NaoValidado []string `json:"nao_validado"`
}

// New ...
func New() *Classify {
	return &Classify{}
}

// Add ...
func (c *Classify) Add(telefone string, tipo int) {
	if tipo == VALIDO {
		c.Valido = append(c.Valido, telefone)
	}

	if tipo == NAO_VALIDO {
		c.NaoValido = append(c.NaoValido, telefone)
	}

	if tipo == NAO_VALIDADO {
		c.NaoValidado = append(c.NaoValidado, telefone)
	}
}

// Remove ...
func (c *Classify) Remove(telefone string, tipo int) {
	if tipo == VALIDO {
		c.Valido = remove(c.Valido, telefone)
	}

	if tipo == NAO_VALIDO {
		c.NaoValido = remove(c.NaoValido, telefone)
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
