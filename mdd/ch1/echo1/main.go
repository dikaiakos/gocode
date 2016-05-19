package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string                   // implicitly initialised to zero value
	for i := 1; i < len(os.Args); i++ { // short variable declaration ":=" used to
		// declare i and give it appropriate type based on initial value
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
