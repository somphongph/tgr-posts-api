package repositories

import (
	"context"
	"tgr-posts-api/modules/posts/domains"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Storer interface {
	GetById(string) (domains.Post, error)
	GetAll() ([]domains.Post, error)
	Add(*domains.Post) error
	Update(*domains.Post) error
	Delete(string) error
}

type MongoDBStore struct {
	*mongo.Collection
}

func InitMongoDBStore() *MongoDBStore {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(viper.GetString("MONGO_CONNECTION")))
	if err != nil {
		panic("failed to connect database")
	}
	collection := client.Database(viper.GetString("MONGO_DB_NAME")).Collection("posts")

	return &MongoDBStore{Collection: collection}
}

func (s *MongoDBStore) GetById(in string) (domains.Post, error) {
	id, _ := primitive.ObjectIDFromHex(in)

	var (
		ctx    = context.Background()
		filter = bson.M{"_id": id}
		result domains.Post
	)

	// Find
	err := s.Collection.FindOne(ctx, filter).Decode(&result)

	return result, err
}

func (s *MongoDBStore) GetAll() ([]domains.Post, error) {

	var (
		ctx    = context.Background()
		filter = bson.M{}
		result []domains.Post
	)

	// Find All
	cursor, err := s.Collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var item domains.Post
		cursor.Decode(&item)
		result = append(result, item)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return result, err
}

func (s *MongoDBStore) Add(p *domains.Post) error {
	var ctx = context.Background()

	// Insert
	_, err := s.Collection.InsertOne(ctx, p)
	return err
}

func (s *MongoDBStore) Update(p *domains.Post) error {
	update := bson.M{
		"$set": p,
	}
	var ctx = context.Background()

	// Update
	_, err := s.Collection.UpdateByID(ctx, p.Id, update)
	return err
}

func (s *MongoDBStore) Delete(in string) error {
	id, _ := primitive.ObjectIDFromHex(in)

	var (
		ctx    = context.Background()
		filter = bson.M{"_id": id}
	)

	// Delete
	_, err := s.Collection.DeleteOne(ctx, filter)
	return err
}
