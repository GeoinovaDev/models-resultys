package email

import (
	"strings"
)

// Email dados
type Email struct {
	Nome     string `json:"nome" bson:"nome"`
	Email    string `json:"email" bson:"email"`
	CreateAt string `json:"create_at" bson:"create_at"`
	Status   string `json:"status" bson:"status"`
	Fonte    int    `json:"fonte" bson:"fonte"`
	Ranking  int    `json:"ranking" bson:"ranking"`
	Repeat   int    `json:"repeat" bson:"repeat"`
}

// New ...
func New(email string) *Email {
	return &Email{Email: email}
}

// Raw ...
func (e Email) Raw() string {
	return e.Email
}

// GetDomain ...
func (e *Email) GetDomain() string {
	return ExtractDomain(e.Email)
}

// GetName ...
func (e *Email) GetName() string {
	return ExtractName(e.Email)
}

// ExtractName ...
func ExtractName(email string) string {
	p := strings.Split(email, "@")

	if len(p) > 1 {
		return p[0]
	}

	return ""
}

// ExtractDomain ...
func ExtractDomain(email string) string {
	p := strings.Split(email, "@")

	if len(p) > 1 {
		return p[1]
	}

	return ""
}
