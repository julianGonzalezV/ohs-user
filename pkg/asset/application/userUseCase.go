package application

import (
	"context"
	"ms-asset/pkg/asset/domain/entity"
	"ms-asset/pkg/asset/domain/service"
	"ms-asset/pkg/asset/infrastructure/request"
)

// UserUseCaseInterface provides operations to be executed.
type UserUseCaseInterface interface {
	Add(ctx context.Context, requestData request.AssetRequest) error
	GetByClient(ctx context.Context, clientId string) ([]*entity.Asset, error)
	Get(ctx context.Context, sku string) (*entity.Asset, error)
}

type assetUseCase struct {
	service service.AssetServiceInterface
}

// NewBusinessApp creates the business application from App Layer
func New(service service.AssetServiceInterface) UserUseCaseInterface {
	return &assetUseCase{service}
}

// Add adds the given record to storage
func (app *assetUseCase) Add(ctx context.Context, requestData request.AssetRequest) error {
	b := entity.New(itemsRequestToDomain(requestData.Items), requestData.Price, requestData.BusinessId, requestData.Sku, requestData.Name,
		requestData.Description, requestData.Category, requestData.State, requestData.ProductType, requestData.Image)
	return app.service.Add(ctx, b)

}

// GetByClient searches all records into the storage
func (app *assetUseCase) GetByClient(ctx context.Context, businessId string) ([]*entity.Asset, error) {
	return app.service.GetByClient(ctx, businessId)
}

// Get searches all records into the storage
func (app *assetUseCase) Get(ctx context.Context, sku string) (*entity.Asset, error) {
	return app.service.Get(ctx, sku)
}

func itemsRequestToDomain(items []request.ItemRequest) []entity.Item {
	var results []entity.Item
	for _, item := range items {
		results = append(results, entity.Item{Name: item.Name,
			Description: item.Description, Category: item.Category, State: item.State})
	}
	return results
}
