package luminati

import "gopkg.in/mgo.v2/bson"

// Luminati model
type Luminati struct {
	ID       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	User     string        `json:"user" bson:"user"`
	Password string        `json:"password" bson:"password"`
	Host     string        `json:"host" bson:"host"`
}
