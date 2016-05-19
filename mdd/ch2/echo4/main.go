package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	fmt.Printf("%d %d\n", x, y)
	return x
}

//func (c Celsius) String() string
func main() {

	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}

	fmt.Printf("gcd(10, 4):%d\n", gcd(10, 4))
}
