package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jp-chl/test-go-clean-architecture/domain/model"
	"github.com/jp-chl/test-go-clean-architecture/domain/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	ErrRedirectNotFound = errors.New("redirect Not Found")
)

type mongoRepository struct {
	client   *mongo.Client
	database string
	timeout  time.Duration
}

func newMongoClient(mongoURL string, mongoTimeout int) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(mongoTimeout)*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected...")
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	fmt.Println("Ping ok...")
	return client, nil
}

func NewMongoRepository(mongoURL, mongoDB string, mongoTimeout int) (repository.RedirectRepository, error) {
	repository := &mongoRepository{
		timeout:  time.Duration(mongoTimeout) * time.Second,
		database: mongoDB,
	}
	client, err := newMongoClient(mongoURL, mongoTimeout)
	if err != nil {
		return nil, errors.New("unable to connect to Mongo")
	}
	repository.client = client
	return repository, nil
}

func (r *mongoRepository) Find(code string) (*model.Redirect, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	result := &model.Redirect{}
	collection := r.client.Database(r.database).Collection("redirects")
	filter := bson.M{"code": code}
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, ErrRedirectNotFound
		}
		return nil, errors.New("error while trying to find element")
	}
	return result, nil
}

func (r *mongoRepository) Store(redirect *model.Redirect) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection("redirects")
	_, err := collection.InsertOne(
		ctx,
		bson.M{
			"code":       redirect.Code,
			"url":        redirect.URL,
			"created_at": redirect.CreatedAt,
		},
	)
	if err != nil {
		fmt.Printf("error while trying to store [%s]", err.Error())
		return errors.New("error while trying to store")
	}
	return nil
}
