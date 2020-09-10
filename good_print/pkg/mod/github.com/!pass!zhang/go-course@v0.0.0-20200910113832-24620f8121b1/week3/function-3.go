// 闭包函数（函数内在包含子函数，并最终return子函数）

package main 

import "fmt"

func number1(x int) func(y int)  int {
	return func (y int) int {
		return x + y
	}
}

func number2() func() {
	return func() {
		fmt.Println("Hello World")
	}
}

func demo3()  {
	n1 := number1(7)
	fmt.Println(n1(8))

	n2 := number2()
	n2()
}

func main()  {
	demo3()
}