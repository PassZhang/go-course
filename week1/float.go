// 浮点数用于表示带小数的数字，Go 提供 float32 和 float64 两种浮点类型
package main

import "fmt"

func main() {
	var (
		f1 float32 = 3.1415926
		f2 float32 = 3E-3
		//f3 float64 = 3.1E10
	)
	// 算术运算符：
	fmt.Println(f1 + f2)
	fmt.Println(f1 - f2)
	fmt.Println(f1 * f2)
	fmt.Println(f1 / f2)
	f1++
	f2--
	fmt.Println(f1, f2)

	// 关系运算
	fmt.Println(f1 > f2)
	fmt.Println(f1 >= f2)
	fmt.Println(f1 < f2)
	fmt.Println(f1 <= f2)

	// 赋值运算
	f1 += f1
	fmt.Println(f1)
	f1 -= f1
	fmt.Println(f1)
	f1 *= f1
	fmt.Println(f1)
	f1 /= f1
	fmt.Println(f1)

	// 类型转换
	fmt.Println("%T %T\n", f1, float64(f1))
}
