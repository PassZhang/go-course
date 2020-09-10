package main

import "fmt"

/*
if 布尔表达式 {
	在布尔表达式为 true 时执行
 }
*/

// func main() {
// 	/* 定义局部变量 */
// 	var a int = 10

// 	/* 使用 if 语句 判断布尔表达式 */
// 	if a < 20 {
// 		/* 如果条件为true 则执行以下语句 */
// 		fmt.Printf("a 小于 20\n")
// 	}
// 	fmt.Printf("a 的值为 %d\n", a)
// }

func main() {
	// 老婆
	// 买10个包子
	// 如果有卖西瓜的，买一个西瓜
	fmt.Println("买十个包子")

	var y string
	fmt.Print("有没有卖西瓜的:")
	fmt.Scan(&y)
	fmt.Println("你输入的是:", y)

	if y == "yes" {
		fmt.Println("买一个西瓜")
	}
}
