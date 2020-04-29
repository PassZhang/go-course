package main

import "fmt"

func main() {
	// 声明
	var names [10]string
	var scores [10]int

	fmt.Printf("%T, %T\n", names, scores)
	fmt.Printf("%q\n", names)
	fmt.Println(scores)

	// 字面量
	/*
		1. 指定数组长度：[length]type{v1, ,v2, ..., vlength}
		2. 使用初始化元素数量推到数组长度：[...]type{v1, v2, ..., vlength}
		3. 对指定位置元素进行初始化：[length]type{im:vm, ..., sin:in}
	*/
	var users [3]string = [3]string{"kk", "蜗牛", "ada"}
	bounds := [...]int{10, 20, 30, 40, 50}
	teachers := [5]string{"kk", "ada"}
	nums := [10]int{1: 10, 3: 30, 5: 50, 8: 80}

	fmt.Printf("%q\n", users)
	fmt.Println(bounds)
	fmt.Printf("%q\n", teachers)
	fmt.Println(nums)

	//  关系运算
	fmt.Println(bounds == [...]int{10, 20, 30, 40, 50})
	fmt.Println(bounds != [...]int{20, 10, 30, 40, 50})

	// 获取数组长度
	langs := [...]string{"GO", "Python", "c#", "kk", "蜗牛", "ada"}

	fmt.Println(len(langs))

	// 访问和修改
	fmt.Println(langs[0], langs[5], langs[len(langs)-1], langs[len(langs)-2], langs[len(langs)-3], langs[len(langs)-4])

	fmt.Printf("%q\n", langs)
	langs[0] = "golang"
	langs[1] = "Python"
	langs[2] = "C#"
	langs[3] = "KK"
	fmt.Printf("%q\n", langs)

	// 切片
	fmt.Printf("%T, %q\n", langs[1:4:4], langs[1:2:3])

	// 遍历，可以通过for+len+访问方式或者for-range范式对数组中元素进行遍历
	for i := 0; i < len(langs); i++ {
		fmt.Printf("%d: %q\n", i, langs[i])

	}
	fmt.Printf("---------------------\n")

	for i, v := range langs {
		fmt.Printf("%d: %q\n", i, v)
	}
	// 使用for-range遍历数组， range返回两个元素分别为数组元素索引和值

}
