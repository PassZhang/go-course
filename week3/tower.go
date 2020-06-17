package main

import "fmt"

/*
汉诺塔小游戏：
n个盘子 start（开始） end（终点） temp（借助）

n  start -> temp -> end
第一步 n - 1 start -> end -> temp
第二步 start -> end
第三步 n - 1 temp -> start -> end
终止条件 1 -> start -> end

*/

func tower(start, end, temp string, layer int) {
	if layer == 1 {
		fmt.Println(start, "->", end)
		return // 无返回值，表示函数结束
	}
	tower(start, temp, end, layer-1)
	fmt.Println(start, "->", end)
	tower(temp, end, start, layer-1)

}

func main() {
	tower("塔1", "塔2", "塔3", 3)
}
