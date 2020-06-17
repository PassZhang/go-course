package main

import "fmt"

// func main() {
// 	num1 := make([]int, 0, 10)
// 	fmt.Println(num1, len(num1), cap(num1))

// 	num2 := num1[2:3]
// 	fmt.Println(num2, len(num2), cap(num2))

// }

// func main() {
// 	complexArray1 := [3][]string{[]string{"d", "e", "f"},
// 		[]string{"g", "h", "i"},
// 		[]string{"j", "k", "l"},
// 	}

// 	complexArray2 := complexArray1

// 	// 若是修改数组中的切片的某个元素，会影响原数组
// 	complexArray2[0][0] = "a"
// 	fmt.Println(complexArray1)
// 	fmt.Println(complexArray2)

// 	//若是修改数组的某个元素就不会影响原数组
// 	complexArray2[0] = []string{"z", "z", "z"}
// 	fmt.Println(complexArray1)
// 	fmt.Println(complexArray2)
// }

// slice不是为删除而设计的数据结构
// package main

// import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8}

	// 删除元素4
	s = append(s[:3], s[5:]...)
	fmt.Println(s)

}
