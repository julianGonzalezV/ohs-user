package repositoryimpl

import (
	"context"
	"fmt"
	"log"
	"ms-asset/pkg/asset/domain/entity"
	"ms-asset/pkg/asset/domain/repository"
	"ms-asset/shared/customerror"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type assetRepository struct {
	db *mongo.Client
}

// NewRepository creates a mongo repository with the necessary dependencies
func New(db *mongo.Client) repository.AssetRepository {
	return assetRepository{db: db}
}

// Create saves a new record
func (r assetRepository) Create(ctx context.Context, p *entity.Asset) error {
	collection := r.db.Database("test").Collection("assets")
	// Insert a single document
	insertResult, err := collection.InsertOne(context.TODO(), p)
	if err != nil {
		fmt.Println("Error insertando", err)
		log.Fatal(err)
	}
	fmt.Println("insertResult", insertResult.InsertedID)
	return nil
}

// FetchByID returns the record related to given ID
func (r assetRepository) FetchByID(ID string) (*entity.Asset, error) {
	fmt.Println("FetchByID", ID)
	collection := r.db.Database("test").Collection("assets")
	resultStruct := &entity.Asset{}
	result := collection.FindOne(context.TODO(), bson.M{"sku": ID})
	if result.Err() == mongo.ErrNoDocuments {
		return nil, customerror.ErrRecordNotFound
	}
	result.Decode(&resultStruct)
	return resultStruct, nil

}

// Fetch return all records saved in storage
func (r assetRepository) FetchByClient(ID string) ([]*entity.Asset, error) {
	collection := r.db.Database("test").Collection("assets")
	var results []*entity.Asset
	cur, error := collection.Find(context.TODO(), bson.M{"businessid": ID})
	if error != nil {
		return nil, customerror.ErrMongo
	}

	for cur.Next(context.TODO()) {
		//Value into which the single document can be decoded
		var elem entity.Asset
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &elem)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	return results, nil
}
