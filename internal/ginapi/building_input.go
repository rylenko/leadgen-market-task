package ginapi

import "github.com/rylenko/leadgen-market-task/internal/domain"

// Input JSON model to create or update buildings.
type buildingInput struct {
	Name string         `json:"name" binding:"required"`
	City string         `json:"city" binding:"required"`
	HandoverYear uint64 `json:"handover_year" binding:"required"`
	FloorsCount uint64  `json:"floors_count" binding:"required"`
}

// Converts JSON input to building information domain model.
func (input *buildingInput) toInfo() *domain.BuildingInfo {
	return domain.NewBuildingInfo(
		input.Name, input.City, input.HandoverYear, input.FloorsCount)
}
