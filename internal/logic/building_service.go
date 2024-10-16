package logic

import "github.com/rylenko/leadgen-market-task/internal/domain"

// BuildingService is an interface that describes the required capabilities of
// the building service.
type BuildingService interface {
	// Create must create a structure within the system or return an error. For
	// example, insert into the repository.
	Create(building *domain.Building) error

	// GetAll must get all buildings according to the passed filter parameters or
	// return an error.
	GetAll(filterParams *BuildingFilterParams) (domain.Buildings, error)

	// Init must initialize service before work.
	Init() error
}
