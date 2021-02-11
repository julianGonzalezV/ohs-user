package service

// This component is in charge of handle domain business rules
import (
	"context"
	"ohs-user/pkg/user/domain/dependency"
	"ohs-user/pkg/user/domain/entity"
)

// UserServiceInterface interface that establishes functions to be implemented
type UserServiceInterface interface {
	Create(ctx context.Context, b *entity.User) error
	GetSession(ctx context.Context, businessID string) (*entity.User, error)
	ChangePassword(ctx context.Context, sku string) ([]*entity.User, error)
}

type userService struct {
	userManagerDependency dependency.UserManager
}

// New creates a domain service for core logic
func New(dependency dependency.UserManager) UserServiceInterface {
	return &userService{dependency}
}

// Create a new user
func (service *userService) Create(ctx context.Context, p *entity.User) error {
	service.userManagerDependency.Create(p)
	return nil
}

// GetSession searches all records related to a business
func (service *userService) GetSession(ctx context.Context, businessID string) (*entity.User, error) {
	return service.userManagerDependency.GetSession(businessID)
}

// ChangePassword searches a record
func (service *userService) ChangePassword(ctx context.Context, sku string) ([]*entity.User, error) {
	return service.userManagerDependency.ChangePassword(sku)
}
