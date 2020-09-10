package main

import "fmt"

type callback func(val, total int) int

func demo7(val int, callback callback) int {
	if val > 0 {
		return callback(val, 0)
	} else {
		return val
	}
}

func addFunc(val int, total int) int {
	if val > 0 {
		return addFunc(val-1, total+val)

	} else {
		return total
	}
}

func main() {
	fmt.Println(demo7(5, addFunc))
}
