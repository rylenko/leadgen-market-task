package main

import (
	"context"
	"fmt"
	"log"

	"github.com/rylenko/leadgen-market-task/internal/domain"
	"github.com/rylenko/leadgen-market-task/internal/logic"
	"github.com/rylenko/leadgen-market-task/internal/pgx"
)

func main() {
	uri := "postgresql://admin:adminpwd@db:5432/db"

	// Open buildings repository.
	repository, err := pgx.OpenBuildingRepositoryImpl(context.Background(), uri)
	if err != nil {
		log.Fatal("failed to open building repository", err)
	}
	defer repository.Close()

	// Create a new instance of building service.
	service := logic.NewBuildingServiceImpl(repository)

	// Try to init build service.
	if err := service.Init(context.Background()); err != nil {
		log.Fatal("failed to initialize build service", err)
	}

	info := domain.NewBuildingInfo("info 1", "Moscow", 2025, 9)
	building, err := service.Create(context.Background(), info)
	if err != nil {
		log.Fatal("failed to create %+v: %v", info, err)
	}
	fmt.Printf("info created: %+v\n", building)
}
