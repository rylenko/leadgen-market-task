package pgx

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rylenko/leadgen-market-task/internal/domain"
	"github.com/rylenko/leadgen-market-task/internal/logic"
)

const (
	createCityIndexStatement = `
		CREATE INDEX IF NOT EXISTS building_city_index
			ON building USING HASH (city);
	`

	createFloorsCountIndexStatement = `
		CREATE INDEX IF NOT EXISTS building_floors_count_index
			ON building (floors_count);
	`

	createHandoverYearIndexStatement = `
		CREATE INDEX IF NOT EXISTS building_handover_year_index
			ON building (handover_year);
	`

	createTableStatement = `
		CREATE TABLE IF NOT EXISTS building (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			city TEXT NOT NULL,
			handover_year INTEGER NOT NULL CHECK (handover_year >= 0),
			floors_count INTEGER NOT NULL CHECK (floors_count >= 0)
		);
	`

	getAllQueryPrefix = `
		SELECT id, name, city, handover_year, floors_count FROM building
	`

	insertQuery = `
		INSERT INTO building (name, city, handover_year, floors_count)
			VALUES ($1, $2, $3, $4) RETURNING (id);
	`
)

// BuildingRepositoryImpl is a pgx implementation of buildings repository.
type BuildingRepositoryImpl struct {
	pool *pgxpool.Pool
}

// Closes opened repository implementation.
func (repository *BuildingRepositoryImpl) Close() {
	repository.pool.Close()
}

func (repository *BuildingRepositoryImpl) GetAll(
		ctx context.Context,
		filters *logic.BuildingFilters) ([]*domain.Building, error) {
	var buildings []*domain.Building

	// Build query with its arguments.
	query, args := buildGetAllQuery(filters)

	// Try to execute query.
	rows, err := repository.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to get all buildings with filter parameters %+v: %v", filters, err)
	}
	defer rows.Close()

	// Scan rows to the buildings slice.
	for rows.Next() {
		// Try to scan a building.
		var (
			building domain.Building
			info domain.BuildingInfo
		)
		err := rows.Scan(
			&building.Id,
			&info.Name,
			&info.City,
			&info.HandoverYear,
			&info.FloorsCount)
		if err != nil {
			return nil, fmt.Errorf("failed to scan a building: %v", err)
		}
		building.Info = &info

		// Append scanned building to the slice.
		buildings = append(buildings, &building)
	}

	// Check rows error after iterations completion.
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error after rows iteration: %v", err)
	}

	return buildings, nil
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
		ctx context.Context, info *domain.BuildingInfo) (*domain.Building, error) {
	// Execute insertion query in the database.
	row := repository.pool.QueryRow(
		ctx, insertQuery, info.Name, info.City, info.HandoverYear, info.FloorsCount)

	// Scan returned id of a new building in the database.
	var id int64
	if err := row.Scan(&id); err != nil {
		return nil, fmt.Errorf("failed to scan id of a new building: %v", err)
	}

	return domain.NewBuilding(id, info), nil
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

// Opens a new connection to building repository.
func OpenBuildingRepositoryImpl(
		ctx context.Context, uri string) (*BuildingRepositoryImpl, error) {
	// Try to open a new database connection pool.
	pool, err := pgxpool.New(ctx, uri)
	if err != nil {
		return nil, err
	}

	// Create a new database wrapper instance.
	impl := &BuildingRepositoryImpl{pool: pool}
	return impl, nil
}

// Builds query to get all buildings according to filter parameters.
func buildGetAllQuery(
		filters *logic.BuildingFilters) (query string, args []any) {
	query = getAllQueryPrefix
	var conditions []string

	// Add city filter if parameter is not nil.
	if filters.City != nil {
		condition := fmt.Sprintf("city = $%d", len(args) + 1)
		conditions = append(conditions, condition)
		args = append(args, *filters.City)
	}

	// Add handover year filter if parameter is not nil.
	if filters.HandoverYear != nil {
		condition := fmt.Sprintf("handover_year = $%d", len(args) + 1)
		conditions = append(conditions, condition)
		args = append(args, *filters.HandoverYear)
	}

	// Add floors count filter if parameter is not nil.
	if filters.FloorsCount != nil {
		condition := fmt.Sprintf("floors_count = $%d", len(args) + 1)
		conditions = append(conditions, condition)
		args = append(args, *filters.FloorsCount)
	}

	// Join conditions with query prefix.
	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	// Close a query and return it.
	query += ";"
	return query, args
}
