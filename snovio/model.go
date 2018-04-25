package snovio

import "gopkg.in/mgo.v2/bson"

// Snovio model
type Snovio struct {
	ID           bson.ObjectId `json:"id" bson:"_id,omitempty"`
	ClientID     string        "json:`client_id` bson:`client_id`"
	ClientSecret string        "json:`client_secret` bson:`client_secret`"
}
