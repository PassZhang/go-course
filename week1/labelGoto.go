// 可以通过goto语句任意跳转到当前函数指定的label位置

package main

import "fmt"

// func main() {
// 	sum := 0
// 	i := 1

// START: // 定义START标签
// 	if i > 100 {
// 		goto END // 跳转到END标签
// 	} else {
// 		sum += i
// 		i++
// 		goto START // 跳转到START标签
// 	}
// END: // 定义END标签
// 	fmt.Println(sum)
// }

func main() {
CI:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j == 3 {
				continue CI //跳过外层循环
			}
			fmt.Println(i, j)
		}
	}

	// BI:
	// 	for i := 0; i < 5; i++ {
	// 		for j := 0; j < 5; j++ {
	// 			if j == 3 {
	// 				break BI // 跳出外层循环
	// 			}
	// 			fmt.Println(i, j)
	// 		}
	// 	}
}
