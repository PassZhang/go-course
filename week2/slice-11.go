// 例子2
// len才是代表了slice的有效元素个数（而不是cap），只有有效元素才能进行读写

package main 

import "fmt"

func main()  {
	number1 := make([]int, 0, 10)

	// number1[0] = 1

	fmt.Println(number1)
}
