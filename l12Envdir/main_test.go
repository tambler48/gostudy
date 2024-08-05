package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"reflect"
	"testing"
)

func TestEnvDir(t *testing.T) {

	name := "/Users/anduser/_projects/gostudy/l12Envdir/main.go"
	args := []string{
		"/Users/anduser/_projects/gostudy/l12Envdir/envs",
		"subCommand/subCommand",
		"arg1", "arg2", "arg3",
	}
	cmd := exec.Command(name, args...)
	err := cmd.Run()
	fmt.Println(err)
}

func TestReadDir(t *testing.T) {
	path, err := os.MkdirTemp("", "envdir")
	defer func(path string) {
		err := os.RemoveAll(path)
		if err != nil {

		}
	}(path)
	expect := map[string]string{}
	fileInfo, content, err := createFile(path)
	if err != nil {
		t.Fatal(err)
	}
	expect[fileInfo.Name()] = string(content)

	fileInfo, content, err = createFile(path)
	if err != nil {
		t.Fatal(err)
	}
	expect[fileInfo.Name()] = string(content)

	result, err := readDir(path)
	if err != nil {
		t.Fatal(err)
	}
	reflect.DeepEqual(expect, result)

}

func createFile(path string) (os.FileInfo, []byte, error) {
	file1, err := os.CreateTemp(path, "CUSTOM")
	if err != nil {
		return nil, nil, err
	}
	fileString := getRandomSliceLetters(10)
	_, err = file1.Write(fileString)
	if err != nil {
		return nil, nil, err
	}
	fileInfo1, _ := file1.Stat()
	return fileInfo1, fileString, nil
}

func getRandomSliceLetters(length int) []byte {
	letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return b
}
