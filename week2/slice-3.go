package main 

import "fmt"

func main()  {
	number1 := make([]int, 1, 2)

	fmt.Printf("%p\n", number1)     //图1
	fmt.Printf("%p\n", &number1[0]) //图2
	fmt.Printf("%p\n", &number1)    //图3 切片变量的地址(ptr len cap)


	// https://box.kancloud.cn/ab4e434f7fbfcbd78c57e6f8a6c3f444_886x579.jpg
}