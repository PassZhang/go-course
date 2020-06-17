package main

import "fmt"

func main() {
	//[start:end] [start:] [:end] [:]  -- 第三个参数: 递增
	number8 := []int{9, 12, 35, 41, 95, 60, 17, 66, 9, 10}
	fmt.Println("Number8-1:", number8[3:6])

	//切片是引用类型(对数组的引用)
	fmt.Printf("Number8-2: %d, Addr:%p, Len:%d, Cap:%d \r\n", number8, number8, len(number8), cap(number8))

	number9 := number8
	number9[3] = 77
	fmt.Printf("Number8-3: %d, Addr:%p, Len:%d, Cap:%d\n", number8, number8, len(number8), cap(number8))

	fmt.Println("")
	number8 = append(number8, 11)
	number9[3] = 88
	fmt.Printf("Number8-4: %d, Addr:%p Len:%d, Cap:%d\n", number8, number8, len(number8), cap(number8))
	fmt.Printf("%#v\n", number8)
	fmt.Printf("%v\n", number8)

	// 1. 使用 append 对切片增加一个或多个元素并返回修改后切片，当长度在容量范围内时只增加长度，容量和底层数组不变。当长度超过容量范围则会创建一个新的底层数组并对容量进行智能运算(元素数量<1024 时，约按原容量 1 倍增加，>1024 时约按原容量 0.25倍增加), 这也就是number8-4 打印cap:20的原因

}
