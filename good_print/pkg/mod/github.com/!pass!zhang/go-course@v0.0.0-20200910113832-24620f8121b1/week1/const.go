package main

import (
	"fmt"
)

// const 常量
const (
	packageName string = "string"
	packageMsg         = "package"
)

func main() {
	const name string = "kk"
	const msg = "msg" // 常量可以不使用
	fmt.Println(name)

	// name = "slience" 常量一旦定义不能修改
}
