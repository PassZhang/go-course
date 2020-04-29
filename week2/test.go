package main

import "fmt"

func main() {
	var nums []int = []int{1, 2, 3, 4, 5, 101, 5, 100, 2, 10, 101, 46, 100}

	Secondmax := SecondMaxNum(nums)
	fmt.Println("第二大的数是：", Secondmax)

	Secondmax = SecondMaxNum2(nums)
	fmt.Println("第二大的数是：", Secondmax)

	BubbleSort(nums)
	fmt.Println(nums)

	InsertSort(nums)
	fmt.Println(nums)
}
