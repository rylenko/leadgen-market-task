package ginapi

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/rylenko/leadgen-market-task/internal/ginapi/docs"
	"github.com/rylenko/leadgen-market-task/internal/logic"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title                                     Leadgen Market Task
// @version                                   1.0
// @description                               Swagger API of Leadgen Market Task
// @termsOfService                            http://swagger.io/terms/

// @contact.name                              rylenko
// @contact.url                               http://github.com/rylenko
// @contact.email	                      rylenko@tuta.io

// @license.name                              Apache 2.0
// @license.url                               http://www.apache.org/licenses/LICENSE-2.0.html

// @host                                      localhost:8000
// @BasePath                                  /api/v1

// @securityDefinitions.basic                 BasicAuth

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
	// addMiddlewares(engine)

	// Add v1 API controllers.
	v1group := engine.Group("/api/v1")
	addBuildingController(v1group, ctx, buildingService)

	// Add swagger controller.
	addSwaggerController(engine)

	log.Printf("before run")

	// Run engine on accepted address components.
	return engine.Run(addr...)
}

// Registers building handlers to the passed group.
func addBuildingController(
		group *gin.RouterGroup,
		ctx context.Context,
		service logic.BuildingService) {
	// Create a new instance of the controller.
	controller := NewBuildingController(ctx, service)

	// Create buildings sub-group and add controller handlers to it.
	buildings := group.Group("/buildings")
	{
		buildings.GET("", controller.GetAll)
		buildings.POST("", controller.Create)
	}
}

// Adds all middlewares to the passed engine.
// func addMiddlewares(engine *gin.Engine) {
	// engine.Use(printErrorsMiddleware)
// }

// Adds swagger controller to the engine.
func addSwaggerController(engine *gin.Engine) {
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

// Middleware to log context errors.
// func printErrorsMiddleware(c *gin.Context) {
	// c.Next()

	// for _, err := range c.Errors {
		// log.Printf("Error: %v", err)
	// }
// }
