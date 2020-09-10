/*
作用域
1. 作用域之比那里可以使用范围。
2. go 语言使用大括号显示的标识作用域范围，大括号内包含一连串的语句叫做语句块。
3. 语句块可以嵌套，语句块内定义的变量不能再语句块外使用。

常见隐式语句块：
1. 全语句块
2. 包语句块
3. 文件语句块
4. if、switch、for、select、case语句块

作用域内定义变量只能被声明一次且变量必须使用，否则编译错误。在不同作用域定义相同的变量，此时局部将覆盖全局。
*/

package main 

import "fmt"

func main() {
	outer := 1
	{
		inner := 2
		fmt.Println(outer, inner)
		outer := 3
		fmt.Println(outer, inner)
	}
	fmt.Println(outer)
}

