package domain

import "github.com/rylenko/leadgen-market-task/internal/domain"

// BuildingFilterParams contains parameters that are used to select certain
// buildings from among the others.
//
// Filter values ​​are pointers. If one of them is nil, then the filter
// parameter is not set.
type BuildingFilterParams struct {
	City *string
	HandoverYear *uint16
	FloorsCount *uint16
}

// NewBuildingFilterParams creates a new instance of filter parameters
// structure.
//
//
func NewBuildingFilterParams(
		city *string, handoverYear, floorsCount *uint16) *BuildingFilterParams {
	return &BuildingFilterParams{
		City: city,
		HandoverYear: handoverYear,
		FloorsCount: floorsCount,
	}
}
