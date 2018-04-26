package token

import (
	"gopkg.in/mgo.v2/bson"

	"git.resultys.com.br/lib/lower/db/mongo"
	"gopkg.in/mgo.v2"
)

// Insert token
func (token *Token) Insert() *Token {
	mongo.New().DB("coletor").C("token").Query(func(c *mgo.Collection) {
		token.TokenID = bson.NewObjectId()
		err := c.Insert(token)
		if err != nil {
			panic(err)
		}
	})

	return token
}

// FetchByTokenID busca token pelo tokenid
// Return *token
func FetchByTokenID(id bson.ObjectId) *Token {
	token := &Token{}

	mongo.New().DB("coletor").C("token").Query(func(c *mgo.Collection) {
		err := c.Find(bson.M{"_id": id}).One(token)
		if err != nil {
			panic(err)
		}
	})

	return token
}

// UpdateStatus atualiza o status do token
func (token *Token) UpdateStatus(status int) {
	mongo.New().DB("coletor").C("token").Query(func(c *mgo.Collection) {
		err := c.Update(bson.M{"_id": token.TokenID}, bson.M{"$set": bson.M{"status": status}})
		if err != nil {
			panic(err)
		}
	})
}

// UpdateTrigger atualiza o status do token
func (token *Token) UpdateTrigger(trigger bool) {
	mongo.New().DB("coletor").C("token").Query(func(c *mgo.Collection) {
		err := c.Update(bson.M{"_id": token.TokenID}, bson.M{"$set": bson.M{"trigger": trigger}})
		if err != nil {
			panic(err)
		}
	})
}
