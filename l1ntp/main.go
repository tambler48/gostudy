package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
)

func main() {
	time, err := ntp.Time("pool.ntp.org")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error is - %v", err)
		os.Exit(1)
	}
	fmt.Println(time)
}
