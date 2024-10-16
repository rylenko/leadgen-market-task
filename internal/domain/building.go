package domain

// Building places information about building and other parameters that allow
// us to distinguish between buildings.
type Building struct {
	Id int64
	Info *BuildingInfo
}

// NewBuilding creates a new instance of building structure.
func NewBuilding(id int64, info *BuildingInfo) *Building {
	return &Building{
		Id: id,
		Info: info,
	}
}
