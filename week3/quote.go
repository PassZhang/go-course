package main

import "fmt"

func main() {
	// age := 30
	// tmpAge := age

	// tmpAge = 31
	// fmt.Println(age, tmpAge)
	// fmt.Printf("%p, %p\n", &age, &tmpAge) // 打印变量指针地址

	// users := make([]string, 10)
	// tmpUsers := users

	// tmpUsers[0] = "kk"
	// fmt.Printf("%#v, %#v\n", users, tmpUsers) // 打印变量
	// fmt.Printf("%p, %p\n", users, tmpUsers)   // 打印变量指针地址
	// fmt.Printf("%p, %p\n", &users, &tmpUsers) // 打印变量的切片指针地址

	// 上面这个测试主要是为了区分值类型和引用类型
	// 值类型对应的有int, float, point, 数组, 结构体
	// 引用类型对应的有：切片，映射，接口。

	// 数组示例
	array := [3]int{}
	tmpArray := array
	tmpArray[0] = 10
	fmt.Println(array, tmpArray)
	fmt.Printf("%p, %p\n", &array, &tmpArray)
	// 切片示例
	users := make([]string, 10)
	tmpUsers := users
	tmpUsers[0] = "kk"
	fmt.Printf("%#v, %#v\n", users, tmpUsers)
	fmt.Printf("%p, %p\n", users, tmpUsers)
	fmt.Printf("%p, %p\n", &users, &tmpUsers)
}

//总结：
// 	// 上面这个测试主要是为了区分值类型和引用类型
// 值类型对应的有int, float, point, 数组, 结构体
// 引用类型对应的有：切片，映射，接口。

// 引用类型：
//假如user是一个切片变量，user切片变量里面存储的是切片的地址(slice address)，切片的地址里面又存储了数组的地址(array address)。同时user切片本身也有一个地址（user address）
//关于变量指针来说，指针本身也是一个值，也有自己的地址，两个不同的指针地址不同，但是他们可以存储相同的值，而他们存的值就是另外一个变量的地址，所以根据他们的值可以找到这个变量
// 切片引用切片，他们的变量地址不一样，是因为变量的指针的地址是不一样的，但是切片中的地址都是一样的。所以如果a切片引用了b切片，再去修改b切片，a切片也会被修改的。因为切片内地址是一样的。
