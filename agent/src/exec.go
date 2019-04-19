package main

import (
	"fmt"
	"os/exec"
)

// Execute : Execute binaries using our own stdio
func Execute(cmd string, args ...string) ([]byte, error) {
	process := exec.Command(cmd, args...)
	output, err := process.CombinedOutput()
	fmt.Println(string(output))
	return output, err
}
