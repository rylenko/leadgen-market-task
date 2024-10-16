package logic

import (
	"context"

	"github.com/rylenko/leadgen-market-task/internal/domain"
)

// BuildingRepository is an interface that describes the required capabilities
// of the building repository.
type BuildingRepository interface {
	// GetAll must get all buildings according to the passed filter parameters or
	// return an error.
	GetAll(
		ctx context.Context,
		filterParams *BuildingFilterParams) ([]*domain.Building, error)

	// Insert must insert a structure to the repository or return an error.
	Insert(
		ctx context.Context, info *domain.BuildingInfo) (*domain.Building, error)

	// Init must initialize repository before queries.
	Init(ctx context.Context) error
}
