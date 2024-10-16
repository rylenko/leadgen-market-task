package logic

// BuildingFilters contains parameters that are used to select certain
// buildings from among the others.
//
// Filter values ​​are pointers. If one of them is nil, then the filter
// parameter is not set.
type BuildingFilters struct {
	City *string
	HandoverYear *uint64
	FloorsCount *uint64
}

// NewBuildingFilters creates a new instance of filter parameters
// structure.
func NewBuildingFilters(
		city *string, handoverYear, floorsCount *uint64) *BuildingFilters {
	return &BuildingFilters{
		City: city,
		HandoverYear: handoverYear,
		FloorsCount: floorsCount,
	}
}
