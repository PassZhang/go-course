// 变量（补充：类型推断的好处）

/*
var name string  // 可以在函数外或者函数内声明一个变量
name = "winnie" // 只能在函数内使用
---
name := "winnie" // 只能在函数内声明
*/

package main 

import "fmt"

// func Calculation() (num int) {
// 	num = 5
// 	return
// }

// func main()  {
// 	var number int
// 	number = Calculation()
// 	fmt.Println(number)
// }
/*
1. 修改某个程序与外界的交互规则，只改变内部的实现而不影响外部
2. 修改Calculation的返回类型时，也要修改number的类型
3. 修改了接口返回值类型，接口外部也要跟着改（number := Calculation()）则不要
4. Go语言的类型推断可以明显提升程序的灵活性，使得代码重构变得更加容易（写起来像动态语言）
*/


func Calculation() (num string) {
	num = "5"
	return 
}

func main()  {
	var number string
	number = Calculation()
	fmt.Println(number)
}
