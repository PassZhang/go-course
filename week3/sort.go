package main

import (
	"fmt"
	"sort"
)

func main() {

	// char => counter
	// A => 65
	// B => 66
	// C => 67
	// D => 68

	stats := [][]int{{'A', 3}, {'B', 2}, {'C', 1}, {'D', 2}}
	// 按使用出现次数进行排序
	// {B, 2}, {D, 2} => 稳定的（默认）
	// {D, 2}, {B, 2} => 不稳定

	// sort.Slice(stats, func(i, j int) bool {
	// 	return stats[i][1] > stats[j][1]
	// })
	// fmt.Println(stats)
	// sort.SliceStables 不稳定排序

	// 升序 <=
	// 降序 >=
	index := sort.Search(len(stats), func(i int) bool {
		return stats[i][1] <= 1
	})
	fmt.Println(index)
	fmt.Println(stats)
}
