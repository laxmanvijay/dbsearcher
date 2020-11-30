package services

import (
	"database/sql"
	"dbsearcher/constants"
	"dbsearcher/models"
	"fmt"

	_ "github.com/lib/pq" // used for db connection
)

type postgres struct {
	connectionString string
	showTables       string
}

// GetConnection provides db connection
func GetConnection(args models.CLIArgs, queries constants.ConnectionStructure) (*sql.DB, error) {
	var connectionString = fmt.Sprintf(queries.ConnectionString, args.Host, args.Port, args.UserName, args.Password, args.DBname, args.SSLmode)

	db, err := sql.Open("postgres", connectionString)
	return db, err
}
