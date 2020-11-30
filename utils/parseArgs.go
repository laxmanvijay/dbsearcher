package utils

import (
	"dbsearcher/models"
	"flag"
	"fmt"
	"os"
)

// ParseCLIArgs parses the arguments from the CLI
func ParseCLIArgs() models.CLIArgs {
	var (
		args   models.CLIArgs
		logger CustomLoggerMethods
	)

	flag.StringVar(&args.UserName, "username", "postgres", "username to connect to database")
	flag.StringVar(&args.Password, "password", "postgres", "password to connect to database")
	flag.StringVar(&args.Host, "host", "localhost", "host to connect to database")
	flag.IntVar(&args.Port, "port", 5433, "username to connect to database")
	flag.StringVar(&args.DBname, "db", "postgres", "name of database")
	flag.StringVar(&args.SSLmode, "ssl", "disable", "does it require ssl?")
	flag.StringVar(&args.SearchText, "search", "", "text to search")
	flag.Var(&args.Engine, "engine", "database engine")
	flag.BoolVar(&args.CheckExactMatch, "checkExactMatch", false, "should the searchText match exactly")
	flag.IntVar(&args.Workers, "workers", 5, "Number of concurrent workers")

	flag.Parse()

	if args.SearchText == "" {
		var noSearchText = "provide a search text using --search flag"
		fmt.Println(noSearchText)
		logger.Error(noSearchText)
		os.Exit(1)
	}

	return args
}
