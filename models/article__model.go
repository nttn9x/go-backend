package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Articles struct {
	ID               bson.ObjectId `bson:"_id,omitempty" json:"id"`
	LongDescription  string        `json:"LongDescription"`
	ShortDescription string        `json:"ShortDescription"`
}
