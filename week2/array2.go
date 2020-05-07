package main

import "fmt"

func main() {
	var names [3]string
	var signIns [3]bool
	var scores [3]float64

	//打印值类型, 返回对应值的类型
	fmt.Printf("%T\n", names)
	fmt.Printf("%T\n", signIns)
	fmt.Printf("%T\n", scores)

	//打印零值，返回对应值类型的零值
	fmt.Printf("%#v\n", names)
	fmt.Printf("%#v\n", signIns)
	fmt.Printf("%#v\n", scores)

	// 赋值
	// 字面量 => 0, 1, 2, 3, n-1
	// 第一种
	//	names = [3]string{"05-牛", "37-猫", "21-羊"}

	//	fmt.Printf("%#v\n", names)

	// names = [1]string{"05-牛"}
	// 第二种
	// names = [...]string{"05-牛", "37-猫", "21-羊", "24-狗"}

	// fmt.Printf("%#v\n", names)

	testnames := [...]string{"05-牛", "37-猫", "21-羊", "24-狗"}
	fmt.Printf("%#v\n", testnames)

	// 第三种
	names = [3]string{1: "kk"} //[3]string{"", "kk", ""}
	fmt.Printf("%#v\n", names)

	// 操作
	// 关系运算 == !=
	fmt.Println(names == [3]string{})
	fmt.Println(names == [3]string{1: "kk"})

	// 元素
	// 访问 & 修改 索引(0, 1, ..., 0-1)
	fmt.Printf("%q\n", names[0])
	names[0] = "03-牛"
	names[2] = "37-猫"
	fmt.Printf("%T %#v\n", names, names)

	// 函数len
	fmt.Println(len(names))

	// 遍历
	for i := 0; i < len(names); i++ {
		fmt.Println(i, names[i])
	}

	fmt.Println("---------------------------------------")
	for i, v := range names {
		fmt.Println(i, v)
	}

	// 定义一个数组，每个元素也是数组
	// 二维数组
	d2 := [3][2]int{0: [2]int{3, 4}, 1: [2]int{1, 2}, 2: [2]int{1: 5}}
	//[2]int = {0, 0}
	//{[2]int, [2]int, [2]int}
	//{{0, 0}, {0, 0}, {0, 0}}
	fmt.Printf("%#v\n", d2)
	fmt.Printf("%#v\n", d2[0])
	fmt.Printf("%#v\n", d2[0][0])
	fmt.Printf("%#v\n", d2[1][1])

}