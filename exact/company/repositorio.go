package company

import (
	"github.com/GeoinovaDev/lower-resultys/db/mongo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Save ...
func (company *Company) Save() *Company {
	mongo.New().DB("exact").C("company").Query(func(c *mgo.Collection) {
		c.Remove(bson.M{"cnpj": company.CNPJ})
		err := c.Insert(company)
		if err != nil {
			panic(err)
		}
	})

	return company
}

// Load ...
func Load() []*Company {
	companies := []*Company{}

	mongo.New().DB("exact").C("company").Query(func(c *mgo.Collection) {
		c.Find(nil).All(&companies)
	})

	return companies
}
