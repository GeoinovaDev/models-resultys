package token

import (
	"sync"

	"github.com/GeoinovaDev/lower-resultys/time/datetime"
	"github.com/GeoinovaDev/models-resultys/empresa"
	"github.com/GeoinovaDev/models-resultys/telefone"
	"gopkg.in/mgo.v2/bson"
)

// Status
const (
	RUNNING = 1
	DONE    = 2
)

type diagnostic struct {
	StartTime string `json:"starttime" bson:"starttime"`
	EndTime   string `json:"endtime" bson:"endtime"`
}

// Token dados
type Token struct {
	TokenID     bson.ObjectId       `json:"token_id" bson:"_id,omitempty"`
	ID          string              `json:"id" bson:"id"`
	CNPJ        string              `json:"cnpj" bson:"cnpj"`
	RazaoSocial string              `json:"razao_social" bson:"razao_social"`
	Fantasia    string              `json:"fantasia" bson:"fantasia"`
	Cidade      string              `json:"cidade" bson:"cidade"`
	Estado      string              `json:"estado" bson:"estado"`
	CEP         string              `json:"cep" bson:"cep"`
	Webhook     string              `json:"webhook" bson:"webhook"`
	WebhookID   string              `json:"webhook_id" bson:"webhook_id"`
	Status      int                 `json:"status" bson:"status"`
	Trigger     bool                `json:"trigger" bson:"trigger"`
	Telefones   []telefone.Telefone `json:"telefones" bson:"telefones"`
	Domains     []string            `json:"domains" bson:"domains"`
	Params      map[string]string   `json:"params" bson:"params"`
	CreateAt    string              `json:"create_at" bson:"create_at"`
	Diagnostic  *diagnostic         `json:"diagnostic" bson:"diagnostic"`
	Latitude    string              `json:"latitude" bson:"latitude"`
	Longitude   string              `json:"longitude" bson:"longitude"`
	Worker      string              `json:"worker" bson:"worker"`

	mutex *sync.Mutex
}

// New cria o token
func New() *Token {
	return &Token{
		TokenID:  bson.NewObjectId(),
		mutex:    &sync.Mutex{},
		CreateAt: datetime.Now().String(),
		Diagnostic: &diagnostic{
			StartTime: datetime.Now().String(),
		},
		Params: map[string]string{},
	}
}

// AddParam ...
func (t *Token) AddParam(param string, value string) *Token {
	t.Params[param] = value

	return t
}

// GetParam ...
func (t *Token) GetParam(param string) string {
	if val, ok := t.Params[param]; ok {
		return val
	}

	return ""
}

// Lock trava o token
func (t *Token) Lock() {
	t.mutex.Lock()
}

// Unlock destrava o token
func (t *Token) Unlock() {
	t.mutex.Unlock()
}

// GetNome retorna nome da empresa
// Return string
func (t *Token) GetNome() string {
	return empresa.GetNome(t.RazaoSocial, t.Fantasia)
}
