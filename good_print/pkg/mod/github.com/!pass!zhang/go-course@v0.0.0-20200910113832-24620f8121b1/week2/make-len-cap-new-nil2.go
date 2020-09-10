// slice, map, channel, pointer, func, interface 零值均为nil

package main

import "fmt"

// * 如果变量没有显式初始化，则被隐式的赋予其类型的零值（zero value）, 数值类型是0， 字符串类型是空字符串""
// slice map pointer, func, interface 零值均为nil

// var field1 []int
// var field2 map[int]string
// var field3 chan int
// var field4 *int
// var field5 func()
// var field6 interface{}

// func main() {
// 	if field1 == nil {
// 		fmt.Println("field1 zero value is nil")
// 	}

// 	if field2 == nil {
// 		fmt.Println("field2 zero value is nil ")
// 	}

// 	if field3 == nil {
// 		fmt.Println("field3 zero value is nil")
// 	}

// 	if field4 == nil {
// 		fmt.Println("field4 zero value is nil")
// 	}

// 	if field5 == nil {
// 		fmt.Println("field5 zero value is nil")
// 	}

// 	if field6 == nil {
// 		fmt.Println("field6 zero value is nil")
// 	}

// 	fmt.Println(field1)
// 	fmt.Println(field2)
// 	fmt.Println(field3)
// 	fmt.Println(field4)
// 	fmt.Println(field5)
// 	fmt.Println(field6)
// }

///////////
var err error

func main() {
	fmt.Println(err)
}

// 看出来打印结果是nil，但是为什么上面没有写呢？
// 因为error 不是基本类型，是自定义类型，底层是interface

// go源码中这样写到
/*
type error interface {
	Error() string
}
*/
