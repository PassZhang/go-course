package main

import "fmt"

func main() {
	// - >
	number1 := make([]int, 5, 10)

	// fmt.Println(number1[5])  // 超出下标索引

	number2 := number1[2:7]
	fmt.Println(number2, cap(number2))

}
