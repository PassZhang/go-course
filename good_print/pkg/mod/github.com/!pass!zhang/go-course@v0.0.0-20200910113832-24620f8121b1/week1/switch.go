// switch 语句用于基于不同条件执行不同动作，每一个 case 分支都是唯一的，从上至下逐一测试，直到匹配为止。

/*
switch var1 {
    case val1:
        ...
    case val2:
        ...
    default:
        ...
}
*/

package main

import "fmt"

func main() {
	grade := ""
	switch {
	case grade == "A":
		fmt.Printf("优秀！\n")
	case grade == "B", grade == "C":
		fmt.Printf("良好\n")
	case grade == "D":
		fmt.Printf("及格\n")
	case grade == "F":
		fmt.Printf("不及格\n")
	default:
		fmt.Printf("差\n")
	}
	fmt.Printf("你的等级是 %s\n", grade)
}
