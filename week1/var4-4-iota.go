// iota
// package main 
package main 

import "fmt"

// iota 是常量的计数器从0开始，组中每定义一个常量自动递增1
// 每遇到一个const关键字，iota就会重置为0

const (
	GB int = 1 << (iota * 10)
	MB int = 1 << (iota * 10)
	KB int = 1 << (iota * 10)
)

const (
	a int = iota 
	b
)

func main()  {
	fmt.Println(GB, MB, KB)
	fmt.Println(a, b)
}