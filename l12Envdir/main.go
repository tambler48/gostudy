package main

import (
	"bytes"
	//"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	EnvDir()
}

func EnvDir() {
	arguments := os.Args[1:]
	if len(arguments) < 2 {
		fmt.Println("Usage: l12Envdir <path> <command> <arguments(optional)>")
		return
	}
	envdir := arguments[0]
	command := arguments[1]
	argumentsToSend := arguments[2:]

	envVariables, err := readDir(envdir)
	if err != nil {
		panic(err)
	}
	exitCode := runCmd(command, envVariables, argumentsToSend)
	os.Exit(exitCode)
}

func runCmd(name string, envs map[string]string, args []string) int {
	cmd := exec.Command(name, args...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	cmd.Stdin = os.Stdin
	envSlice := make([]string, 0, len(envs))
	for _, env := range envs {
		envSlice = append(envSlice, env)
	}
	cmd.Env = envSlice
	err := cmd.Run()

	if err != nil {
		fmt.Println(err)
		fmt.Println(stderr.String())
	}
	fmt.Println(stdout.String())
	return cmd.ProcessState.ExitCode()
}

func readDir(dir string) (map[string]string, error) {
	envVariables := make(map[string]string)
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		if !file.IsDir() { // add recursive read
			of, e := os.Open(filepath.Join(dir, file.Name()))
			if e != nil {
				return nil, e
			}
			content, e := io.ReadAll(of)
			if e != nil {
				return nil, e
			}
			envVariables[file.Name()] = string(content)
		}
	}
	return envVariables, nil
}

/*
envdir utility

Goal: Implement the envdir utility in Go.
This utility allows you to run programs by getting environment variables from a specific directory.
See man envdir Example go-envdir /path/to/evndir command arg1 arg2

Create a separate package (module) for this DZ in the repository

Implement a function like ReadDir(dir string) (map[string]string, error),
which scans the specified directory and returns all environment variables defined in it.

Implement a function like RunCmd(cmd []string, env map[string]string) int ,
which runs a program with arguments (cmd) with an overridden environment.

Implement a main function that parses command line arguments and calls ReadDir and RunCmd

Test the utility.
You can test the entire utility using a shell script, or you can write unit tests for individual functions.
Evaluation criteria: Standard input/output/error streams must be passed to the called program.
The exit code of the envdir utility must match the exit code of the program.
The code must pass go vet and golint checks
The teacher must be able to download and install the package using go get / go install
*/
