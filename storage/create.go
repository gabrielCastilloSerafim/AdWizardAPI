package storage

import (
	"errors"

	"github.com/gabrielCastilloSerafim/AdWizardAPI/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db MongoStorage) CreateEvent(event *models.Event) error {
	_, err := db.EventCollection.InsertOne(*db.Context, event)
	if err != nil {
		return errors.New("error creating event on database")
	}
	return nil
}

func (db MongoStorage) CreateCampaign(campaign *models.Campaign) (string, error) {
	insertResult, err := db.CampaignCollection.InsertOne(*db.Context, campaign)
	campaignId := (insertResult.InsertedID.(primitive.ObjectID)).Hex()
	if err != nil {
		return "", errors.New("error creating campaign in database")
	}
	return campaignId, nil
}

func (db MongoStorage) CreateAppUser(appUser *models.AppUser) error {
	_, err := db.AppUserCollection.InsertOne(*db.Context, appUser)
	if err != nil {
		return errors.New("error creating app user in database")
	}
	return nil
}
