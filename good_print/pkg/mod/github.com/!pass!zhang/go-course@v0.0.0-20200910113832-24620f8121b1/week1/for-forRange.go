// for 和 for range 区别
// range 可以将每个独立的Unicode码点分离出来
//for range只能用于遍历系统自带的array slice map 和chan ,数据结构就不能用range

package main

import (
	"fmt"
	"reflect"
	"time"
)

type student struct {
	name    string
	age     int
	content string
}

func main() {
	// for range 只能用于遍历系统自带的array slice map 和chanal， 数据结构就不能用for range
	s1 := student{name: "张三", age: 18, content: "Hello World"}

	for k := 0; k < reflect.TypeOf(s1).NumFiled(); k++ {
		fmt.Println(reflect.TypeOf(s1).Field(k).Name, "=>", reflect.ValueOf(s1).Field(k))

	}

	number1 := []int{4, 8, 12, 16, 20}
	// len := len(number1) 用变量代替，不用每次循环都计算number1 长度
	for i := 1; i < len(number1); i++ {
		if i == 0 {
			number1 = append(number1, 24)
		}
		fmt.Print(&number1[i], " - ")
	}
	fmt.Println("")

	number3 := []int{4, 8, 12, 16, 20}
	for _, v := range number3 {
		go func() {
			fmt.Print(v, "-")
		}()
	}

	time.Sleep(time.Second)
}
