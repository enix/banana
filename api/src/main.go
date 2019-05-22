package main

import (
	"flag"

	"k8s.io/klog"

	"enix.io/banana/src/routes"
	"enix.io/banana/src/services"
)

// Assert : Ensure that the given error is a nil pointer
// 					otherwise print it and exit process with status code 1
func Assert(err error) {
	if err != nil {
		klog.Fatal(err)
	}
}

func main() {
	klog.InitFlags(nil)
	flag.Set("v", "1")
	flag.Parse()

	err := services.OpenVaultConnection()
	Assert(err)
	err = services.OpenDatabaseConnection()
	Assert(err)
	router, err := routes.InitializeRouter()
	Assert(err)
	router.Run(":80")
}
