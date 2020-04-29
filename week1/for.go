/*
for 循环是一个循环控制结构，可以执行指定次数的循环。

语法
Go 语言的 For 循环有 3 种形式，只有其中的一种使用分号。

和 C 语言的 for 一样：

for init; condition; post { }
和 C 的 while 一样：

for condition { }
和 C 的 for(;;) 一样：

for { }
init： 一般为赋值表达式，给控制变量赋初值；
condition： 关系表达式或逻辑表达式，循环控制条件；
post： 一般为赋值表达式，给控制变量增量或减量。
for语句执行过程如下：

1、先对表达式 1 赋初值；

2、判别赋值表达式 init 是否满足给定条件，若其值为真，满足循环条件，则执行循环体内语句，然后执行 post，进入第二次循环，再判别 condition；否则判断 condition 的值为假，不满足条件，就终止for循环，执行循环体外语句。

*/
// 循环求取10以内的数字相加结果
// package main
//
// import "fmt"
//
// func main() {
// sum := 0
// for i := 0; i <= 10; i++ {
// sum += i
// }
// fmt.Println(sum)
// }

// for-each range 循环
package main

import "fmt"

// func main() {
// 	strings := []string{"google", "runoob"}
// 	for i, s := range strings {
// 		fmt.Println(i, s)
// 	}
// }

// 	numbers := [6]int{1, 2, 3, 5}
// 	for i, x := range numbers {
// 		fmt.Printf("第 %d 位x的值= %d\n", i, x)
// 	}
// }

// func main() {
// 	sum := 0
// 	for i := 1; i <= 100; i++ {
// 		sum += i
// 	}
// 	fmt.Println(sum)
// }

// func main() {
// 	sum := 0
// 	for i := 1; i <= 100; i++ {
// 		sum += i
// 	}
// 	fmt.Println(sum)
// }

func main() {
	// 在控制台打印1..10
	for index := 1; index <= 10; index++ {
		fmt.Println(index)
	}

	//计算1..100的和

	total := 0

	for index := 1; index <= 100; index++ {
		total += index
	}
	fmt.Println(total)
}
