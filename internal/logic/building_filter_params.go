package logic

// BuildingFilterParams contains parameters that are used to select certain
// buildings from among the others.
//
// Filter values ​​are pointers. If one of them is nil, then the filter
// parameter is not set.
type BuildingFilterParams struct {
	City *string
	HandoverYear *uint64
	FloorsCount *uint64
}

// NewBuildingFilterParams creates a new instance of filter parameters
// structure.
//
//
func NewBuildingFilterParams(
		city *string, handoverYear, floorsCount *uint64) *BuildingFilterParams {
	return &BuildingFilterParams{
		City: city,
		HandoverYear: handoverYear,
		FloorsCount: floorsCount,
	}
}
