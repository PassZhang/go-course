// 数组的元素也可以是数组类型，此时称为多维数组
package main

import "fmt"

func main() {
	// 声明&初始化
	var students00 [2][3]string
	students01 := [...][2]string{{"kk", "woniu"}, {"ada", "wd"}, {"xuegao", "xiaolin"}} // 多维数组只有第一维可以使用变量数量推测
	students02 := [3][3]string{{"kk", "woniu", "ada"}, {"wd", "xiaoming", "xiaohong"}, {"xiaoli", "xuegao", "xiaolin"}}
	students03 := [3][3]string{0: {"kk", "woniu", "ada"}, 1: {"xiao", "li"}, 2: {"wd"}}
	students04 := [3][3]string{0: {0: "kk", "woniu", "ada"}, 1: {"xiaozhang", "xiaoli", "xiaohong"}, 2: {0: "xiaohong", 1: "wd", 2: "xuegao"}}

	fmt.Printf("%T, %T, %T, %T, %T\n", students00, students01, students02, students03, students04)
	fmt.Printf("%q\n", students00)
	fmt.Printf("%q\n", students01)
	fmt.Printf("%q\n", students02)
	fmt.Printf("%q\n", students03)
	fmt.Printf("%q\n", students04)

	// 访问&修改
	fmt.Printf("%q\n", students01[1])
	fmt.Println(students01[0][0])

	// 修改元素
	students01[1] = [2]string{"无敌", "卧底"}
	students01[0][0] = "KK"
	fmt.Printf("%q\n", students01)

	fmt.Printf("----------------------\n")
	// 遍历元素
	for i := 0; i < len(students02); i++ {
		for j := 0; j < len(students02[i]); j++ {
			fmt.Printf("[%d, %d]: %q\n", i, j, students02[i][j])
		}

	}
	fmt.Printf("----------------------\n")
	for i, line := range students03 {
		for j, v := range line {
			fmt.Printf("[%d, %d]: %q\n", i, j, v)
		}
	}
}
