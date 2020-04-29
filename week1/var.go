package main

import "fmt"

// 包级别 设置字符串变量
var packageVar string = "hello packageVar"

func main() {

	// 函数级别 设置字符串变量
	var funcVar string = "hello funcVar"

	{
		// 块级别变量
		var blockVar string = "hello blockVar"

		{
			// 子块变量
			var innerBlockVar string = "hello innerBlockVar "
			// 打印子块变量，块变量，函数变量，包变量
			fmt.Println(innerBlockVar, blockVar, funcVar, packageVar)
		}
		// 打印块变量，函数变量，包变量
		fmt.Println(blockVar, funcVar, packageVar)
	}
	// 打印函数变量
	fmt.Println(funcVar)
}
