// 当前程序的包名（一个可执行程序只有一个main包）
// 一般建议package的名称和目录名字保持一致
package main

// 导入其他包
// 如果缺少或未使用的包，程序都无法编译通过
import "fmt"

// 通过const 关键字来进行常量的定义
const number1 = 10

// 通过var 关键字来声明变量
var number2 int

// 数组
var number3 [5]int

// 切片
var number4 []int

// map
var number5 map[string]int

// 一般类型声明
type number6 int

// 结构声明
type number7 struct{}

// 接口声明
type number8 interface{}

// 只有通过func 关键字来进行函数的声明
// 只有package名称为main的包才可以包含main 函数
func main() {
	fmt.Println("Hello World!")
}
