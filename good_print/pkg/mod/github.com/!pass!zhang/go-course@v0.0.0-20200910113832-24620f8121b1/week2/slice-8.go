package main

import "fmt"

func main() {
	// `copy` 函数 从源slice的srv中负载元素到目标dst， 并返回复制的元素的个数

	number1 := []int{1, 2, 3, 4, 5}
	number2 := []int{6, 7, 8}
	number3 := copy(number1, number2)
	fmt.Println(number1, number2, number3)

	number4 := []int{1, 2, 3, 4, 5}
	number5 := []int{6, 7, 8}
	number6 := copy(number5, number4)
	fmt.Println(number4, number5, number6)

}
