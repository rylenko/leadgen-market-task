package domain

// Building places information about buildings.
type Building struct {
	Name string
	City string
	HandoverYear uint16
	FloorsCount uint16
}

// NewBuilding creates a new instance of Building structure.
func NewBuilding(
		name string, city string, handoverYear, floorsCount uint16) *Building {
	return &Building{
		Name: name,
		City: city,
		HandoverYear: handoverYear,
		FloorsCount: floorsCount,
	}
}
