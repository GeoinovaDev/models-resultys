package token

import (
	"sync"

	"git.resultys.com.br/motor/models/empresa"
	"git.resultys.com.br/motor/models/telefone"
	"gopkg.in/mgo.v2/bson"
)

// Status
const (
	RUNNING = 1
	DONE    = 2
)

// Token dados
type Token struct {
	TokenID     bson.ObjectId       `json:"token_id" bson:"_id,omitempty"`
	ID          string              `json:"id" bson:"id"`
	CNPJ        string              `json:"cnpj" bson:"cnpj"`
	RazaoSocial string              `json:"razao_social" bson:"razao_social"`
	Fantasia    string              `json:"fantasia" bson:"fantasia"`
	Cidade      string              `json:"cidade" bson:"cidade"`
	Estado      string              `json:"estado" bson:"estado"`
	Webhook     string              `json:"webhook" bson:"webhook"`
	Status      int                 `json:"status" bson:"status"`
	Trigger     bool                `json:"trigger" bson:"trigger"`
	Telefones   []telefone.Telefone `json:"telefones" bson:"telefones"`

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
	return empresa.GetNome(e.RazaoSocial, e.Fantasia)
}
