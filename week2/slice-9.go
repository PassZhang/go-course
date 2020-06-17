package main

import "fmt"

func main() {
	// append 向slice里面追加一个或者多个元素，然后返回一个和slice一样类型的slice

	number1 := []int{1, 2, 3}
	fmt.Println(number1, len(number1), cap(number1))

	number1 = append(number1, 4, 5)
	fmt.Println(number1, len(number1), cap(number1)) // 按初始容量成倍扩大

	number3 := append(number1, number1...) //用省略号自动振凯，以使用每个元素
	fmt.Println(number3)
}