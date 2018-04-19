package coleta

import (
	"sync"

	"git.resultys.com.br/prospecta/models/email"
	"git.resultys.com.br/prospecta/models/telefone"
)

var telefoneMutex *sync.Mutex
var emailMutex *sync.Mutex

func init() {
	if telefoneMutex == nil {
		telefoneMutex = &sync.Mutex{}
	}

	if emailMutex == nil {
		emailMutex = &sync.Mutex{}
	}
}

// Coleta dados
type Coleta struct {
	Emails    []email.Email
	Telefones []telefone.Telefone
}

// PopuleEmails popula coletor com emails
func (c *Coleta) PopuleEmails(emails []email.Email) {
	emailMutex.Lock()

	for i := 0; i < len(emails); i++ {
		c.Emails = append(c.Emails, emails[i])
	}

	emailMutex.Unlock()
}

// PopuleTelefones popule coletor com telefones
func (c *Coleta) PopuleTelefones(telefones []telefone.Telefone) {
	telefoneMutex.Lock()

	for i := 0; i < len(telefones); i++ {
		c.Telefones = append(c.Telefones, telefones[i])
	}

	telefoneMutex.Unlock()
}
