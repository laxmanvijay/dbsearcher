package services

import (
	"database/sql"
	"dbsearcher/constants"
	"dbsearcher/models"
	"sync"
)

// Worker defines a independent function that works concurrently with other workers
func Worker(db *sql.DB, queries constants.ConnectionStructure, args models.CLIArgs, job <-chan string, recv chan<- string, wg *sync.WaitGroup) {
	for j := range job {
		SearchTable(db, queries, args, j, recv)
	}
	wg.Done()
}
