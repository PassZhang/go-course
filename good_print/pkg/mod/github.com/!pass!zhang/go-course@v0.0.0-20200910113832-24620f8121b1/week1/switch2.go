// 选择语句switch
/* 
支持三种形式
1. switch init; variable {

}

2. switch variable {

}

switch {}
*/

package main 

import "fmt"

func main()  {
	switch number1 := 3; number1 {
	case 1:
		fmt.Println(1)
case 3:
		fmt.Println(3)
case 5:
	fmt.Println(5)
}
// 不需要写break, 一旦条件符合自动结束


// 多个case 执行同一个结果（错误）
// number2 :=7 
// switch number2 {
// case:7
// case:8
// fmt.Println("number2 - 7 or 8")
// }

//多个case执行同一结果
number3 := 7 
switch number3 {
case 7, 8:
	fmt.Println("number3 - 7 or 8")
}

number4 := 7
switch {
case number4 == 7:
	fmt.Println(7)
	// 如希望继续执行下一个case，可以使用fallthrough语句
	fallthrough
case number4 == 8:
	fmt.Println(8)
}

}