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
	Fetch(filter primitive.M, sort primitive.D, page, limit int) ([]entities.Post, error)
	Count(filter primitive.M) (int64, error)
	Add(*entities.Post) error
	Update(*entities.Post) error
	Delete(string) error
}

type mongoDBStore struct {
	*mongo.Collection
}

func InitPostRepository(cfg *configs.MongoDB) *mongoDBStore {
	ctx := context.Background()
	opts := options.Client().ApplyURI(cfg.Connection)

	// Create a new client and connect to the server
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}

	collection := client.Database(cfg.DbName).Collection(tableName)

	return &mongoDBStore{Collection: collection}
}

func (s *mongoDBStore) GetById(in string) (entities.Post, error) {
	id, _ := primitive.ObjectIDFromHex(in)

	ctx := context.Background()
	filter := bson.M{"_id": id}
	var result entities.Post

	// Find
	err := s.Collection.FindOne(ctx, filter).Decode(&result)

	return result, err
}

func (s *mongoDBStore) Fetch(filter primitive.M, sort primitive.D, page, limit int) ([]entities.Post, error) {
	ctx := context.Background()

	l := int64(limit)
	skip := int64(page*limit - limit)
	opts := options.FindOptions{Limit: &l, Skip: &skip, Sort: &sort}

	// Find
	cursor, err := s.Collection.Find(ctx, &filter, &opts)
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

func (s *mongoDBStore) Count(filter primitive.M) (int64, error) {
	ctx := context.Background()

	opts := options.Count().SetHint("_id_")

	// Count
	count, err := s.CountDocuments(ctx, filter, opts)
	if err != nil {
		return 0, err
	}

	return count, err
}

func (s *mongoDBStore) Add(p *entities.Post) error {
	ctx := context.Background()

	p.Status = "active"
	p.CreatedBy = util.GetUserId()
	p.CreatedOn = time.Now()
	p.UpdatedBy = util.GetUserId()
	p.UpdatedOn = time.Now()

	// Insert
	_, err := s.Collection.InsertOne(ctx, p)
	return err
}

func (s *mongoDBStore) Update(p *entities.Post) error {
	ctx := context.Background()
	update := bson.M{
		"$set": p,
	}

	p.UpdatedBy = util.GetUserId()
	p.UpdatedOn = time.Now()

	// Update
	_, err := s.Collection.UpdateByID(ctx, p.Id, update)
	return err
}

func (s *mongoDBStore) Delete(in string) error {
	id, _ := primitive.ObjectIDFromHex(in)

	ctx := context.Background()
	filter := bson.M{"_id": id}

	// Delete
	_, err := s.Collection.DeleteOne(ctx, filter)
	return err
}
