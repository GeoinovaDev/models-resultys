package empresa

// Basic struct
type Basic struct {
	RazaoSocial string
	Fantasia    string
	CNPJ        string
}

// GetNome ...
func (basic *Basic) GetNome() string {
	return GetNome(basic.RazaoSocial, basic.Fantasia)
}
