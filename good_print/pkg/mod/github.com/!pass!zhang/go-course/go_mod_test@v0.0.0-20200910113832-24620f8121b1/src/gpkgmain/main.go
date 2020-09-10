package main

import (
	"fmt"

	"github.com/imsilence/chapter08/modtest/pkgname/pkg01" // 导入包pkg01， 路径gpkgname/pkg01
	"github.com/imsilence/chapter08/modtest/pkgname/pkg02" // 导入包pkg02， 路径gpkgname/pkg02
)

func main() {
	fmt.Println("gpkgmain")
	fmt.Println(pkg01.Name) // 调用pkg01包中的成员Name
	fmt.Println(pkg02.Name) // 调用pkg02包中的成员Name
}
