// 数值类型
// 1. 整型
// go提供了5种有符号，5种无符号、1种指针、1种3单字节、1种单个unicode字符(unicode码点)，共13种整数类型，零值均为0.

package main

import "fmt"




func main() {
	var n1, n2, n3 int = 1, 8, -2 
 	var n4 uint = 2

	// 算术运算
	fmt.Println(n1 + n2) //9 
	fmt.Println(n1 + n2) //-7
	fmt.Println(n1 * n2) //8 
	fmt.Println(n1 / n2) //0 
	fmt.Println(n1 % n2) //1

	n1++ //2
	n2-- //7
	fmt.Println(n1, n2) //2 7 

	// 关系运算符
	fmt.Println(n1 > n2)
	fmt.Println(n1 >= n2)
	fmt.Println(n1 < n2)
	fmt.Println(n1 <= n2)
	fmt.Println(n1 == n2)
	fmt.Println(n1 != n2)

	// 位运算
	// 2 => 0010
	// 7 => 0111
	// -2 => 1110
	// n4 => uint 2 0010
	fmt.Println(n1 & n2) // 0010
	fmt.Println(n1 | n2) // 0111
	fmt.Println(n1 ^ n2) // 0101
	fmt.Println(n1 &^ n2) // 0000
	fmt.Println(n2 << n4) //0001 1100
	fmt.Println(n2 >> n4) //0001
	fmt.Println(n3 << n4) //1111 1000
	fmt.Println(n3 >> n4) //1111 1111
} 
