// 映射
// 映射是存储一系列无序的key/value对，通过key来对value进行操作（增删查改）。映射的key只能为可以使用==运算符的值类型（字符串、数字、布尔、数组），value可以为任意类型

package main

import "fmt"

func main() {
	// 声明
	// map声明需要指定组成元素key和value的类型，在声明后，会被初始化为nil，便是不存在的映射。

	var tels map[string]string
	var points map[[2]int]float64
	fmt.Printf("%T, %t, %v\n", tels, tels == nil, tels)
	fmt.Printf("%T, %t, %v\n", points, points == nil, points)

	//初始化
	// 1. 使用字面量初始化：map[ktype]vtype{k1:v1, k2:v2, ..., kn:vn}
	// 2. 使用字面量初始化空映射：map[ktype]vtype{}
	// 3. 使用make函数初始化
	// 	make(map[ktype]vtype),通过make函数创建映射

	fmt.Printf("----------------------------------\n")

	tels = map[string]string{"KK": "15200000000", "woniu": "185000000000"}
	fmt.Printf("%q\n", tels)

	fmt.Printf("----------------------------------\n")
	points = map[[2]int]float64{{1, 2}: 3, {4, 5}: 6}
	fmt.Println(points)

	scores := map[string]int{"kk": 80, "woniu": 90}
	fmt.Println(scores)

	heighs := map[string]string{}
	fmt.Println(heighs)

	weights := make(map[string]float64)
	fmt.Println(weights)
	fmt.Printf("----------------------------------\n")

	// 操作
	// 1. 获取元素的数量
	// 使用len函数获取映射元素的数量
	fmt.Println(len(tels), len(points), len(scores), len(heighs), len(weights))
	fmt.Printf("----------------------------------\n")

	//访问
	students := map[int]string{1: "kk", 2: "wonin"}
	students01 := map[int]map[string]string{1: map[string]string{"name": "kk", "tel": "1520000000"}}
	fmt.Printf("%v, %q, %q\n", students, students[1], students[3])
	fmt.Printf("%v, %q, %q, %t\n", students01, students01[1], students01[3], students01[3] == nil)
	// 当访问 key 存在与映射时则返回对应的值，否则返回值类型的零值

	// 判断key 是否存在
	// 通过key访问元素可以接收两个值，第一个值为value，第二个值为bool类型表示元素是否存在，若存在为true，否则为false
	student, ok := students[1]
	fmt.Printf("%t, %v\n", ok, student)

	student, ok = students[2]
	fmt.Printf("%t, %v\n", ok, student)

	student01, ok := students01[1]
	fmt.Printf("%t, %v\n", ok, student01)

	student01, ok = students01[2]
	fmt.Printf("%t, %v\n", ok, student01)

	// 修改和增加
	// 使用key对映射复制时当key存在修改key对应的value，若key不存在则增加key和value
	students[1] = "KK"
	students[3] = "WD"

	fmt.Println(students)

	students01[1]["tel"] = "1520000000"                                                     //key存在，修改
	students01[1]["addr"] = "西安市"                                                           // key不存在，增加
	students01[2] = map[string]string{"name": "woniu", "tel": "15800000000", "addr": "北京市"} //key不存在，增加
	fmt.Println(students01)

	students01[2] = map[string]string{"name": "woniu", "tel": "15811111111", "addr": "北京市"} // key存在，修改
	fmt.Println(students01)
	fmt.Printf("----------------------------------\n")

	// 删除
	// 使用delete函数删除映射中已经存在的key
	delete(students, 1)
	delete(students, 4)
	fmt.Println(students, 4)

	delete(students01[1], "addr")
	delete(students01, 2)
	fmt.Println(students01)
	fmt.Printf("----------------------------------\n")

	//遍历
	//可以通过for-range对映射中的元素进行遍历，range返回两个元素分别为映射的key和value
	for k, v := range students {
		fmt.Printf("%v:%v\n", k, v)

	}
}
