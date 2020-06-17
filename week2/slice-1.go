// 切片slice
// 例子1当切片长度小于1024时，扩容为原来的2倍，大于或者等于1024时，扩容1.25倍
// slice[start:end:cap]

package main

import "fmt"

/*
1. 定义一个切片格式：make([]Type, length长度,capacity容量)
2. 一般是使用make()创建：用len获取切片长度，cap获取切片容量
3. 一个切片在未初始化之前默认为nil，长度为0
4. 切片是基于数组类型做的一层封装，自动扩容
5. 与数组相比切片的的长度是不固定的，可以追加元素，在追加时可以使切片的容量增大
*/

func main() {
	var number1 []string
	number2 := []string{"A", "B", "C"}
	number3 := make([]int, 2, 2)
	number4 := append(number3, []int{3, 4, 5}...) // number4 := append(number3, number3...) 用省略号自动展开切片，以使用每个元素

	// 长度和容量 切片再切片
	// cap 省略时 默认等于len
	// number5... 展开number5切片成一个一个元素
	// var a [5]int = [5]int{5, 4, 3, 2, 1} // soft.Ints(a[:])

	fmt.Println("Number1:", number1)
	fmt.Println("Number2:", number2)
	fmt.Printf("Number3: %d, len:%d, cap:%d\n", number3, len(number3), cap(number3))
	fmt.Printf("Number4: %d, len:%d, cap:%d\n", number4, len(number4), cap(number4))

	number5 := []string{"A", "B", "C"}
	number6 := []string{"D", "E"}
	number7 := copy(number6, number5) //等位替换原切片（返回替换成功的个数）,用后面的值替换前面的值
	fmt.Printf("Number5: %d\n", number5)

	fmt.Printf("Number6: copy success number %d --- %s\n", number7, number6)
	fmt.Println(number6)

}
