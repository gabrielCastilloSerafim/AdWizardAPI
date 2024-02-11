package storage

import (
	"errors"

	"github.com/gabrielCastilloSerafim/AdWizardAPI/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db MongoStorage) GetAllAppUsers() ([]models.AppUser, error) {
	var appUsers []models.AppUser
	cursor, err := db.AppUserCollection.Find(*db.Context, bson.M{})
	if err != nil {
		return nil, err
	}
	err = cursor.All(*db.Context, &appUsers)
	if err != nil {
		return nil, err
	}
	return appUsers, nil
}

func (db MongoStorage) GetAppUserWithIp(ip string, appUser *models.AppUser) error {
	userMatch := db.AppUserCollection.FindOne(*db.Context, bson.M{"userIp": ip})
	userMatch.Decode(appUser)
	if userMatch != nil {
		return errors.New("could not find app user match for given id")
	}
	return nil
}

func (db MongoStorage) GetAllEvents() ([]models.Event, error) {
	cursor, err := db.EventCollection.Find(*db.Context, bson.M{})
	if err != nil {
		return nil, err
	}
	var events []models.Event
	err = cursor.All(*db.Context, &events)
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (db MongoStorage) GetEventByCampaignId(campaignId string) (*models.Event, error) {
	cursor, err := db.EventCollection.Find(*db.Context, bson.M{"campaignId": campaignId})
	if err != nil {
		return nil, err
	}
	var events []models.Event
	err = cursor.All(*db.Context, &events)
	if err != nil {
		return nil, err
	}
	if len(events) == 0 {
		return nil, errors.New("could not find event for given campaing id")
	}
	return &events[0], nil
}

func (db MongoStorage) GetAllCampaigns() ([]models.Campaign, error) {
	cursor, err := db.CampaignCollection.Find(*db.Context, bson.M{})
	if err != nil {
		return nil, err
	}
	var campaings []models.Campaign
	err = cursor.All(*db.Context, &campaings)
	if err != nil {
		return nil, err
	}
	return campaings, nil
}

func (db MongoStorage) GetCampaignById(campaignId string) (string, error) {
	objId, err := primitive.ObjectIDFromHex(campaignId)
	if err != nil {
		return "", err
	}
	campaignMatch := db.CampaignCollection.FindOne(*db.Context, bson.M{"_id": objId})
	campaign := new(models.Campaign)
	campaignMatch.Decode(campaign)
	return campaign.AppStoreLink, nil
}
