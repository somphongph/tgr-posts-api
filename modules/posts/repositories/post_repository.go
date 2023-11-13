package repositories

import (
	"context"
	"tgr-posts-api/configs"
	"tgr-posts-api/modules/posts/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	tableName = "posts"
)

type PostRepository interface {
	GetById(string) (entities.Post, error)
	GetAll() ([]entities.Post, error)
	Add(*entities.Post) error
	Update(*entities.Post) error
	Delete(string) error
}

type MongoDBStore struct {
	*mongo.Collection
}

func InitMongoDBStore(cfg *configs.MongoDB) *MongoDBStore {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(cfg.Connection))
	if err != nil {
		panic("failed to connect database")
	}
	collection := client.Database(cfg.DbName).Collection(tableName)

	return &MongoDBStore{Collection: collection}
}

func (s *MongoDBStore) GetById(in string) (entities.Post, error) {
	id, _ := primitive.ObjectIDFromHex(in)

	var (
		ctx    = context.Background()
		filter = bson.M{"_id": id}
		result entities.Post
	)

	// Find
	err := s.Collection.FindOne(ctx, filter).Decode(&result)

	return result, err
}

func (s *MongoDBStore) GetAll() ([]entities.Post, error) {

	var (
		ctx    = context.Background()
		filter = bson.M{}
		result []entities.Post
	)

	// Find All
	cursor, err := s.Collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var item entities.Post
		cursor.Decode(&item)
		result = append(result, item)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return result, err
}

func (s *MongoDBStore) Add(p *entities.Post) error {
	var ctx = context.Background()

	// Insert
	_, err := s.Collection.InsertOne(ctx, p)
	return err
}

func (s *MongoDBStore) Update(p *entities.Post) error {
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
