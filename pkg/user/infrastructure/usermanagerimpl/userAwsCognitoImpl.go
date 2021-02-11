package usermanagerimpl

import (
	"fmt"
	"ohs-user/pkg/user/domain/dependency"
	"ohs-user/pkg/user/domain/entity"
)

type userAwsCognito struct {
	//db *mongo.Client
}

// New  create a User manager
func New() dependency.UserManager {
	return userAwsCognito{}
}

// Create a new user
func (r userAwsCognito) Create(p *entity.User) error {
	fmt.Println("Creating user")
	return nil
}

// FetchByID returns the record related to given ID
func (r userAwsCognito) GetSession(ID string) (*entity.User, error) {
	fmt.Println("GetSession")
	resultStruct := &entity.User{}
	return resultStruct, nil

}

// Fetch return all records saved in storage
func (r userAwsCognito) ChangePassword(ID string) ([]*entity.User, error) {
	var results []*entity.User
	return results, nil
}
