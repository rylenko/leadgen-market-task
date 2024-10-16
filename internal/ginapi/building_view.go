package ginapi

import "github.com/rylenko/leadgen-market-task/internal/domain"

// Building JSON view to make responses.
type buildingView struct {
	id int64            `json:"id"`
	name string         `json:"name"`
	city string         `json:"city"`
	handoverYear uint64 `json:"handover_year"`
	floorsCount uint64  `json:"floors_count"`
}

func getBuildingView(building *domain.Building) *buildingView {
	return &buildingView{
		id: building.Id,
		name: building.Info.Name,
		city: building.Info.City,
		handoverYear: building.Info.HandoverYear,
		floorsCount: building.Info.FloorsCount,
	}
}
