package main

import "fmt"

func main() {
	// const B 不允许使用
	const (
		A = "TEST"
		B // 使用前一个常量的值进行初始化（B）
		C // 使用前一个常量的值进行初始化（c=>b）
		D = "TESTD"
		E //使用前一个常量的执行进行初始化（D）
		F //使用前一个常量的值进行初始化（E=>D）
	)
	fmt.Println(B, C)
	fmt.Println(E, F)
}
