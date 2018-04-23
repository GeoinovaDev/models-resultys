package facebook

// Page struct
type Page struct {
	Nome     string
	Title    string
	URL      string
	Endereco string
	Curtidas string
	Horarios string

	Telefones []string
	Sites     []string
	Emails    []string
}

// PopuleFromMap page
func (page *Page) PopuleFromMap(m map[string]interface{}) {
	telefones := m["telefones"].([]interface{})
	emails := m["emails"].([]interface{})
	sites := m["sites"].([]interface{})

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
