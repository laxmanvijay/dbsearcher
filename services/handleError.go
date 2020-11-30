package services

import (
	"dbsearcher/utils"
)

var logger utils.CustomLoggerMethods

func init() {
	logger = utils.GetLogger()
}

// HandleErr is a convient function to handle errors
func HandleErr(err error, soft bool) {
	if err != nil {
		if !soft {
			panic(err)
		} else {
			logger.Error(err)
		}
	}
}
