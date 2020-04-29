package main

import "fmt"

// 定义变量的四种方式
func main() {
	// 定义了类型并且初始化了值
	var name string = "kk"

	// 定义变量类型，但不初始化值
	var zeroString string
	// 初始化使用类型对应的零值(空字符串"")

	// 定义变量省略类型，不能省略初始化值
	var typeString = "google"
	// 通过对应的值类型推到变量的类型

	// 短声明(必须在函数内使用，不能在包级别使用)
	shortString := "baidu"
	// 通过对应的值类型推到变量的类型

	fmt.Println(name, zeroString, typeString, shortString)

}
