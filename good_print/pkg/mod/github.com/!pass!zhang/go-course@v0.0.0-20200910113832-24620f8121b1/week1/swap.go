// 可以通过赋值运算=更新变量的值，Go 语言支持通过元组赋值同时更新多个变量的值，例如：n1, n2 = 1, 2，可用于两个变量的交换 x, y = y,x

package main

import "fmt"

func main() {
	n1, n2 := 1, 2

	fmt.Println(n1, n2)

	n1, n2 = n2, n1
	fmt.Println(n1, n2)
}
