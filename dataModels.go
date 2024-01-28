package main

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Campaign struct {
	Id           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string             `bson:"name" json:"name"`
	Source       string             `bson:"source" json:"source"`
	AppStoreLink string             `bson:"appStoreLink" json:"appStoreLink"`
	CreatedAt    time.Time          `bson:"createdAt"`
	UpdatedAt    time.Time          `bson:"updatedAt"`
}

type Event struct {
	Id         primitive.ObjectID `bson:"_id,omitempty"`
	Name       string             `bson:"name" json:"name"`
	CampaignId string             `bson:"campaignId" json:"campaignId"`
	CreatedAt  time.Time          `bson:"createdAt"`
	UpdateAt   time.Time          `bson:"updateAt"`
	UserId     string             `bson:"userId" json:"userId"`
}

type AppUser struct {
	Id         primitive.ObjectID `bson:"_id,omitempty"`
	CampaignId string             `bson:"campaignId"`
	UserIp     string             `bson:"userIp"`
	Email      string             `bson:"email"`
	CreatedAt  time.Time          `bson:"createdAt"`
	UpdatedAt  time.Time          `bson:"updatedAt"`
}
