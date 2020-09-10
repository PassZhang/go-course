// 切片不是引用类型吗？为什么我给number2追加一个元素，number1的值却没有改变（容量为10，追加的时候也没有扩容）

// package main

// import "fmt"

// func main() {
// 	number1 := make([]int, 0, 10)

// 	number2 := number1
// 	number2 = append(number2, 5)

// 	fmt.Println(number1)
// 	fmt.Println(number2)
// }

package main

import "fmt"

func main() {
	// 两个slice 虽然公用一个数组地城，但是他们的len和cap是独立的。
	number1 := make([]int, 0, 10)
	number2 := number1

	number1 = append(number1, 1)
	fmt.Println(number1)
	fmt.Println(number2)

	number2 = append(number2, 2)
	fmt.Println(number1)
	fmt.Println(number2)
}
