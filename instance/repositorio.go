package instance

import (
	"git.resultys.com.br/lib/lower/db/mongo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Insert instance
func (instance *Instance) Insert() *Instance {
	mongo.New().DB("compute").C("instances").Query(func(c *mgo.Collection) {
		err := c.Insert(instance)
		if err != nil {
			panic(err)
		}
	})

	return instance
}

// FetchByAlloc busca instancias pelo tipo alloc
func FetchByAlloc(alloc string) (all []Instance) {
	all = []Instance{}

	mongo.New().DB("compute").C("instances").Query(func(c *mgo.Collection) {
		err := c.Find(bson.M{"alloc": alloc}).All(&all)
		if err != nil {
			panic(err)
		}
	})

	return
}
