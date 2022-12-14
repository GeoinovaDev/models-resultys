package snovio

import (
	"github.com/GeoinovaDev/lower-resultys/db/mongo"
	mgo "gopkg.in/mgo.v2"
)

// Insert insere snovio model
func (snovio *Snovio) Insert() *Snovio {
	mongo.New().DB("vendor").C("snovio").Query(func(c *mgo.Collection) {
		err := c.Insert(snovio)
		if err != nil {
			panic(err)
		}
	})

	return snovio
}

// Fetch o atual apikey
func Fetch() *Snovio {
	snovio := &Snovio{}

	mongo.New().DB("vendor").C("snovio").Query(func(c *mgo.Collection) {
		c.Find(nil).One(&snovio)
	})

	return snovio
}
