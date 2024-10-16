package logic

import (
	"context"
	"fmt"

	"github.com/rylenko/leadgen-market-task/internal/domain"
)

// BuildingService implementation that interacts with the repository to work
// with data.
type BuildingServiceImpl struct {
	repository BuildingRepository
}

// Create inserts passed building to the database or returns an error.
func (service *BuildingServiceImpl) Create(
		ctx context.Context, info *domain.BuildingInfo) (*domain.Building, error) {
	// Try to insert accepted building to the repository.
	building, err := service.repository.Insert(ctx, info)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to insert building to the repository: %v", err)
	}

	return building, nil
}

// GetAll gets all buildings according to the passed filter parameters or
// returns and error.
func (service *BuildingServiceImpl) GetAll(
		ctx context.Context, filters *BuildingFilters) ([]*domain.Building, error) {
	// Try to get all buildings using repository and accepted filter parameters.
	buildings, err := service.repository.GetAll(ctx, filters)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to get all buildings with filter params %+v: %v", filters, err)
	}

	return buildings, nil
}

// Initializes service before work. For example, initializes repository.
func (service *BuildingServiceImpl) Init(ctx context.Context) error {
	// Try to initialize service repository.
	if err := service.repository.Init(ctx); err != nil {
		return fmt.Errorf("failed to init repository: %v", err)
	}

	return nil
}

// NewBuildingServiceImpl creates a new instance of building service
// implementation using passed repository.
func NewBuildingServiceImpl(
		repository BuildingRepository) *BuildingServiceImpl {
	return &BuildingServiceImpl{
		repository: repository,
	}
}
