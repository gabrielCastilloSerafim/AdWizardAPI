package storage

import "go.mongodb.org/mongo-driver/bson"

func (db MongoStorage) DeleteAllAppUsers() error {
	_, err := db.AppUserCollection.DeleteMany(*db.Context, bson.M{})
	return err
}

func (db MongoStorage) DeleteAllEvents() error {
	_, err := db.EventCollection.DeleteMany(*db.Context, bson.M{})
	return err
}

func (db MongoStorage) DeleteAllCampaigns() error {
	_, err := db.CampaignCollection.DeleteMany(*db.Context, bson.M{})
	return err
}
