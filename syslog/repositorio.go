package syslog

import (
	"git.resultys.com.br/lib/lower/db/mongo"
	mgo "gopkg.in/mgo.v2"
)

// Insert ...
func Insert(log interface{}) {
	mongo.New().DB("syslog").C("log").Query(func(c *mgo.Collection) {
		err := c.Insert(log)
		if err != nil {
			panic(err)
		}
	})
}
