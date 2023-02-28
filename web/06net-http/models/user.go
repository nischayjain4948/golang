package models

import "gopkg.in/mgo.v2/bson"

type User struct {
	ID     bson.ObjectId `json:"id" bson:"_id"`        //while exporting
	Name   string        `json:"name" bson:"name"`     //while exporting
	Gender string        `json:"gender" bson:"gender"` //while exporting
	Age    int           `json:"age" bson:"age"`       //while exporting
}
