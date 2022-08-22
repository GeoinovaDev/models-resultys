package luminati

import (
	"github.com/GeoinovaDev/lower-resultys/db/mongo"
	mgo "gopkg.in/mgo.v2"
)

// Insert insere luminati model
func (luminati *Luminati) Insert() *Luminati {
	mongo.New().DB("vendor").C("luminati").Query(func(c *mgo.Collection) {
		err := c.Insert(luminati)
		if err != nil {
			panic(err)
		}
	})

	return luminati
}

// FetchAll ...
func FetchAll() []Luminati {
	luminati := []Luminati{}

	mongo.New().DB("vendor").C("luminati").Query(func(c *mgo.Collection) {
		c.Find(nil).All(&luminati)
	})

	return luminati
}
