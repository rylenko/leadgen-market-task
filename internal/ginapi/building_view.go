package ginapi

import "github.com/rylenko/leadgen-market-task/internal/domain"

// Building JSON view to make responses.
type buildingView struct {
	Id int64            `json:"id"`
	Name string         `json:"name"`
	City string         `json:"city"`
	HandoverYear uint64 `json:"handover_year"`
	FloorsCount uint64  `json:"floors_count"`
}

// Gets building view from building domain model.
func getBuildingView(building *domain.Building) *buildingView {
	return &buildingView{
		Id: building.Id,
		Name: building.Info.Name,
		City: building.Info.City,
		HandoverYear: building.Info.HandoverYear,
		FloorsCount: building.Info.FloorsCount,
	}
}
