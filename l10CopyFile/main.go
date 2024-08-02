package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"time"
)

var from string
var to string
var offset int
var limit int

func init() {
	flag.StringVar(&from, "from", "l10CopyFile/from.txt", "File source of copying")
	flag.StringVar(&to, "to", "l10CopyFile/to.txt", "Target file to copying")
	flag.IntVar(&offset, "offset", 0, "Offset in source file")
	flag.IntVar(&limit, "limit", 0, "How many bytes to copy")
}

func main() {
	flag.Parse()
	start := time.Now()

	err := CopyManager(from, to, offset, limit)

	end := time.Now()
	if err == nil {
		duration := end.Sub(start)
		fmt.Printf("Copied in %s, %v bytes\n", duration, limit)
	} else {
		fmt.Println(err)
	}
}

func CopyManager(from string, to string, offset int, limit int) error {
	fromFile, fromErr := os.OpenFile(from, os.O_RDONLY, 0)
	defer func(fromFile *os.File) {
		err := fromFile.Close()
		if err != nil {

		}
	}(fromFile)

	err := validate(fromFile, fromErr, offset)

	if err != nil {
		return err
	}
	toFile, err := os.OpenFile(to, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	defer func(toFile *os.File) {
		err := toFile.Close()
		if err != nil {

		}
	}(toFile)
	err = startCoping(fromFile, toFile, offset, limit)

	if err != nil {
		return err
	}
	return nil
}

func prepareLimit(fromFile *os.File, limit int, offset int) int {
	if limit == 0 {
		fromFileInfo, _ := fromFile.Stat()
		limit = int(fromFileInfo.Size()) - offset
	}
	return limit
}
func getBufferSize(limit int) int {
	defBuffSize := 1000
	if limit < defBuffSize {
		defBuffSize = limit
	}
	return defBuffSize
}

func startCoping(fromFile *os.File, toFile *os.File, offset int, limit int) error {
	limit = prepareLimit(fromFile, limit, offset)
	buffSize := getBufferSize(limit)
	buffer := make([]byte, buffSize)

	_, err := fromFile.Seek(int64(offset), io.SeekStart)
	if err != nil {
		return err
	}
	copied := 0
	for copied < limit {
		read, err := fromFile.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		_, err2 := toFile.Write(buffer)

		if err2 != nil {
			return err
		}

		copied += read
		progress := copied * 100 / limit
		fmt.Printf("\rCopied %v%%", progress)
		if (limit - copied) < buffSize {
			buffer = make([]byte, limit-copied)
		}
	}

	fmt.Print("\n")

	return nil
}

func validate(fromFile *os.File, err error, offset int) error {
	if err != nil && os.IsNotExist(err) {
		return errors.New("ERROR: file not exists")
	}
	fromFileInfo, err := fromFile.Stat()
	size := fromFileInfo.Size()

	if !fromFileInfo.Mode().IsRegular() {
		return errors.New("ERROR: source file is not a regular file")
	}
	if size == 0 {
		return errors.New("ERROR: source file is empty")
	}
	if int(size) < offset {
		return errors.New("ERROR: offset too big")
	}
	return nil
}

/*
go run main.go -from=from.txt -to=to.txt -offset=5 -limit=10

Copying files
Goal: Implement a file copying utility
The utility should accept the following arguments
* source file (From)
* copy file (To)
* Source offset (Offset), default - 0
* Number of bytes to copy (Limit), default - the entire file from From
Output copy progress to console in %, for example using github.com/cheggaaa/pb
The program may NOT process files whose length is unknown (for example /dev/urandom).
Create a separate package (module) for this homework in the repository
Implement a function like Copy(from string, to string, limit int, offset int) error
Write unit tests for the Copy function
Implement the main function that analyzes the command line parameters and calls Copy
Check the installation and operation of the utility manually
Evaluation criteria: The function must pass all tests
All files required for the tests must be created in the test itself
The code must pass go vet and golint checks
The teacher must be able to download, check and install the package using go get / go test / go install*/
