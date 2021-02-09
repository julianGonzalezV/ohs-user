package service

// This component is in charge of handle domain business rules
import (
	"context"
	"ms-asset/pkg/asset/domain/entity"
	"ms-asset/pkg/asset/domain/repository"
	"ms-asset/shared/customerror"
)

// assetServiceInterface interface that establishes functions to be implemented
type AssetServiceInterface interface {
	Add(ctx context.Context, b *entity.Asset) error
	GetByClient(ctx context.Context, businessId string) ([]*entity.Asset, error)
	Get(ctx context.Context, sku string) (*entity.Asset, error)
}

type assetService struct {
	repository repository.AssetRepository
}

// New creates a domain service for core logic
func New(repository repository.AssetRepository) AssetServiceInterface {
	return &assetService{repository}
}

// Addasset adds the given record
func (service *assetService) Add(ctx context.Context, p *entity.Asset) error {
	_, error := service.repository.FetchByID(p.Sku)
	if error == customerror.ErrRecordNotFound {
		return service.repository.Create(ctx, p)
	}
	return nil
}

// GetassetsByBusiness searches all records related to a business
func (service *assetService) GetByClient(ctx context.Context, businessId string) ([]*entity.Asset, error) {
	return service.repository.FetchByClient(businessId)
}

// Getasset searches a record
func (service *assetService) Get(ctx context.Context, sku string) (*entity.Asset, error) {
	return service.repository.FetchByID(sku)
}
