package repository

import (
	"context"
	"ohs-user/pkg/user/domain/entity"
)

// UserManagement interface that establishes functions to be implemented
type UserManagement interface {
	// Create saves a new record
	Create(ctx context.Context, p *entity.User) error
	// FetchByID returns the record with given ID
	GetSession(Id string) (*entity.User, error)
	// FetchByClient returns the record with given business ID
	ChangePassword(businessId string) ([]*entity.User, error)
}
