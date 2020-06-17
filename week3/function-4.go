package main

import (
	"fmt"
)

func Number(i int) func() {
	return func() {
		fmt.Println(i)
	}
}
func main() {
	number1 := make([]func(), 3)
	number2 := make([]func(), 3)

	// TODO:1
	// 闭包可以直接使用外部的变量
	for i := 0; i < 3; i++ {
		number1[i] = func() {
			fmt.Println(i)
		}
	}

	fmt.Println("-----TODO1-----")
	for _, v := range number1 {
		v()
	}
}
	//TODO:2
	fmt.Println("-----TODO2-----")
	for _, v := range number2 {
		v()
	}
}
