package token

import (
	"git.resultys.com.br/lib/lower/db/mongo"
	"gopkg.in/mgo.v2"
)

// Insert token
func Insert(token *Token) *Token {
	mongo.New().DB("coletor").C("token").Query(func(c *mgo.Collection) {
		err := c.Insert(token)
		if err != nil {
			panic(err)
		}
	})

	return token
}
