// 常量
// go 支持，字符串，布尔值和数值的常量
package main 

import "fmt"

// 用const关键字定义常量
const (

	num1 = "a"
	num2 = 2
)

func main()  {
	fmt.Println(num1)
	fmt.Println(num2)
	
	const num3 = 10 - num2 

	// 表达式里面可以有常量， 但不能有变量
	/*
		num4 := 3 
		const num5 = 10 - num4 // 错误
	*/

	// 常量定义后就不能再修改了（只读）
	/*
		num1 = 3 //错误

	*/
}