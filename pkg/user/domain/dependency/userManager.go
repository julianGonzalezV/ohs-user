package dependency

import (
	"ohs-user/pkg/user/domain/entity"
)

// UserManager interface that establishes functions to be implemented
type UserManager interface {
	// Create saves a new record
	Create(p *entity.User) error
	// GetSession returns the record with given ID
	GetSession(ID string) (*entity.User, error)
	// ChangePassword returns the record with given business ID
	ChangePassword(businessID string) ([]*entity.User, error)
}
