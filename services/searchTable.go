package services

import (
	"database/sql"
	"dbsearcher/constants"
	"dbsearcher/models"
	"fmt"
	"strings"
)

//SearchTable performs search
func SearchTable(db *sql.DB, queries constants.ConnectionStructure, args models.CLIArgs, tableName string, recv chan<- string) {
	fmt.Printf("scanning %s\n", tableName)
	if strings.HasPrefix(tableName, queries.InternalTablePrefix) {
		return
	}
	query := fmt.Sprintf(queries.DescribeTable, tableName)

	rows, err := db.Query(query)
	HandleErr(err, true)
	if err != nil {
		return
	}
	defer rows.Close()
	var props models.TableDescription
	for rows.Next() {
		var entry models.TableDescriptionEntry
		err := rows.Scan(&entry.Name, &entry.DataType)
		HandleErr(err, true)
		props.Props = append(props.Props, entry)
	}

	selectAll := fmt.Sprintf(queries.SelectAll, tableName)

	rawResult := make([][]byte, len(props.Props))
	result := make([]string, len(props.Props))

	temp := make([]interface{}, len(props.Props))
	for i := range rawResult {
		temp[i] = &rawResult[i]
	}

	rows, err = db.Query(selectAll)
	HandleErr(err, true)
	if err != nil {
		return
	}
	defer rows.Close()
	i := 1
	for rows.Next() {
		err = rows.Scan(temp...)
		HandleErr(err, true)
		if err != nil {
			continue
		}
		for i, raw := range rawResult {
			if raw == nil {
				result[i] = constants.InvalidData
			} else {
				result[i] = string(raw)
			}
		}
		j := 0
		for _, res := range result {
			if !args.CheckExactMatch {
				if strings.Contains(res, args.SearchText) {
					recv <- fmt.Sprintf(constants.ResponseStructure, tableName, props.Props[j].Name, i)
				}
			} else {
				if res == args.SearchText {
					recv <- fmt.Sprintf(constants.ResponseStructure, tableName, props.Props[j].Name, i)
				}
			}
			j++
		}
		i++
	}
}
