package coleta

import (
	"sync"

	"git.resultys.com.br/motor/models/email"
	"git.resultys.com.br/motor/models/empresa"
	"git.resultys.com.br/motor/models/facebook"
	"git.resultys.com.br/motor/models/linkedin"
	"git.resultys.com.br/motor/models/site"
	"git.resultys.com.br/motor/models/telefone"
	"git.resultys.com.br/motor/models/twitter"
)

// Coleta dados
type Coleta struct {
	Empresa empresa.Empresa

	Emails    []*email.Email
	Telefones []*telefone.Telefone
	Facebooks []*facebook.Page
	Linkedins []*linkedin.Linkedin
	Twitters  []*twitter.Twitter
	Sites     []*site.Site

	emailMutex    sync.Mutex
	telefoneMutex sync.Mutex
	facebookMutex sync.Mutex
	linkedinMutex sync.Mutex
	twitterMutex  sync.Mutex
	siteMutex     sync.Mutex
}

// PopuleFacebook facebook
func (c *Coleta) PopuleFacebook(facebooks []*facebook.Page) {
	c.facebookMutex.Lock()

	for i := 0; i < len(facebooks); i++ {
		c.Facebooks = append(c.Facebooks, facebooks[i])
	}

	c.facebookMutex.Unlock()
}

// PopuleEmail popula coletor com emails
func (c *Coleta) PopuleEmail(emails []*email.Email) {
	c.emailMutex.Lock()

	for i := 0; i < len(emails); i++ {
		c.Emails = append(c.Emails, emails[i])
	}

	c.emailMutex.Unlock()
}

// PopuleTelefone popule coletor com telefones
func (c *Coleta) PopuleTelefone(telefones []*telefone.Telefone) {
	c.telefoneMutex.Lock()

	for i := 0; i < len(telefones); i++ {
		c.Telefones = append(c.Telefones, telefones[i])
	}

	c.telefoneMutex.Unlock()
}

// PopuleSite popule coletor com telefones
func (c *Coleta) PopuleSite(sites []*site.Site) {
	c.siteMutex.Lock()

	for i := 0; i < len(sites); i++ {
		c.Sites = append(c.Sites, sites[i])
	}

	c.siteMutex.Unlock()
}

// PopuleLinkedin popule coletor com telefones
func (c *Coleta) PopuleLinkedin(linkedins []*linkedin.Linkedin) {
	c.linkedinMutex.Lock()

	for i := 0; i < len(linkedins); i++ {
		c.Linkedins = append(c.Linkedins, linkedins[i])
	}

	c.linkedinMutex.Unlock()
}

// PopuleTwitter popule coletor com telefones
func (c *Coleta) PopuleTwitter(twitters []*twitter.Twitter) {
	c.twitterMutex.Lock()

	for i := 0; i < len(twitters); i++ {
		c.Twitters = append(c.Twitters, twitters[i])
	}

	c.twitterMutex.Unlock()
}

func popule(mutex *sync.Mutex, arr []interface{}, dados []interface{}) {
	mutex.Lock()

	for i := 0; i < len(dados); i++ {
		arr = append(arr, dados[i])
	}

	mutex.Unlock()
}
