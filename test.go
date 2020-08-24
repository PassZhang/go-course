package main

import "fmt"

var n1, n2, n3 int = 1, 8, -2
var n4 uint = 2

func main() {
	n1 += n2
	n2 -= n1
	n4 <<= n4

	fmt.Println()
}
