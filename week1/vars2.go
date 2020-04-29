package main

import "fmt"

func main() {

	var name string = "kk"

	var zeroString string = "zero" //定义变量类型，但不初始化值
	// 初始化使用类型对应的零值（空字符串""）
	fmt.Printf("%T %#v\n", zeroString, zeroString)
	fmt.Println(name)

	var typeString = "kk" //定义变量可以省略类型，不能省略初始化值

	shortString := "kk" //短声明（必须在函内包含函数内子块使用，不能在包级别使用）
	// 通过对应的值类型推到变量的类型

	fmt.Println(name, zeroString, typeString, shortString)
}
