package snovio

import (
	"git.resultys.com.br/lib/lower/db/mongo"
	"git.resultys.com.br/lib/lower/log"
	"git.resultys.com.br/lib/lower/net/loopback"
)

// Insert insere snovio model
func Insert(snovio *Snovio) *Snovio {
	mongo := mongo.New()
	if mongo == nil {
		return nil
	}
	defer mongo.Close()

	collection := mongo.DB("vendor").C("snovio")

	err := collection.Insert(snovio)
	if err != nil {
		log.Logger.Save(err.Error(), log.WARNING, loopback.IP())
		return nil
	}

	return snovio
}

// Fetch o atual apikey
func Fetch() *Snovio {
	mongo := mongo.New()
	if mongo == nil {
		return nil
	}
	defer mongo.Close()

	collection := mongo.DB("vendor").C("snovio")

	snovio := &Snovio{}
	collection.Find(nil).One(&snovio)

	return snovio
}
