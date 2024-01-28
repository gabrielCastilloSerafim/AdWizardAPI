package main

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

type Event struct {
	Id         primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Name       string             `bson:"name" json:"name"`
	CampaignId string             `bson:"campaignId" json:"campaignId"`
	CreatedAt  time.Time          `bson:"createdAt" json:"createdAt"`
	UpdateAt   time.Time          `bson:"updateAt" json:"updateAt"`
	UserId     string             `bson:"userId" json:"userId"`
}

type AppUser struct {
	Id         primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	CampaignId string             `bson:"campaignId" json:"campaignId"`
	UserIp     string             `bson:"userIp" json:"userIp"`
	Email      string             `bson:"email" json:"email"`
	CreatedAt  time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt  time.Time          `bson:"updatedAt" json:"updatedAt"`
}
