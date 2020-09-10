package main

import "fmt"

func main() {
	isGirl := false

	fmt.Printf("%T, %#v\n", isGirl, isGirl)
	//%T 是显示类型，%v是显示结果

	//操作
	// 逻辑运算
	a, b, c, d := true, true, false, false
	// 与：左操作数与右操作数都为true,结果为true &&

	fmt.Println("a ,b:", a && b) // true && true : true
	fmt.Println("a ,c:", a && c) // true && false : false
	fmt.Println("c ,b:", c && b) // false && true : false
	fmt.Println("c ,d:", c && d) // false && false : false

	// 或者： 左右操作数只要有一个为true，结果为true ||
	fmt.Println("a, b:", a || b) // true || true : true
	fmt.Println("a, c:", a || c) // true || false : true
	fmt.Println("c, b:", c || b) // flase || true : true
	fmt.Println("c, d:", c || d) // flase || flase : flase

	// 非：取反 true => false, false => true !

	fmt.Println("!a", !a) // !true : false
	fmt.Println("!c", !c) // !false : true

	// 或者： 左右操作数只要有一个为true，结果为true ||
	fmt.Println("a, b:", a || b) // true || true : true
	fmt.Println("a, c:", a || c) // true || false : true
	fmt.Println("c, b:", c || b) // flase || true : true
	fmt.Println("c, d:", c || d) // flase || flase : flase

	// 关系：
	fmt.Println(a == b) // true == true : true
	fmt.Println(a != c) // true != false : true
	fmt.Println(a == c) // true == false : false
	fmt.Println(c != b) // flase != true : true

	fmt.Printf("%t, %t\n", a, c)

	var bb bool //定义为空变量，默认为flase
	fmt.Println(bb)

}
