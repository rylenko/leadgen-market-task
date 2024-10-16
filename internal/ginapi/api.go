package ginapi

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rylenko/leadgen-market-task/internal/logic"
)

// Launches API using passed context, services and address components.
func Launch(
		ctx context.Context,
		buildingService logic.BuildingService,
		addr ...string) error {
	// Initialize services.
	if err := buildingService.Init(ctx); err != nil {
		fmt.Errorf("failed to initialize building service: %v", err)
	}

	// Create, fill engine with middlewares and handlers and run it.
	engine := gin.Default()
	useMiddlewares(engine)
	useHandlers(ctx, engine, buildingService)
	return engine.Run(addr...)
}

// Registers all handlers to the passed engine.
func useHandlers(
		ctx context.Context,
		engine *gin.Engine,
		buildingService logic.BuildingService) {
	// Building handlers.
	engine.GET("/buildings", createBuildingsGettingHandler(ctx, buildingService))
}

// Adds all middlewares to the passed engine.
func useMiddlewares(router *gin.Engine) {
	router.Use(printErrorsMiddleware)
}
