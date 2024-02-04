package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Event struct {
	Id         primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Name       string             `bson:"name" json:"name"`
	CampaignId string             `bson:"campaignId" json:"campaignId"`
	CreatedAt  time.Time          `bson:"createdAt" json:"createdAt"`
	UpdateAt   time.Time          `bson:"updateAt" json:"updateAt"`
	UserId     string             `bson:"userId" json:"userId"`
}
