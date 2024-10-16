package pgx

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rylenko/leadgen-market-task/internal/domain"
	"github.com/rylenko/leadgen-market-task/internal/logic"
)

const (
	createCityIndexStatement = ```
		CREATE INDEX IF NOT EXISTS building_city_index
			ON building (city);
	```

	createFloorsCountIndexStatement = ```
		CREATE INDEX IF NOT EXISTS building_floors_count_index
			ON building (floors_count)
	```

	createHandoverYearIndexStatement = ```
		CREATE INDEX IF NOT EXISTS building_handover_year_index
			ON building (handover_year);
	```

	createTableStatement = ```
		CREATE TABLE IF NOT EXISTS building (
			id INTEGER PRIMARY KEY,
			name TEXT NOT NULL,
			city TEXT NOT NULL,
			handover_year INTEGER NOT NULL,
			floors_count INTEGER NOT NULL
		);
	```

	insertQuery = ```
		INSERT INTO building (name, city, handover_year, floors_count)
			VALUES ($1, $2, $3, $4) RETURNING (id);
	```
)

// BuildingRepositoryImpl is a pgx implementation of buildings repository.
type BuildingRepositoryImpl struct {
	pool *pgxpool.Pool
}

func (repository *BuildingRepositoryImpl) GetAll(
		filterParams *logic.BuildingFilterParams) (domain.Buildings, error) {

}

// Init creates database table and indexes if they are not exists.
func (repository *BuildingRepositoryImpl) Init(ctx context.Context) error {
	// Try to create database table.
	if err := repository.createTable(ctx); err != nil {
		return fmt.Errorf("failed to create table: %v", err)
	}

	// Try to create index on city field.
	if err := repository.createCityIndex(ctx); err != nil {
		return fmt.Errorf("failed to create city index: %v", err)
	}

	// Try to create index on handover year field.
	if err := repository.createHandoverYearIndex(ctx); err != nil {
		return fmt.Errorf("failed to create handover year index: %v", err)
	}

	// Try to create index on floors count field.
	if err := repository.createFloorsCountIndex(ctx); err != nil {
		return fmt.Errorf("failed to create floors count index: %v", err)
	}

	return nil
}

// Insert inserts a new building to the database.
func (repository *BuildingRepositoryImpl) Insert(
		building *domain.Building) (id int64, err error) {
	// Execute insertion query in the database.
	row := repository.pool.QueryRow(
		ctx,
		insertQuery,
		building.Name,
		building.City,
		building.HandoverYear,
		building.FloorsCount)

	// Scan returned id of a new building in the database.
	var id int64
	if _, err := row.Scan(&id); err != nil {
		return id, fmt.Errorf("failed to scan id of a new building: %v", err)
	}

	return id, nil
}

// Creates city index in the database.
func (repository *BuildingRepositoryImpl) createCityIndex(
		ctx context.Context) error {
	_, err := repository.pool.Exec(ctx, createCityIndexStatement)
	return err
}

// Creates floors count index in the database.
func (repository *BuildingRepositoryImpl) createFloorsCountIndex(
		ctx context.Context) error {
	_, err := repository.pool.Exec(ctx, createFloorsCountIndexStatement)
	return err
}

// Creates handover year index in the database.
func (repository *BuildingRepositoryImpl) createHandoverYearIndex(
		ctx context.Context) error {
	_, err := repository.pool.Exec(ctx, createHandoverYearIndexStatement)
	return err
}

// Creates buildings table in the database.
func (repository *BuildingRepositoryImpl) createTable(
		ctx context.Context) error {
	_, err := repository.pool.Exec(ctx, createTableStatement)
	return err
}

// Creates a new instance of pgx building repository implementation.
func NewBuildingRepositoryImpl(pool *pgxpool.Pool) *BuildingRepositoryImpl {
	return &BuildingRepositoryImpl{
		pool: pool,
	}
}
