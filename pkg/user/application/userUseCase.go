package application

import (
	"context"
	"ohs-user/pkg/user/domain/entity"
	"ohs-user/pkg/user/domain/service"
	"ohs-user/pkg/user/infrastructure/request"
)

// UserUseCaseInterface provides operations to be executed.
type UserUseCaseInterface interface {
	SignUp(ctx context.Context, requestData request.UserRequest) error
	SignIn(ctx context.Context, clientID string) (*entity.User, error)
	ChangePassword(ctx context.Context, sku string) ([]*entity.User, error)
}

type userUseCase struct {
	service service.UserServiceInterface
}

// New creates the  application from Application Layer
func New(service service.UserServiceInterface) UserUseCaseInterface {
	return &userUseCase{service}
}

// Add adds the given record to storage
func (app *userUseCase) SignUp(ctx context.Context, requestData request.UserRequest) error {
	b := entity.New(itemsRequestToDomain(requestData.Items), requestData.Price, requestData.BusinessId, requestData.Sku, requestData.Name,
		requestData.Description, requestData.Category, requestData.State, requestData.ProductType, requestData.Image)
	return app.service.Create(ctx, b)

}

// GetByClient searches all records into the storage
func (app *userUseCase) SignIn(ctx context.Context, businessID string) (*entity.User, error) {
	return app.service.GetSession(ctx, businessID)
}

// Get searches all records into the storage
func (app *userUseCase) ChangePassword(ctx context.Context, sku string) ([]*entity.User, error) {
	return app.service.ChangePassword(ctx, sku)
}

func itemsRequestToDomain(items []request.ItemRequest) []entity.Item {
	var results []entity.Item
	for _, item := range items {
		results = append(results, entity.Item{Name: item.Name,
			Description: item.Description, Category: item.Category, State: item.State})
	}
	return results
}
