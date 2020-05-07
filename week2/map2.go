// 映射
// 映射是存储一系列无序的key/value对，通过key来对value进行操作（增删查改）。映射的key只能为可以使用==运算符的值类型（字符串、数字、布尔、数组），value可以为任意类型

package main

import "fmt"

func main() {
	// 每个同学的成绩
	// key => ID value => 成绩

	var scores map[string]float64

	// 零值
	fmt.Printf("%T,  %#v\n", scores, scores)

	// 初始化
	// 字面量
	scores = map[string]float64{} //空的映射
	fmt.Printf("%T, %#v\n", scores, scores)

	scores = map[string]float64{"22": 80, "23": 90, "24": 100}
	fmt.Printf("%T, %#v\n", scores, scores)

	// make
	scores = make(map[string]float64) //map[string]float64{}
	fmt.Printf("%T, %#v\n", scores, scores)

	scores = map[string]float64{"22": 80, "23": 90, "24": 100}

	// 操作
	// 获取映射长度
	fmt.Println(len(scores))

	// key => value
	// 查找
	fmt.Println(scores["22"])
	fmt.Println(scores["23"])
	fmt.Println(scores["24"])

	// 判断key是否存在
	v, ok := scores["22"]
	fmt.Println(ok, v)

	v, ok = scores["23"]
	fmt.Println(ok, v)

	// 改
	scores["22"] = 100
	fmt.Println(scores)

	// 增
	scores["25"] = 110
	fmt.Println(scores)

	// 删除
	delete(scores, "22")
	fmt.Println(scores)

	// 遍历映射
	for v := range scores {
		fmt.Println(v, scores[v])

	}
	fmt.Println("----------------------------\n")
	for k, v := range scores {
		fmt.Println(k, v)
	}
}
