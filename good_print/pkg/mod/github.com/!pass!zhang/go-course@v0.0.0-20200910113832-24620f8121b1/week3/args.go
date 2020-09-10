package main

import "fmt"

// 可变参数的函数

func test(args ...string) {
	fmt.Printf("%T, %#v\n", args, args)
}

// 1. 可变参数 在一个方法中只能有一个
// 2. 并且可变参数必须放置函数声明的参数列表最后

// 举例（函数至少有n个参数）
// add （n1, n2, ...）
func add(n1, n2 int, args ...int) int {
	total := n1 + n2
	for _, v := range args {
		total = total + v
	}
	return total
}

func cacl(n1, n2 int, args ...int) int {
	// 直接调用add 把add方法结果进行返回
	// args 切片
	// swith(len(args))
	// add(n1, n2, args[], args[1], ...)

	return add(n1, n2, args...) // 解操作
}

func main() {
	test()
	test("1")
	test("1", "2", "3")

	fmt.Println(add(1, 2))
	fmt.Println(add(1, 2, 3))

	params := []int{1, 2, 30, 4}
	fmt.Println(add(1, 2, params...))
}
