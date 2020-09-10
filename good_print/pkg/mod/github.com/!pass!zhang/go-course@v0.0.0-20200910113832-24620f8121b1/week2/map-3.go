package main

import (
	"fmt"
)

// func main() {
// 	var m map[string]int
// 	if _, ok := m["first"]; !ok {
// 		m = make(map[string]int, 5)
// 	}

// 	m["first"] = 11111
// 	fmt.Println(m)

// }

// delete

func main() {
	// 删除对应的键值对，如果集合为空或者找不到对应的键，则误操作
	number := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}
	fmt.Println(number)

	delete(number, "b")
	fmt.Println(number)

	delete(number, "f")
	fmt.Println(number)

}

// 为什么map会提示写并发和读的错误，变量又不会？？？
// 在标准库的源代码中能看到，map内部进行了包含，如果发现了并发访问就会报fatal error （不是panic， 所以无法refcover）。 普通的变量类型没有这个机制。
