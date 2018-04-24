package facebook

// Page struct
type Page struct {
	Nome     string `json:"nome"`
	Title    string `json:"title"`
	URL      string `json:"url"`
	Endereco string `json:"endereco"`
	Curtidas string `json:"curtidas"`
	Horarios string `json:"horarios"`

	Telefones []string `json:"telefones"`
	Sites     []string `json:"sites"`
	Emails    []string `json:"emails"`
}

// PopuleFromMap page
func (page *Page) PopuleFromMap(m map[string]interface{}) {

	telefones := m["telefones"].([]interface{})
	emails := m["emails"].([]interface{})
	sites := m["sites"].([]interface{})

	if val, ok := m["nome"]; ok {
		page.Nome = val.(string)
	}

	if val, ok := m["title"]; ok {
		page.Title = val.(string)
	}

	if val, ok := m["url"]; ok {
		page.URL = val.(string)
	}

	if val, ok := m["endereco"]; ok {
		page.Endereco = val.(string)
	}

	if val, ok := m["curtidas"]; ok {
		page.Curtidas = val.(string)
	}

	if val, ok := m["horarios"]; ok {
		page.Horarios = val.(string)
	}

	for i := 0; i < len(telefones); i++ {
		page.Telefones = append(page.Telefones, telefones[i].(string))
	}

	for i := 0; i < len(emails); i++ {
		page.Emails = append(page.Emails, emails[i].(string))
	}

	for i := 0; i < len(sites); i++ {
		page.Sites = append(page.Sites, sites[i].(string))
	}
}
