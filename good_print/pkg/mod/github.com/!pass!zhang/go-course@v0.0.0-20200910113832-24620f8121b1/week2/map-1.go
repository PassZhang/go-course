package main

import (
	"fmt"
)

// map的key 必须是支持== 或者!= 比较运算的类型，不可以是函数，map或者slice

// map的value 则没有任何的限制

// make 根据size 大小来初始化分配内存， 不过分配后的map长度为0。 如果size被忽略了，那么会在初始化分配内存的时候， 分配一个小尺寸的内存

// map的迭代顺序是随机的

func main() {
	// map在使用之前必须使用make 而不是new来创建的，值为nil的map是空的，并且不能赋值

	var number1 map[int]string
	number2 := make(map[int]string, 5)
	number3 := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}

	fmt.Printf("Number1: %v, len:%d \r\n", number1, len(number1))
	fmt.Printf("Number2: %v, len:%d \r\n", number2, len(number2))
	fmt.Printf("Number3-1: %v, len:%d \r\n", number3, len(number3))

	delete(number3, "c")
	fmt.Printf("Number3-2: %v, len:%d \r\n", number3, len(number3))

	number3["f"] = 6
	fmt.Printf("Number3-3: %v, len:%d \r\n", number3, len(number3))

	// 引用类型
	var number4 = map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 4}
	fmt.Println("Number4-1:", number4)
	number5 := number4
	number5["e"] = 5
	fmt.Println("Number4-2:", number4)
}
