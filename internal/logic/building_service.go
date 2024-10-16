package logic

import (
	"context"

	"github.com/rylenko/leadgen-market-task/internal/domain"
)

// BuildingService is an interface that describes the required capabilities of
// the building service.
type BuildingService interface {
	// Create must create a structure within the system or return an error. For
	// example, insert into the repository.
	Create(
		ctx context.Context, building *domain.BuildingInfo) (*domain.Building, error)

	// GetAll must get all buildings according to the passed filter parameters or
	// return an error.
	GetAll(
		ctx context.Context, filters *BuildingFilters) ([]*domain.Building, error)

	// Init must initialize service before work.
	Init(ctx context.Context) error
}
