// 遍历range
// package main

// import "fmt"

// func main() {
// 	// range 关键字用户for循环中迭代数组（array）、切片（slice）、通道（channel）或者集合（map）的元素

// 	number := [5]string{"a", "b", "c", "d", "e"}
// 	for k, v := range number {
// 		fmt.Println(k, v)
// 	}
// }

// 找出下面错误
package main

import "fmt"

type student struct {
	Name string
	Age  int
}

func main() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhouYi", Age: 24},
		{Name: "zhouEr", Age: 23},
		{Name: "zhouSan", Age: 37},
	}

	for _, stu := range stus {
		m[stu.Name] = &stu
		fmt.Println(stu)
	}
}
