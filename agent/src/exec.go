package main

import (
	"os"
	"os/exec"
)

// Execute : Execute binaries using our own stdio
func Execute(cmd string, args ...string) error {
	process := exec.Command(cmd, args...)
	process.Stdout = os.Stdout
	process.Stderr = os.Stderr
	return process.Run()
}
