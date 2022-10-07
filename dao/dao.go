package dao

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/njupt-sast/atsast-apply-module-server/config"
)

var Client *mongo.Client
var UserColl *mongo.Collection
var ExamColl *mongo.Collection
var ConfigColl *mongo.Collection
var InvitationColl *mongo.Collection

const timeLimit = 5 * time.Second

func init() {
	ctx, cancel := context.WithTimeout(context.Background(), timeLimit)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.Database.Uri))
	if err != nil {
		panic(err)
	}

	// check database connection
	Client = client
	if err = Client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	// create unique index for user collection
	UserColl = client.Database(config.Database.Name).Collection("user")
	if _, err = UserColl.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.D{{Key: "userId", Value: 1}},
		Options: options.Index().SetUnique(true).SetName("userId"),
	}); err != nil {
		panic(err)
	}

	// create unique index for exam collection
	ExamColl = client.Database(config.Database.Name).Collection("exam")
	if _, err = ExamColl.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.D{{Key: "examId", Value: 1}},
		Options: options.Index().SetUnique(true).SetName("examId"),
	}); err != nil {
		panic(err)
	}

	ConfigColl = client.Database(config.Database.Name).Collection("config")

	InvitationColl = client.Database(config.Database.Name).Collection("invitation")
}
