package main

import (
	"io"
	"math/rand"
	"os"
	"testing"
)

func TestCopyManager(t *testing.T) {
	sourceContent := getRandomSliceLetters(400)
	cases := []struct {
		input struct {
			output []byte // source file content
			offset int
			length int
		}
		want []byte // target file content
	}{
		{
			input: struct {
				output []byte
				offset int
				length int
			}{
				output: sourceContent,
				offset: 10,
				length: 20,
			},
			want: sourceContent[10:30],
		},
		{
			input: struct {
				output []byte
				offset int
				length int
			}{
				output: sourceContent,
				offset: 10,
				length: 0,
			},
			want: sourceContent[10:],
		},
		{
			input: struct {
				output []byte
				offset int
				length int
			}{
				output: sourceContent,
				offset: 0,
				length: 0,
			},
			want: sourceContent,
		},
	}

	for _, testCase := range cases {

		outputFile, err := os.CreateTemp("", "output")
		if err != nil {
			t.Fatal(err)
		}
		_, err = outputFile.Write(testCase.input.output)
		if err != nil {
			t.Fatal(err)
		}
		inputFile, err := os.CreateTemp("", "input")
		if err != nil {
			t.Fatal(err)
		}

		err = CopyManager(outputFile.Name(), inputFile.Name(), testCase.input.offset, testCase.input.length)
		if err != nil {
			t.Fatal(err)
		}
		content, err := io.ReadAll(inputFile)
		if err != nil {
			t.Fatal(err)
		}

		if string(content) != string(testCase.want) {
			t.Errorf("Content is not equal")
		}
		closeAndDelete(outputFile)
		closeAndDelete(inputFile)
	}
}

func getRandomSliceLetters(length int) []byte {
	letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return b
}

func closeAndDelete(file *os.File) {
	defer func(outputFile *os.File) {
		err := outputFile.Close()
		if err != nil {

		}
	}(file)
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {

		}
	}(file.Name())
}
