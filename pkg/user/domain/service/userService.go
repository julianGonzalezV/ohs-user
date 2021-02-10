package service

// This component is in charge of handle domain business rules
import (
	"context"
	"ohs-user/pkg/user/domain/dependency"
	"ohs-user/pkg/user/domain/entity"
	"ohs-user/shared/customerror"
)

// UserServiceInterface interface that establishes functions to be implemented
type UserServiceInterface interface {
	Create(ctx context.Context, b *entity.User) error
	GetSession(ctx context.Context, businessId string) ([]*entity.User, error)
	ChangePassword(ctx context.Context, sku string) (*entity.User, error)
}

type userService struct {
	userManagerDependency dependency.UserManagement
}

// New creates a domain service for core logic
func New(dependency dependency.UserManagement) UserServiceInterface {
	return &userService{dependency}
}

// Create a new user
func (service *UserService) Create(ctx context.Context, p *entity.User) error {
	_, error := service.userManagerDependency.Create(p.Sku)
	if error == customerror.ErrRecordNotFound {
		return service.userManagerDependency.Create(ctx, p)
	}
	return nil
}

// GetSession searches all records related to a business
func (service *UserService) GetSession(ctx context.Context, businessId string) ([]*entity.User, error) {
	return service.userManagerDependency.GetSession(businessId)
}

// ChangePassword searches a record
func (service *UserService) ChangePassword(ctx context.Context, sku string) (*entity.User, error) {
	return service.userManagerDependency.ChangePassword(sku)
}
