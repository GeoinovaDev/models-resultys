package snovio

import (
	"git.resultys.com.br/lib/lower/db/mongo"
	mgo "gopkg.in/mgo.v2"
)

// Insert insere snovio model
func Insert(snovio *Snovio) *Snovio {
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
		err := c.Find(nil).One(&snovio)
		if err != nil {
			panic(err)
		}
	})

	return snovio
}
