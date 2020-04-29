/*
break 语句
go 语言中break 语句用于以下两方面
1. 用于循环语句中跳出循环，并开始执行循环之后的语句。
2. break在switch(开关语句)中在执行一条case后跳出语句的作用。
3. 在多种循环中，可以用标号bale标出想break的循环。
*/

// 变量a大于15的时候跳出循环
package main

import "fmt"

func main() {
	// 定义局部变量
	var a int = 10

	// for 循环
	for a < 20 {
		fmt.Printf("a 的值为： %d\n", a)
		a++
		if a > 15 {
			// 使用break语句跳出循环
			break
		}
	}
	for j := 0; j < 2; j++ {
		for i := 0; i < 10; i++ {
			if i == 5 {
				break
			}
			fmt.Println(i)
		}
	}
}
