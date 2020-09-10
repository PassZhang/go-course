// make len cap new nil 
/*
make 和 new 的区别：
1. make 只能用来分配及初始化 类型为slice，map，chanal的数据，new可以分配任意类型的数据
2. new分配返回的是指针，即类型*T，make返回引用，即T
3. new分配的空间被清零，make分配后，会进行初始化

*/ 

// make
/* 
格式：func make(Type, size IntegerType) Type

1. 内建函数make分配和初始化一个slice或map或者channel对象，并且只能是这三种对象
2. 和new类似，第一个参数也是一个类型而不是一个值，不同的是make返回类型的引用而不是指针，而返回值也依赖于具体传入的类型
3. slice：第二个参数指定他的长度，此时他的容量和长度相同，可以用第三参数来来指定不同的容量大小，但不能小于他的长度（第二个参数）
4. map：根据size大小来初始化分配内存，不过分配后的map长度为0.如果size被忽略了，那么会在初始化分配内存的的时候分配一个小尺寸的内存。
channel：管道缓冲区依据缓冲区容量倍初始化。如果容量为0或者被忽略，管道是没有缓冲区的。

*/

// len 
// 格式：func len(v Type) int

// cap 
// 格式：func cap(v Type) int
// map变量被创建后，你可以指定map的容量，但是不可以在map上使用cap()方法

// new 
// 格式: func new(Tyep) *Type
// 内建函数new用来分配内存，它的第一个参数是一个类型，不是一个值，他的返回值是一个指向新分配类型零值的指针
// package main

// import "fmt"

// func main()  {
// 	number1 := [5]int{}
// 	number2 := new([5]int)

// 	fmt.Println(number1)
// 	fmt.Println(number2)
// }


// package main 

// import "fmt"

// type person struct {
// 	name string
// 	age int

// }


// func main()  {
// 	p1 := person{}
// 	p2 := &person{}
// 	p3 := new(person)

// 	fmt.Println(p1) //返回类型
// 	fmt.Println(p2) //返回指针
// 	fmt.Println(p3) // 和p2一样
// }



// nil 
// nil 是一个预定义标识符，其代表（用作）一些类型的零值；这些类型包括：pointer, channel, func, interface, map, slice

package main 

import "fmt" 

func main()  {
	var n1 []int

	var n2 map[int]string

	var n3 chan int

	if n1 == nil {
		fmt.Println("n1")
	}

	if n2 == nil {
		fmt.Println("n1")

	}

	if n3 == nil {
		fmt.Println("n1")
	}

	
}

