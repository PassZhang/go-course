// 无序输出和有序输出
package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var number6 map[int][]int
	number6 = make(map[int][]int, 5)

	for i := 0; i < 5; i++ {
		_, ok := number6[i]
		if !ok {
			number6[i] = make([]int, 0, 3)
		}
		number6[i] = append(number6[i], r.Intn(100), r.Intn(100), r.Intn(100))
	}
	fmt.Println("无序输出")
	for k, v := range number6 {
		fmt.Printf("K: %d, v: %d \r\n", k, v)

	}

	fmt.Println("有序输出")

	var number7 []int
	for k, _ := range number6 {
		number7 = append(number7, k)

	}

	sort.Ints(number7)

	for _, v := range number7 {
		fmt.Printf("k: %d, v: %d \r\n", v, number6[v])

	}

}
