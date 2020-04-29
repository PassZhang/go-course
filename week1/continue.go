/*
contineu语句
Go 语言的 continue 语句 有点像 break 语句。但是 continue 不是跳出循环，而是跳过当前循环执行下一次循环语句。

for 循环中，执行 continue 语句会触发 for 增量语句的执行。

在多重循环中，可以用标号 label 标出想 continue 的循环。
*/

// 变量a大于15的时候跳出循环
package main

import "fmt"

func main() {
	// 定义局部变量
	var a int = 10

	// for 循环
	for a < 20 {
		if a == 15 {
			// 使用continue语句跳出循环
			a++
			continue
		}
		fmt.Printf("a 的值为： %d\n", a)
		a++

	}
}
