package main

import "fmt"

/*
if 布尔表达式 {
	在布尔表达式为 true 时执行
 }
*/
func main() {
	var s int // 声明变量 s 是需要判断的数
	fmt.Printf("请输入一个数字：")
	fmt.Scan(&s)

	if s%2 == 0 { // 取 s 除以 2 的余数是否等于 0
		fmt.Printf("s 是偶数\n") // 如果成立
	} else {
		fmt.Printf("s 不是偶数\n") // 否则
	}
	fmt.Printf("s 的值是：", s)
}
