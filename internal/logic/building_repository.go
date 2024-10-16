package logic

import "github.com/rylenko/leadgen-market-task/internal/domain"

// BuildingRepository is an interface that describes the required capabilities
// of the building repository.
type BuildingRepository interface {
	// GetAll must get all buildings according to the passed filter parameters or
	// return an error.
	GetAll(filterParams *BuildingFilterParams) (domain.Buildings, error)

	// Insert must insert a structure to the repository or return an error.
	Insert(building *domain.Building) (id int64, err error)

	// Init must initialize repository before queries.
	Init() error
}
