package instance

import "gopkg.in/mgo.v2/bson"

// Instance model
type Instance struct {
	ID    bson.ObjectId `json:"id" bson:"_id,omitempty"`
	IP    string        `json:"ip" bson:"ip"`
	Alloc string        `json:"alloc" bson:"alloc"`
}
