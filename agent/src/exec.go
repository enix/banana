package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"k8s.io/klog"
)

func executeWithExtraFD(cmd string, args ...string) ([]byte, []byte, *os.File, error) {
	klog.Info("spawning plugin with arguments: ", args)
	process := exec.Command(cmd, args...)

	stdoutPipe, err := process.StdoutPipe()
	if err != nil {
		return nil, nil, nil, err
	}
	stderrPipe, err := process.StderrPipe()
	if err != nil {
		return nil, nil, nil, err
	}

	extraPipeR, extraPipeW, err := os.Pipe()
	if err != nil {
		return nil, nil, nil, err
	}
	process.ExtraFiles = []*os.File{extraPipeW}

	err = process.Start()
	if err != nil {
		return nil, nil, nil, err
	}

	stdout, err := ioutil.ReadAll(stdoutPipe)
	if err != nil {
		return nil, nil, nil, err
	}
	stderr, err := ioutil.ReadAll(stderrPipe)
	if err != nil {
		return nil, nil, nil, err
	}

	err = process.Wait()
	extraPipeW.Close()

	fmt.Println(string(stderr))
	return stdout, stderr, extraPipeR, err
}

// execute : Execute binaries
func execute(cmd string, args ...string) ([]byte, []byte, error) {
	stdout, stderr, _, err := executeWithExtraFD(cmd, args...)
	return stdout, stderr, err
}
