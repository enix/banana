package main

import (
	"os"

	"enix.io/banana/src/logger"
	"enix.io/banana/src/routes"
	"enix.io/banana/src/services"
)

// Assert : Ensure that the given error is a nil pointer
// 					otherwise print it and exit process with status code 1
func Assert(err error) {
	if err != nil {
		logger.LogError(err)
		os.Exit(1)
	}
}

func main() {
	err := services.OpenVaultConnection()
	Assert(err)
	err = services.OpenDatabaseConnection()
	Assert(err)
	err = services.OpenStorageConnection()
	Assert(err)
	router, err := routes.InitializeRouter()
	Assert(err)
	router.Run(":80")
}
