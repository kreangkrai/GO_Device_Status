package Models

import "go.mongodb.org/mongo-driver/bson/primitive"

type DataModel struct {
	ID     primitive.ObjectID `json:"_id" bson:"_id"`
	Device string             `json:"device" bson:"device"`
	Date   string             `json:"date" bson:"date"`
	Status string             `json:"status" bson:"status"`
}
