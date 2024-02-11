package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AppUser struct {
	Id         primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	CampaignId string             `bson:"campaignId" json:"campaignId"`
	UserIp     string             `bson:"userIp" json:"userIp"`
	Email      string             `bson:"email" json:"email"`
	CreatedAt  time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt  time.Time          `bson:"updatedAt" json:"updatedAt"`
}

func (appUser *AppUser) MarshalBSON() ([]byte, error) {
	if appUser.CreatedAt.IsZero() {
		appUser.CreatedAt = time.Now()
	}
	appUser.UpdatedAt = time.Now()

	type my AppUser
	return bson.Marshal((*my)(appUser))
}
