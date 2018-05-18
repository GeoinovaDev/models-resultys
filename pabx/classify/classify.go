package classify

import (
	"git.resultys.com.br/motor/models/empresa"
	"gopkg.in/mgo.v2/bson"
)

// Classify struct
type Classify struct {
	ClassifyID bson.ObjectId  `json:"classify_id" bson:"_id,omitempty"`
	ID         string         `bson:"id"`
	Empresa    *empresa.Basic `json:"empresa" bson:"empresa"`

	Pertence    []string `json:"pertence" bson:"pertence"`
	NaoPertence []string `json:"nao_pertence" bson:"nao_pertence"`
	Contador    []string `json:"contador" bson:"contador"`
	Valido      []string `json:"valido" bson:"valido"`
	NaoValidado []string `json:"nao_validado" bson:"nao_validado"`
}

// New ...
func New() Classify {
	return Classify{}
}

// ExistTelefoneConfirmado ...
func (classify *Classify) ExistTelefoneConfirmado() bool {
	if len(classify.Pertence) > 0 {
		return true
	}

	return false
}

// ExistTelefoneValido ...
func (classify *Classify) ExistTelefoneValido() bool {
	if len(classify.Valido) > 0 {
		return true
	}

	return false
}
