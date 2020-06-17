// slice 区块
package main

import "fmt"

func main() {
	//区间这一块不是很好理解
	s1 := make([]int, 0, 5)
	fmt.Println(s1)
	s2 := s1[2:5]
	s2 = s1[2:] //右区间未指定长度按切表的长度填充
	s2 = s1[2:6]
	s2 = s1[5:5]

	fmt.Println(s2)
}
