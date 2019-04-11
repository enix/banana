package main

import (
	"os"
)

// CheckEnvVariablePresence : Verify if the given key is present in process env
func CheckEnvVariablePresence(key string) bool {
	return len(os.Getenv(key)) > 0
}
