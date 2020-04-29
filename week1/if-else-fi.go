package main

import "fmt"

/*
if 布尔表达式 1 {
   // 在布尔表达式 1 为 true 时执行
   if 布尔表达式 2 {
    // 在布尔表达式 2 为 true 时执行
 }
}
*/

func main() {
	// 定义局部变量
	var a int = 1001
	var b int = 200

	// 判断条件
	if a == 100 {
		if b == 200 {
			fmt.Printf("a 的值为100，b 的值为200\n")
		}
	}
	fmt.Printf("a 值为：%d\n", a)
	fmt.Printf("b 值为：%d\n", b)
}
