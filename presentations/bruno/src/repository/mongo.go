package repository

import (
	"catalog/domain"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type mongoRepository struct {
	client  *mongo.Client
	db      string
	timeout time.Duration
}

func (r *mongoRepository) Find(code string) (*domain.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	product := &domain.Product{}
	collection := r.client.Database(r.db).Collection("products")
	filter := bson.M{"code": code}
	err := collection.FindOne(ctx, filter).Decode(product)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, domain.ErrProductNotFound
		}
		return nil, fmt.Errorf("repository search error: %w", err)
	}
	return product, nil
}

func (r *mongoRepository) Store(product *domain.Product) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.db).Collection("products")
	_, err := collection.InsertOne(
		ctx,
		bson.M{
			"code":  product.Code,
			"name":  product.Name,
			"price": product.Price,
		},
	)
	if err != nil {
		return fmt.Errorf("error writing to repository: %w", err)
	}
	return nil
}

func (r *mongoRepository) Update(product *domain.Product) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.db).Collection("products")
	_, err := collection.UpdateOne(
		ctx,
		bson.M{"code": product.Code},
		bson.D{
			{Key: "$set", Value: bson.D{{Key: "name", Value: product.Name}, {Key: "price", Value: product.Price}}},
		},
		&options.UpdateOptions{
			Upsert: &[]bool{false}[0],
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *mongoRepository) FindAll() ([]*domain.Product, error) {
	var items []*domain.Product

	collection := r.client.Database(r.db).Collection("products")
	cur, err := collection.Find(context.Background(), bson.D{})

	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())
	for cur.Next(context.TODO()) {

		var item domain.Product
		if err := cur.Decode(&item); err != nil {
			log.Fatal(err)
			return nil, err
		}
		items = append(items, &item)
	}
	return items, nil
}

func (r *mongoRepository) Delete(code string) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	filter := bson.M{"code": code}

	collection := r.client.Database(r.db).Collection("products")
	_, err := collection.DeleteOne(ctx, filter)

	if err != nil {
		return err
	}

	return nil
}

func newMongoClient(mongoServerURL string, timeout int) (*mongo.Client, error) {
	clientOptions := options.Client().
		ApplyURI(mongoServerURL)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	return client, nil
}

// NewMongoRepository  mongoDB adapter is created by constructing a mongo repository based on the mongoRepository struct
func NewMongoRepository(mongoServerURL, mongoDb string, timeout int) (domain.Repository, error) {
	mongoClient, err := newMongoClient(mongoServerURL, timeout)
	if err != nil {
		return nil, fmt.Errorf("client error: %w", err)
	}

	_, _ = mongoClient.Database(mongoDb).Collection("products").Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys: bson.M{
			"code": 1,
		},
		Options: &options.IndexOptions{
			Unique: &[]bool{true}[0],
		},
	})

	repo := &mongoRepository{
		client:  mongoClient,
		db:      mongoDb,
		timeout: time.Duration(timeout) * time.Second,
	}

	return repo, nil
}
