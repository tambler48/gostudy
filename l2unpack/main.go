package main

import (
	"l2unpack/unpack"
)

func main() {

	res := unpack.Unpack(`qwe\45`)
	print(res)

}
