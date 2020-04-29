// 使用 fallthrough 会强制执行后面的 case 语句，fallthrough 不会判断下一条 case 的表达式结果是否为 true。

package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("1、case条件语句为false")
		fallthrough
	case true:
		fmt.Println("2、case条件语句为true")
		fallthrough
	case false:
		fmt.Println("3、case条件语句为false")
		fallthrough
	case true:
		fmt.Println("4、case条件语句为true")
	case false:
		fmt.Println("5、case条件语句为false")
		fallthrough
	default:
		fmt.Println("6、默认 case")
	}
}

/* 
1. 支持多条件匹配
2. 不同的case之间不使用break分隔，默认只会执行一个case。
3. 如果想要执行多个case，需要使用fallthrough关键字，也可用break终止。
*/