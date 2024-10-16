package domain

// BuildingInfo places information about buildings.
type BuildingInfo struct {
	Name string
	City string
	HandoverYear uint64
	FloorsCount uint64
}

// NewBuildingInfo creates a new instance of building information.
func NewBuildingInfo(
		name, city string, handoverYear, floorsCount uint64) *BuildingInfo {
	return &BuildingInfo{
		Name: name,
		City: city,
		HandoverYear: handoverYear,
		FloorsCount: floorsCount,
	}
}
