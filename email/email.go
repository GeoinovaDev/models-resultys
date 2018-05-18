package email

import (
	"strings"
)

// Email dados
type Email struct {
	Email    string `json:"email" bson:"email"`
	CreateAt string `json:"create_at" bson:"create_at"`
	Status   string `json:"status" bson:"status"`
	Fonte    int    `json:"fonte" bson:"fonte"`
	Ranking  int    `json:"ranking" bson:"ranking"`
}

// Raw ...
func (e Email) Raw() string {
	return e.Email
}

// IsOwner checa se o email eh de dominio privado
func IsOwner(email string) bool {
	b := true
	arr := []string{
		"globo", "gmail", "hotmail", "ig.com", "ibest", "oi.com", "live",
		"yahoo", "bol", "uol", "msn", "r7.com", "aol", "brturbo", "outlook", "terra",
		"hotamail", "hotail", "gamil", "outlok", "htmail", "ymail", "hotmaiul", "yhaoo",
	}

	for i := 0; i < len(arr); i++ {
		if strings.Index(email, arr[i]) > -1 {
			b = false
			break
		}
	}

	return b
}
