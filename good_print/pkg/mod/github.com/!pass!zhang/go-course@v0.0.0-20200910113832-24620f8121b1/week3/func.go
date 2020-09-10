package main

import "fmt"

// 导入包（标准包，自定义包，第三方包）

// 包级别的变量，常量，函数
// 无参 无返回值
func sayHello() {
	fmt.Println("hello, world")
}

// 有参 无返回值
func sayHi(name string, name2 string) {
	fmt.Println("Hi:", name, name2)
}

// 有参 有返回值
func add(n1 int, n2 int) int {
	return n1 + n2
}

// 有参 有返回值
func test(a int, b string) {
	fmt.Println(a, b)

}

func main() {
	//引用方法
	// main -> func -> ... -> sayHello

	// 调用 方法名
	sayHello()

	// 调用 方法名（参数）
	sayHi("小明", "小红")

	// // 有参数 有返回值
	// n := add(1, 2)
	// fmt.Println(n)

	var n int
	n = add(1, 2)
	fmt.Println(n)

	//调用参数值引用
	test(1, "100")
}
