package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] { // _ is the blank identifier; range returns pair of values: index, elt
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
