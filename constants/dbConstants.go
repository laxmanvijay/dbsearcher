package constants

import (
	"errors"
)

const (
	postgresConnectionString = "host=%s port=%d user=%s password=%s dbname=%s sslmode=%s"

	postgresListAllTablesQuery = "SELECT tablename FROM pg_catalog.pg_tables"

	postgresDescribeTable = "SELECT column_name, data_type FROM information_schema.COLUMNS WHERE TABLE_NAME = '%s'"

	postgresInternalTablePrefix = "pg_"

	selectAll = "SELECT * FROM %s"
)

// ConnectionStructure represents the required queries
type ConnectionStructure struct {
	ConnectionString    string
	ListTables          string
	DescribeTable       string
	SelectAll           string
	InternalTablePrefix string
}

// GetConstMap gives a postgres config map
func GetConstMap() map[EngineType]ConnectionStructure {
	return map[EngineType]ConnectionStructure{
		PostgreSQL: ConnectionStructure{
			ConnectionString:    postgresConnectionString,
			ListTables:          postgresListAllTablesQuery,
			DescribeTable:       postgresDescribeTable,
			SelectAll:           selectAll,
			InternalTablePrefix: postgresInternalTablePrefix,
		},
	}
}

// GetQueries gives a datamodel
func GetQueries(engine EngineType) (ConnectionStructure, error) {
	config := GetConstMap()

	if _, ok := config[engine]; ok {
		return config[engine], nil
	}
	return ConnectionStructure{}, errors.New("Invalid engine type")
}
