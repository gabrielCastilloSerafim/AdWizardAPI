package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Event struct {
	Id         primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Name       string             `bson:"name" json:"name"`
	CampaignId string             `bson:"campaignId" json:"campaignId"`
	CreatedAt  time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt  time.Time          `bson:"updateAt" json:"updatedAt"`
	UserId     string             `bson:"userId" json:"userId"`
}

func (event *Event) MarshalBSON() ([]byte, error) {
	if event.CreatedAt.IsZero() {
		event.CreatedAt = time.Now()
	}
	event.UpdatedAt = time.Now()

	type my Event
	return bson.Marshal((*my)(event))
}
