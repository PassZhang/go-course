// 迭代
package main

import "fmt"

func demo6(number int) int {
	total := 0
	for {
		if number > 0 {
			total += number
			number--

		} else {
			break
		}
	}
	return total
}

func main() {
	fmt.Println(demo6(5))
}
