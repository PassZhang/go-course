package main

import "fmt"

func main() {
	number1 := make([]int, 4, 11)

	number2 := number1[2:7]

	number2[3] = 3
	fmt.Println(number1)
	fmt.Println(number2)

	number2[3] = 4
	fmt.Println(number1)
	fmt.Println(number2)

}
