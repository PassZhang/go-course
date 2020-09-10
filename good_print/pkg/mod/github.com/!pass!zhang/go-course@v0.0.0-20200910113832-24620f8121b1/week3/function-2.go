// 匿名函数：通常不希望再次使用（即只使用一次）的函数可以定义为匿名函数

package main

import (
	"fmt"
)

func demo2() {
	// f1为函数地址
	f1 := func(x, y int) int {
		z := x + y
		return z
	}
	fmt.Println(f1)
	fmt.Println(f1(5, 6))

	//直接创建匿名函数并运行
	f2 := func(x, y int) int {
		z := x + y
		return z
	}(5, 6)
	fmt.Println(f2)

	// 直接创建匿名函数并运行（无参数的形式）()调用匿名函数

	func() {
		fmt.Println(5 + 6)
	}()
}

func main() {
	demo2()
}
