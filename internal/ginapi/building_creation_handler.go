package ginapi

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rylenko/leadgen-market-task/internal/logic"
)

func createBuildingsCreationHandler(
		ctx context.Context, service logic.BuildingService) gin.HandlerFunc {
	return func (c *gin.Context) {
		var data buildingInput

		// Try to bind creaetion data to structure.
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Use service to create a new building.
		building, err := service.Create(ctx, data.toInfo())
		if err != nil {
			c.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
			return
		}

		// Convert building domain model to JSON building view and make response with
		// it.
		view := getBuildingView(building)
		c.JSON(http.StatusCreated, view)
	}
}
