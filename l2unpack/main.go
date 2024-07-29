package main

import (
	"l2unpack/unpack"
)

func main() {

	res := unpack.Unpack(`qwe\45`)
	print(res)

}

/*
Unpacking a string
Create a Go function that performs primitive unpacking of a string containing repeating characters / runes, for example:
* “a4bc2d5e” => “aaaabccddddddddde”
* “abcd” => “abcd”
* “45” => “” (invalid string)
Additional task: support escape - sequences
* `qwe\4\5` => `qwe45` (*)
* `qwe\45` => `qwe44444` (*)
* `qwe\\\5` => `qwe\\\\\` (*)*/
