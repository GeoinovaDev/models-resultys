package instance

import (
	"git.resultys.com.br/lib/lower/db/mongo"
	"git.resultys.com.br/lib/lower/log"
	"git.resultys.com.br/lib/lower/net/loopback"
	"gopkg.in/mgo.v2/bson"
)

// Insert instance
func Insert(instance *Instance) *Instance {
	mongo := mongo.New()
	if mongo == nil {
		return nil
	}
	defer mongo.Close()

	c := mongo.DB("compute").C("instances")
	err := c.Insert(instance)
	if err != nil {
		log.Logger.Save(err.Error(), log.WARNING, loopback.IP())
		return nil
	}

	return instance
}

// FetchByAlloc busca instancias pelo tipo alloc
func FetchByAlloc(alloc string) []Instance {
	all := []Instance{}
	mongo := mongo.New()
	if mongo == nil {
		return all
	}
	defer mongo.Close()

	c := mongo.DB("compute").C("instances")
	err := c.Find(bson.M{"alloc": alloc}).All(&all)
	if err != nil {
		log.Logger.Save(err.Error(), log.WARNING, loopback.IP())
		return all
	}

	return all
}
