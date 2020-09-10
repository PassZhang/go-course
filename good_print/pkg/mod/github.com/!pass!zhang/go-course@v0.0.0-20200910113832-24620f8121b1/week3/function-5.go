package main

import "fmt"

/*
1. 栈空间变化（堆栈从左到右，增加一个峰值后再计算从右到左缩小）
demo4(5)
5 + demo4(4)
5 + (4 + demo4(3))
5 + (4 + (3 + demo4(2)))
5 + (4 + (3 + (2 + demo4(1))))
5 + (4 + (3 + (2 + 1)))
5 + (4 + (3 + 3))
5 + (4 +  6)
5 + 10
15
*/

func demo4(val int) int {
	if val > 0 {
		return val + demo4(val-1)

	} else {
		return val
	}
}

func main() {
	fmt.Println(demo4(5))
}
