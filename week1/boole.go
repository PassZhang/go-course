// 基础数据类型 --> 布尔类型(boole)

package main

import "fmt"

func main() {
	var (
		isBody bool = true
		isGril bool = false
	)
	// 逻辑运算 && 与
	fmt.Println(isBody && isBody) // 1 1 得 1
	fmt.Println(isBody && isGril) // 1 0 得 0
	fmt.Println(isGril && isBody) // 0 1 得 0
	fmt.Println(isGril && isGril) // 0 0 得 0

	// 逻辑运算 || 或
	fmt.Println(isBody || isBody) // 1 1 得 1
	fmt.Println(isBody || isGril) // 1 0 得 1
	fmt.Println(isGril || isBody) // 0 1 得 1
	fmt.Println(isGril || isGril) // 0 0 得 0

	// 逻辑运算 ！ 非
	fmt.Println(!isBody)
	fmt.Println(!isBody)
	fmt.Println(!isGril)
	fmt.Println(!isGril)

	// 关系运算  == 等于
	fmt.Println(isBody == isGril)
	fmt.Println(isBody != isGril)

	// 打印
	fmt.Println(isBody, isGril)
	fmt.Printf("isBody=%#v, isGril=%#v\n", isBody, isGril)
}
