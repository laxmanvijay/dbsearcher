package main

import (
	"database/sql"
	"dbsearcher/constants"
	"dbsearcher/models"
	"dbsearcher/services"
	"dbsearcher/utils"
	"fmt"
	"sync"
)

var (
	tables  []models.TableListModel
	db      *sql.DB
	err     error
	args    models.CLIArgs
	wg      sync.WaitGroup
	job     chan string
	recv    chan string
	queries constants.ConnectionStructure
)

func main() {

	args = utils.ParseCLIArgs()

	queries, err = constants.GetQueries(args.Engine)
	services.HandleErr(err, false)

	db, err = services.GetConnection(args, queries)
	services.HandleErr(err, false)

	readTables()

	initializeWorkers()

	checkResults()
}

func readTables() {
	fmt.Println("reading all tables...")
	tables = services.ListAllTables(db, queries)
	fmt.Println("done...")
}

func initializeWorkers() {

	job = make(chan string, 200)
	recv = make(chan string, 200)

	go func() {
		for i := 0; i < args.Workers; i++ {
			wg.Add(1)
			go services.Worker(db, queries, args, job, recv, &wg)
		}
		wg.Wait()
		close(recv)
	}()

	go func() {
		for _, name := range tables {
			job <- name.TableName
		}
		close(job)
	}()
}

func checkResults() {
	foundMatches := false
	for res := range recv {
		if !foundMatches {
			fmt.Println("Found matches...")
			foundMatches = true
		}
		fmt.Println(res)
	}
	if !foundMatches {
		fmt.Println("No matches found")
	}
}
