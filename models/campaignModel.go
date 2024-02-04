package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Campaign struct {
	Id           primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Name         string             `bson:"name" json:"name"`
	Source       string             `bson:"source" json:"source"`
	AppStoreLink string             `bson:"appStoreLink" json:"appStoreLink"`
	CreatedAt    time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt    time.Time          `bson:"updatedAt" json:"updatedAt"`
}
