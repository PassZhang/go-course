// for Range 用于变量可迭代对象中的每个元素例如字符串，数组，切片，映射，通道等
// 针对包含Unicode字符的字符串遍历是需要使用for-range，range返回两个元素为字符索引和rune字符，可通过空白标识符忽略需要接收的变量

package main

import "fmt"

func main() {
	text := "我爱中国"
	for i, e := range text {
		fmt.Printf("%d: %c\n", i, e)
	}

	for _, e := range text {
		fmt.Printf("%c\n", e)
	}

}