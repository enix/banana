package main

import (
	"fmt"
	"os/exec"
)

// execute : Execute binaries using our own stdio
func execute(cmd string, args ...string) ([]byte, error) {
	process := exec.Command(cmd, args...)
	output, err := process.CombinedOutput()
	fmt.Println(string(output))
	return output, err
}
