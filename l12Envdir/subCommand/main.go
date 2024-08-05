package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Environ())
	fmt.Println(os.Args[1:])
	//fmt.Println(io.ReadAll(os.Stdin))
	os.Exit(32)
}
