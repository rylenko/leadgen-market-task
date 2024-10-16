package ginapi

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rylenko/leadgen-market-task/internal/logic"
)

func createBuildingsGettingHandler(
		ctx context.Context, service logic.BuildingService) gin.HandlerFunc {
	return func (c *gin.Context) {
		// Try to extract building filters from context.
		filters, err := extractFilters(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Try to get all buildings.
		buildings, err := service.GetAll(ctx, filters)
		if err != nil {
			c.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal error"})
			return
		}

		// Convert building models to JSON view.
		var views []*buildingView
		for _, building := range buildings {
			views = append(views, getBuildingView(building))
		}

		// Make response with building views.
		c.JSON(http.StatusOK, views)
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
