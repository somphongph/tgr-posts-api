package posts

import (
	"context"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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

func (s *MongoDBStore) GetById(pid string) (Post, error) {
	id, _ := primitive.ObjectIDFromHex(pid)

	var (
		ctx    = context.Background()
		filter = bson.M{"_id": id}
		result Post
	)

	// Find
	err := s.Collection.FindOne(ctx, filter).Decode(&result)

	return result, err
}

func (s *MongoDBStore) GetAll() ([]Post, error) {

	var (
		ctx    = context.Background()
		filter = bson.M{}
		result []Post
	)

	// Find All
	cursor, err := s.Collection.Find(ctx, filter)
	defer cursor.Close(ctx)

	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var item Post
		cursor.Decode(&item)
		result = append(result, item)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return result, err
}

func (s *MongoDBStore) Add(p *Post) error {
	var ctx = context.Background()

	// Insert
	_, err := s.Collection.InsertOne(ctx, p)
	return err
}

func (s *MongoDBStore) Update(p *Post) error {
	update := bson.M{
		"$set": p,
	}
	var ctx = context.Background()

	// Update
	_, err := s.Collection.UpdateByID(ctx, p.Id, update)
	return err
}

func (s *MongoDBStore) Delete(pid string) error {
	id, _ := primitive.ObjectIDFromHex(pid)

	var (
		ctx    = context.Background()
		filter = bson.M{"_id": id}
	)

	// Delete
	_, err := s.Collection.DeleteOne(ctx, filter)
	return err
}
