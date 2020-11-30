package models

import (
	"dbsearcher/constants"
)

// CLIArgs stores all parsed cli args
type CLIArgs struct {
	Engine          constants.EngineType
	UserName        string
	Password        string
	Host            string
	Port            int
	DBname          string
	SSLmode         string
	SearchText      string
	Workers         int
	CheckExactMatch bool
}
