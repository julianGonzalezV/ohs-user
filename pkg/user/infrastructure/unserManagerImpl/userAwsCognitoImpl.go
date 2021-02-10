package userManagerImpl

import (
	"context"
	"fmt"
	"log"
	"ohs-user/pkg/user/domain/dependency"
	"ohs-user/pkg/user/domain/entity"
	"ohs-user/shared/customerror"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type userAwsCognito struct {
	db *mongo.Client
}

// NewRepository creates a mongo repository with the necessary dependencies
func New(db *mongo.Client) dependency.UserManagement {
	return userAwsCognito{db: db}
}

// Create saves a new record
func (r userAwsCognito) Create(ctx context.Context, p *entity.User) error {
	collection := r.db.Database("test").Collection("users")
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
func (r userAwsCognito) GetSession(ID string) (*entity.User, error) {
	fmt.Println("FetchByID", ID)
	collection := r.db.Database("test").Collection("users")
	resultStruct := &entity.User{}
	result := collection.FindOne(context.TODO(), bson.M{"sku": ID})
	if result.Err() == mongo.ErrNoDocuments {
		return nil, customerror.ErrRecordNotFound
	}
	result.Decode(&resultStruct)
	return resultStruct, nil

}

// Fetch return all records saved in storage
func (r userAwsCognito) ChangePassword(ID string) ([]*entity.User, error) {
	collection := r.db.Database("test").Collection("users")
	var results []*entity.User
	cur, error := collection.Find(context.TODO(), bson.M{"businessid": ID})
	if error != nil {
		return nil, customerror.ErrMongo
	}

	for cur.Next(context.TODO()) {
		//Value into which the single document can be decoded
		var elem entity.User
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
