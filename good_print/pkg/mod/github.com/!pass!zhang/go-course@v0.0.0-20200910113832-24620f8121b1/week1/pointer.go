// 每个变量在内存中都有对应存储位置（内存地址），可以通过&运算符获取。指针用来存储变量地址的变量
// 声明
package main

import "fmt"

func main() {
	// 零值
	var (
		pointerInt    *int
		pointerString *string
	)

	fmt.Printf("%T %#v\n", pointerInt, pointerInt)
	fmt.Printf("%T %#v\n", pointerString)

	// 赋值
	// 使用现有变量 取地址 &name

	age := 32

	age2 := age

	fmt.Printf("%T, %#v\n", &age, &age)
	fmt.Printf("%T, %#v\n", &age2, &age2)
	pointerInt = &age

	fmt.Println(pointerInt)
	fmt.Println(*pointerInt)

	*pointerInt = 330000
	fmt.Println(age, age2)

	pointerString = new(string)
	fmt.Printf("%#v, %#v\n", pointerString, *pointerString)

	pp := &pointerString
	fmt.Printf("%T\n", pp)
	**pp = "kk"
	fmt.Println(*pointerString)

}
