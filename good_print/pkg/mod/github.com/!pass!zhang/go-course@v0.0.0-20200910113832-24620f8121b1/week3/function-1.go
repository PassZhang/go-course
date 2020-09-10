// 函数（不支持嵌套、重载和默认参数，可以将函数作为一个值进行赋值）
package main

import "fmt"

func add(number1 int, number2 int) (number3 int) {
	number3 = number1 + number2
	return
}

func subtraction(number1, number2 int) int {
	number3 := number1 - number2
	return number3
}

func multiplication(number ...int) int {
	total := 1

	for _, v := range number {
		total *= v

	}
	return total

}

func demo1() {
	fmt.Println(add(8, 7))
	fmt.Println(subtraction(8, 7))
	fmt.Println(multiplication(2, 3, 5, 10))
}

func main() {
	demo1()
}
