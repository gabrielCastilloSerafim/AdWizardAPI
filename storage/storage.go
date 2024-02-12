package storage

import (
	"context"

	"github.com/gabrielCastilloSerafim/AdWizardAPI/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoStorage struct {
	MongoClient        *mongo.Client
	Context            *context.Context
	AppUserCollection  *mongo.Collection
	CampaignCollection *mongo.Collection
	EventCollection    *mongo.Collection
}

type StorageInterface interface {
	// Delete
	DeleteAllAppUsers() error
	DeleteAllEvents() error
	DeleteAllCampaigns() error
	// Read
	GetAllAppUsers() ([]models.AppUser, error)
	GetAppUserWithIp(ip string) (*models.AppUser, error)
	GetAllEvents() ([]models.Event, error)
	GetEventByCampaignId(campaignId string) (*models.Event, error)
	GetAllCampaigns() ([]models.Campaign, error)
	GetCampaignById(campaignId string) (string, error)
	// Create
	CreateEvent(event *models.Event) error
	CreateCampaign(campaign *models.Campaign) (string, error)
	CreateAppUser(appUser *models.AppUser) error
	// Update
}

func InitMongoStorage(mongoClient *mongo.Client) MongoStorage {
	context := context.Background()
	return MongoStorage{
		MongoClient:        mongoClient,
		Context:            &context,
		AppUserCollection:  mongoClient.Database("production").Collection("AppUser"),
		CampaignCollection: mongoClient.Database("production").Collection("Campaign"),
		EventCollection:    mongoClient.Database("production").Collection("Event"),
	}
}
