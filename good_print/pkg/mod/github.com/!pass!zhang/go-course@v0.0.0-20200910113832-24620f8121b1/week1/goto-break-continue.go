// 跳转语句
package main

import "fmt"

func main() {
	// 支持函数内部goto 跳转, continue 跳出当前循环进入下一次循环，break终止循环体
	// break he continue都可以配合标签，在多机嵌套件循环间跳出，这和goto调整执行位置完全不同。
	// 通常建议往后沟通，避免死循环

	for number := 1; number < 5; number++ {
		if number == 3 {
			break
		}
		fmt.Println("break:", number)
	}

	for number := 1; number < 5; number++ {
		if number == 3 {
			continue
		}
		fmt.Println("continue:", number)
	}

lable1:
	for {
		for {
			// 配合标签跳出最外层循环
			// 标签名是大小写敏感的
			break lable1
		}
	}
	fmt.Println("Hello World1")
	goto lable2
	fmt.Println("Hello World2")

lable2:
	fmt.Println("Hello World3")
}
