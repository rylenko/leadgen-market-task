package ginapi

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rylenko/leadgen-market-task/internal/domain"
	"github.com/rylenko/leadgen-market-task/internal/logic"
)

// Controller to handle building routes.
type BuildingController struct {
	ctx context.Context
	service logic.BuildingService
}

// Create godoc
//
// @Summary     Creates a new building
// @Description Creates a new building using passed data
// @ID          create-building
// @Tags        building
// @Accept      json
// @Produce     json
// @Success     201                                       {object} BuildingView
// @Failure     400                                       {object} Error
// @Failure     500                                       {object} Error
// @Router      /buildings                                [post]
func (controller *BuildingController) Create(c *gin.Context) {
	var data BuildingInput

	// Try to bind creaetion data to structure.
	if err := c.ShouldBindJSON(&data); err != nil {
		NewError(http.StatusBadRequest, err.Error()).Push(c)
		return
	}

	// Use service to create a new building.
	building, err := controller.service.Create(controller.ctx, data.toInfo())
	if err != nil {
		c.Error(err)
		NewError(http.StatusInternalServerError, "internal error").Push(c)
		return
	}

	// Convert building domain model to JSON building view and make response with
	// it.
	view := getBuildingView(building)
	c.JSON(http.StatusCreated, view)
}

// Gets all buildings from service repository according to query filters.
func (controller *BuildingController) GetAll(c *gin.Context) {
	// Try to extract building filters from context.
	filters, err := extractFilters(c)
	if err != nil {
		NewError(http.StatusBadRequest, err.Error()).Push(c)
		return
	}

	// Try to get all buildings.
	buildings, err := controller.service.GetAll(controller.ctx, filters)
	if err != nil {
		c.Error(err)
		NewError(http.StatusInternalServerError, "internal error").Push(c)
		return
	}

	// Convert building models to JSON view.
	var views []*BuildingView
	for _, building := range buildings {
		views = append(views, getBuildingView(building))
	}

	// Make response with building views.
	c.JSON(http.StatusOK, views)
}

// Creates a new building controller.
func NewBuildingController(
		ctx context.Context, service logic.BuildingService) *BuildingController {
	return &BuildingController{
		ctx: ctx,
		service: service,
	}
}

// Input JSON model to create or update buildings.
type BuildingInput struct {
	Name string         `json:"name" binding:"required"`
	City string         `json:"city" binding:"required"`
	HandoverYear uint64 `json:"handover_year" binding:"required"`
	FloorsCount uint64  `json:"floors_count" binding:"required"`
}

// Converts JSON input to building information domain model.
func (input *BuildingInput) toInfo() *domain.BuildingInfo {
	return domain.NewBuildingInfo(
		input.Name, input.City, input.HandoverYear, input.FloorsCount)
}

// Building JSON view to make responses.
type BuildingView struct {
	Id int64            `json:"id"`
	Name string         `json:"name"`
	City string         `json:"city"`
	HandoverYear uint64 `json:"handover_year"`
	FloorsCount uint64  `json:"floors_count"`
}

// Gets building view from building domain model.
func getBuildingView(building *domain.Building) *BuildingView {
	return &BuildingView{
		Id: building.Id,
		Name: building.Info.Name,
		City: building.Info.City,
		HandoverYear: building.Info.HandoverYear,
		FloorsCount: building.Info.FloorsCount,
	}
}

// Extracts building filter values from passed context or returns an error if
// at least one query contains invalid value.
func extractFilters(c *gin.Context) (*logic.BuildingFilters, error) {
	// Parse city filter.
	cityFilter := extractStringFilter(c, "city")

	// Parse handover year filter.
	handoverYearFilter, err := extractUInt64Filter(c, "handover_year")
	if err != nil {
		return nil, err
	}

	// Parse floors count filter.
	floorsCountFilter, err := extractUInt64Filter(c, "floors_count")
	if err != nil {
		return nil, err
	}

	return logic.NewBuildingFilters(
		cityFilter, handoverYearFilter, floorsCountFilter), nil
}

func extractStringFilter(c *gin.Context, name string) *string {
	value := c.Query(name)
	if value == "" {
		return nil
	}
	return &value
}

// Extracts uint64 filter by its name from passed context or returns an error
// if query contains invalid value.
func extractUInt64Filter(c *gin.Context, name string) (*uint64, error) {
	// Get query string from context.
	str := c.Query(name)
	if str == "" {
		return nil, nil
	}

	// Try to parse uint64 value from query string.
	value, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("filter %s is not uint64 type", name)
	}
	return &value, nil
}
