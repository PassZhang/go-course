// 多维切片
// 切片的元素也可以是切片类型，此时称为多维切片
package main

import "fmt"

func main()  {
	// 声明&初始化
	points := [][]int{{1, 1}, {1, 2, 3}}
	fmt.Printf("%T, %v, %v, %d\n", points, points, points[0], points[0][0])
	fmt.Printf("-----------------------------\n")

	// 修改
	points[0] = []int{2, 2}
	points[1][1] = 3  // 将第二个元素的第二个数修改成3
	fmt.Println(points)


	// 切片
	fmt.Println(points[0:1])

	// 遍历
	for i := 0; i < len(points); i++ {
		for j := 0; j < len(points[i]); j++ {
			fmt.Printf("[%d, %d]: %v\n", i, j, points[i][j])
		}
	}
	fmt.Printf("-----------------------------\n")

	for i, line := range points{
		for j, v := range line{
			fmt.Printf("[%d, %d]: %v\n", i, j, v)
		}
	}

	//append
	points = append(points, []int{2, 3, 1})
	points[0] = append(points[0], 1)
	fmt.Println(points)

	// copy 
	points2 := [][]int{{}, {}}
	copy(points2, points)
	fmt.Println(points2)
}