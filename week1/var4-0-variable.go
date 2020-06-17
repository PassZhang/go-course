// 如果变量没有显式初始化，则被隐式的赋予其类型的零值（zero value），数字类型是0，字符类型是空字符串""
// slice, map, channel, pointer, func, interface 零值都为nil

package main

import "fmt"

func main() {
	var a int = 1
	fmt.Println(a)

	// var 可以声明一个或者多个变量
	// 已声明未使用的变量会在编译阶段报错
	var b, c string = "a", "b"
	fmt.Println(b, c)

	// 声明为没有相应初始化的变量是零值的（int的零值是0， string的零值是空， slice的零值是nil）
	var d int
	fmt.Println(d)

	// go 将推断初始化变量的类型
	// (:=)简短变量声明一般用于局部变量的声明和初始化， var往往是用在需要先指定变量类型稍后再赋值的。
	// 不过它有一个限制，那就是它只能用在函数内部；在函数外部使用则会无法编译通过，所以一般用var方式来定义全局变量
	e := true
	fmt.Println(e)

	// _ （下划线）是个特殊的变量名，任何赋予它的值都会被丢弃
	_, f := 7, 8
	fmt.Println(f)

}
