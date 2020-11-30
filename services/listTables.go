package services

import (
	"database/sql"
	"dbsearcher/constants"
	"dbsearcher/models"
)

// ListAllTables get names of all tables in the db
func ListAllTables(db *sql.DB, queries constants.ConnectionStructure) []models.TableListModel {

	var tables = make([]models.TableListModel, 1)

	rows, err := db.Query(queries.ListTables)
	HandleErr(err, false)
	defer rows.Close()

	for rows.Next() {
		var tableList models.TableListModel
		err := rows.Scan(&tableList.TableName)
		HandleErr(err, false)
		tables = append(tables, tableList)
	}

	return tables
}
