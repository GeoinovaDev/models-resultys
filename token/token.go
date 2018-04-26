package token

import (
	"strings"
	"sync"

	"gopkg.in/mgo.v2/bson"
)

// RUNNING
const (
	RUNNING = 1
	DONE    = 2
)

// Token dados
type Token struct {
	TokenID  bson.ObjectId `json:"token_id" bson:"_id,omitempty"`
	ID       string        `json:"id" bson:"id"`
	CNPJ     string        `json:"cnpj" bson:"cnpj"`
	Nome     string        `json:"nome" bson:"nome"`
	Fantasia string        `json:"fantasia" bson:"fantasia"`
	Cidade   string        `json:"cidade" bson:"cidade"`
	Estado   string        `json:"estado" bson:"estado"`
	Webhook  string        `json:"webhook" bson:"webhook"`
	Status   int           `json:"status" bson:"status"`
	Trigger  bool          `json:"trigger" bson:"trigger"`

	mutex *sync.Mutex
}

// forca commit

// New cria o token
func New() *Token {
	return &Token{
		mutex: &sync.Mutex{},
	}
}

// Lock trava o token
func (e *Token) Lock() {
	e.mutex.Lock()
}

// Unlock destrava o token
func (e *Token) Unlock() {
	e.mutex.Unlock()
}

// GetNome retorna nome da empresa
// Return string
func (e *Token) GetNome() string {
	nome := e.Fantasia
	if len(nome) == 0 {
		nome = e.Nome
	}

	return clearNome(nome)
}

func clearNome(nome string) string {
	nome = strings.Replace(nome, ".", "", -1)
	nome = strings.Replace(nome, "- ME", "", -1)
	nome = strings.Replace(nome, "LTDA", "", -1)
	nome = strings.Replace(nome, "EPP", "", -1)
	nome = strings.Replace(nome, "S/S", "", -1)
	nome = strings.Replace(nome, "S/C", "", -1)
	nome = strings.Replace(nome, "EIRELI", "", -1)
	nome = strings.Replace(nome, "-", "", -1)
	nome = strings.Replace(nome, `"`, "", -1)
	nome = strings.Replace(nome, `\`, "", -1)
	nome = strings.Replace(nome, `/`, "", -1)
	nome = strings.Replace(nome, `\\`, "", -1)
	nome = strings.Replace(nome, `(`, "", -1)
	nome = strings.Replace(nome, `)`, "", -1)

	nome = removeTipoEmpresa(nome, "ME")
	nome = removeTipoEmpresa(nome, "LTDA")
	nome = removeTipoEmpresa(nome, "EPP")

	return strings.Trim(nome, " ")
}

func removeTipoEmpresa(nome string, tipo string) string {
	nomes := strings.Split(nome, " ")
	if nomes[len(nomes)-1] == tipo {
		return strings.Join(nomes[:len(nomes)-1], " ")
	}

	return strings.Join(nomes, " ")
}
