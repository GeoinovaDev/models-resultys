package company

import (
	"git.resultys.com.br/lib/lower/db/mongo"
	"gopkg.in/mgo.v2"
)

// Save ...
func (c *Company) Save() *Company {
	mongo.New().DB("exact").C("company").Query(func(c *mgo.Collection) {
		err := c.Insert(c)
		if err != nil {
			panic(err)
		}
	})

	return c
}

// Load ...
func Load() []*Company {
	companies := []*Company{}

	mongo.New().DB("exact").C("company").Query(func(c *mgo.Collection) {
		c.Find(nil).All(&companies)
	})

	return companies
}
