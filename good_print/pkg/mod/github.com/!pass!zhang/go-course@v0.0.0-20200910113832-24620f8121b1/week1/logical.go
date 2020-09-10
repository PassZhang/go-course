package main

import "fmt"

func main() {
	var a bool = true
	var b bool = false
	if a && b {
		fmt.Printf("第一行条件为true\n")
	} else {
		fmt.Printf("第一行条件为flase\n")
	}

	if a || b {
		fmt.Printf("第二行条件为true\n")
	}

	/* 修改a和b的值 */
	a = false
	b = true
	if a && b {
		fmt.Printf("第三行条件为true\n")
	} else {
		fmt.Printf("第三行条件为flase\n")
	}

	if !(a || b) {
		fmt.Printf("第四行条件为true\n")
	}
}
