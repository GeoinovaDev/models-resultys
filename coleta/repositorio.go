package coleta

import (
	"git.resultys.com.br/lib/lower/db/mongo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Insert coleta
func (coleta *Coleta) Insert() *Coleta {
	mongo.New().DB("coletor").C("coleta").Query(func(c *mgo.Collection) {
		err := c.Insert(coleta)
		if err != nil {
			panic(err)
		}
	})

	return coleta
}

// Fetch busca coleta por id
func Fetch(id string) (coleta *Coleta) {
	coleta = &Coleta{}

	mongo.New().DB("coletor").C("coleta").Query(func(c *mgo.Collection) {
		err := c.Find(bson.M{"id": id}).One(coleta)
		if err != nil {
			coleta = nil
			panic(err)
		}
	})

	return
}
