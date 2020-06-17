// 数组array
package main

import "fmt"

// 如果再数组的长度位置出现的是"..."省略号，则表示数组的长度是根据初始化值的个数来计算
// 数组的长度必须是常量表达式，因为数组的长度需要在编译阶段确定
// func main() {
// number1 := [3]int{1, 2, 2:3}

// number2 := [3][4]int{
// 	{1, 2, 3, 4},
// 	{5, 6, 7, 8},
// 	{0:9, 1:10, 2:11, 3:12},
// }

// number3 := [2][2][3]int {
// 	{{1, 2, 3}, {4, 5, 6}},
// 	{{7, 8, 9}, {10, 11, 12}},
// }
// fmt.Println(number1)
// fmt.Println(number2)
// fmt.Println(number3)

// // 冒泡排序
// number4 := [...]int{3, 1, 9, 6, 8}

// 	for i := 0; i < len(number4); i++ {
// 		for j := i + 1; j < len(number4); j++ {
// 			if number4[i] < number4[j] {
// 				number4[i], number4[j] = number4[j], number4[i]
// 			}
// 		}
// 	}

// 	fmt.Println(number4)
// }

// 定义数组时，指定数组元素个数必须是常量
func main() {
	/*
		// 定义数组时，指定元素个数必须是常量
		// non-constant array bound a
	*/
	// 报错：non-constant array bound a
	a := 10
	var b [a]int

	fmt.Println(b)
}
