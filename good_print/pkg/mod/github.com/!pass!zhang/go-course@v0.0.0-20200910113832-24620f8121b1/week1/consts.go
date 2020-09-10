package main

import "fmt"

const c1 int = 1
const c2 = 3 - 1

const c3, c4 int = 3, 4

const c5, c6 = "5", c2 + c4

func main() {
	const (
		c7 int = 7
		c8
	)
	fmt.Println(c1, c2, c3, c4, c5, c6, c7, c8)
}
