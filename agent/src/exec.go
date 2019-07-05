package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"

	"k8s.io/klog"
)

// execute : Execute binaries using our own stdio
func execute(cmd string, args ...string) ([]byte, []byte, error) {
	klog.Info("spawning plugin with arguments: ", args)
	process := exec.Command(cmd, args...)

	stdoutPipe, err := process.StdoutPipe()
	if err != nil {
		return nil, nil, err
	}
	stderrPipe, err := process.StderrPipe()
	if err != nil {
		return nil, nil, err
	}

	err = process.Start()
	if err != nil {
		return nil, nil, err
	}

	stdout, err := ioutil.ReadAll(stdoutPipe)
	if err != nil {
		return nil, nil, err
	}
	stderr, err := ioutil.ReadAll(stderrPipe)
	if err != nil {
		return nil, nil, err
	}

	fmt.Print(string(stderr))
	err = process.Wait()
	return stdout, stderr, err
}
