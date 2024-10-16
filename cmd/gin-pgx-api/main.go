package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/rylenko/leadgen-market-task/internal/domain"
	"github.com/rylenko/leadgen-market-task/internal/logic"
	"github.com/rylenko/leadgen-market-task/internal/pgxrepo"
)

const (
	usage = "Usage: <config-path>"
	postgresqlURIFormat = "postgresql://%s:%s@%s:%d/%s"
)

type Config struct {
	Host string     `json:"host"`
	Port int        `json:"port"`
	User string     `json:"user"`
	Password string `json:"password"`
	Db string       `json:"db"`
}

// Builds database URI using parsed config parameters.
func (config *Config) buildURI() string {
	return fmt.Sprintf(
		postgresqlURIFormat,
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Db)
}

// Opens and parses config file using passed path.
func openConfig(path string) (*Config, error) {
	// Try to open config file.
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open config %s: %v", path, err)
	}
	defer file.Close()

	// Create JSON decoder of opened file.
	decoder := json.NewDecoder(file)

	// Try to decode JSON config to the structure.
	var config Config
	if err := decoder.Decode(&config); err != nil {
		return nil, fmt.Errorf("failed to decode JSON config %v", err)
	}

	return &config, nil
}

func main() {
	// Validate arguments count.
	if len(os.Args) != 2 {
		log.Fatal(usage)
	}

	// Try to open and parse config file.
	config, err := openConfig(os.Args[1])
	if err != nil {
		log.Fatalf("failed to open config file: %v", err)
	}

	// Open buildings repository.
	repository, err := pgx.OpenBuildingRepositoryImpl(
		context.Background(), config.buildURI())
	if err != nil {
		log.Fatalf("failed to open building repository: %v", err)
	}
	defer repository.Close()

	// Create a new instance of building service.
	service := logic.NewBuildingServiceImpl(repository)

	// Try to init build service.
	if err := service.Init(context.Background()); err != nil {
		log.Fatalf("failed to initialize build service: %v", err)
	}

	info := domain.NewBuildingInfo("info 333", "unknown", 2031, 1)
	building, err := service.Create(context.Background(), info)
	if err != nil {
		log.Fatalf("failed to create %+v: %v", info, err)
	}
	fmt.Printf("info created: %d, %+v\n", building.Id, building.Info)

	var v uint64 = 2031
	filterParams := logic.NewBuildingFilterParams(nil, &v, nil)
	buildings, err := service.GetAll(context.Background(), filterParams)
	if err != nil {
		log.Fatalf("failed to get all %+v: %v", filterParams, err)
	}
	for _, b := range buildings {
		fmt.Printf("building: %d, %+v\n", b.Id, b.Info)
	}
}
