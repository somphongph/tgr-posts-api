package repositories

import (
	"context"
	"tgr-posts-api/configs"
	"tgr-posts-api/modules/posts/entities"
	"tgr-posts-api/modules/shared/util"
	"time"

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
	Fetch(filter primitive.M, page, limit int) ([]entities.Post, error)
	Add(*entities.Post) error
	Update(*entities.Post) error
	Delete(string) error
}

type MongoDBStore struct {
	*mongo.Collection
}

func InitMongoDBStore(cfg *configs.MongoDB) *MongoDBStore {
	ctx := context.Background()
	opts := options.Client().ApplyURI(cfg.Connection)

	// Create a new client and connect to the server
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}

	collection := client.Database(cfg.DbName).Collection(tableName)

	return &MongoDBStore{Collection: collection}
}

func (s *MongoDBStore) GetById(in string) (entities.Post, error) {
	id, _ := primitive.ObjectIDFromHex(in)

	ctx := context.Background()
	filter := bson.M{"_id": id}
	var result entities.Post

	// Find
	err := s.Collection.FindOne(ctx, filter).Decode(&result)

	return result, err
}

func (s *MongoDBStore) Fetch(filter primitive.M, page, limit int) ([]entities.Post, error) {
	ctx := context.Background()

	l := int64(limit)
	skip := int64(page*limit - limit)
	opt := options.FindOptions{Limit: &l, Skip: &skip}

	// Find
	cursor, err := s.Collection.Find(ctx, &filter, &opt)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var result []entities.Post
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
	ctx := context.Background()

	p.Status = "active"
	p.CreatedBy = util.GetUserId()
	p.CreatedOn = time.Time{}
	p.UpdatedBy = util.GetUserId()
	p.UpdatedOn = time.Time{}

	// Insert
	_, err := s.Collection.InsertOne(ctx, p)
	return err
}

func (s *MongoDBStore) Update(p *entities.Post) error {
	ctx := context.Background()
	update := bson.M{
		"$set": p,
	}

	p.UpdatedBy = util.GetUserId()
	p.UpdatedOn = time.Time{}

	// Update
	_, err := s.Collection.UpdateByID(ctx, p.Id, update)
	return err
}

func (s *MongoDBStore) Delete(in string) error {
	id, _ := primitive.ObjectIDFromHex(in)

	ctx := context.Background()
	filter := bson.M{"_id": id}

	// Delete
	_, err := s.Collection.DeleteOne(ctx, filter)
	return err
}
