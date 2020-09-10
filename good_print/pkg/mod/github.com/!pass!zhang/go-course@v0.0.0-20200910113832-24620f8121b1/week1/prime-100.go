package main

import "fmt"

func main() {
	var count int
	var flag bool
	count = 1
	for count < 100 {
		count++
		flag = true
		for tmp := 2; tmp < count; tmp++ {
			if count%tmp == 0 {
				flag = false
			}
		}
		// 每一个 if else 都需要加入括号 同时else 位置不能在新一行
		if flag == true {
			fmt.Println(count, "素数")
		} else {
			continue
		}
	}
}
