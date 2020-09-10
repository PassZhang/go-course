package main

import (
	"fmt"

	"github.com/PassZhang/go-course/go_mod_test/src/gpkgname/pkg01" // 导入包pkg01， 路径gpkgname/pkg01
	"github.com/PassZhang/go-course/go_mod_test/src/gpkgname/pkg02" // 导入包pkg02， 路径gpkgname/pkg02
)

func main() {
	fmt.Println("go_mod")
	fmt.Println(pkg01.Name) // 调用pkg01包中的成员Name
	fmt.Println(pkg02.Name) // 调用pkg02包中的成员Name
}
