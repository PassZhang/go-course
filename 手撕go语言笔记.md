# 初识go

## 简介
go是一门开发源码的编程语言，可容易的构建简单、可靠、和高效的软件

## 历史
Go 语言是由谷歌的开发工程师(罗伯特·格瑞史莫、罗勃·派克、肯·汤普逊等)于 2007 年
开始设计，利用 20%的自由时间开发的实验项目，并于 2009 年以 BSD-style 授权（完全开
源）首次公开发布，于 2012 年正式发布。

## 创造
开发者使用编程语言的三大分类（执行速度、简易程度、开发难度）
- 执行速度快、编译速度慢（编译型）： c，c++
- 执行速度较慢、编译速度快（解释型）：java，.net
- 执行速度慢、开发难度小（动态脚本）：Python、PHP

Go 语言在 3 个条件做了平衡：易于开发、快速编译、高效执行.

## 特性
Go 语言被称为 21 世纪的 C 语言，Go 从其他开发语言的借鉴了许多优秀的设计思想，例如从
C 语言借鉴表达式、流程控制、基础数据类型、参数传递、指针等，从 Oberon-2 语言借鉴
的包的导入和声明等，从 Oberon 语言借鉴的面向对象特性中方法的声明语法，从 Limbo 语 言中借鉴的 CSP(通信顺序进程, communicating sequential processes)，从 APL 语言借鉴
的 iota 语法，从 Scheme 语言借鉴的作用域和嵌套函数……

- 静态类型并具有丰富的内置类型: bool、byte、rune、int、float、string、array、
slice、map
- 函数多返回值
- 错误处理机制：使用 defer、panic、recover 定义标准的错误流程
- 语言层并发：使用关键字go将函数以Goroutine方式运行，使用CSP模型作为Goroutine
的通信方式
- 面向对象：使用类型、组合、接口来实现面向对象思想
- 反射
- CGO：用于调用 C 语言实现的模块
- 自动垃圾回收
- 静态编译
- 交叉编译
- 易于部署
- 基于 BSD 协议完全开放

## 应用
Go 语言主要用于服务端开发，其定位是开发大型软件，常用于： 
- 服务器编程：日志处理、数据打包、虚拟机处理、文件系统、分布式系统、数据库代理
等 
- 网络编程：Web 应用、API 应用、下载应用
- 内存数据库
- 云平台
- 机器学习
- 区块链
- ……

使用 Go 开发的项目列表：https://github.com/golang/go/wiki/Projects
- Go 
- docker
- kubernetes
- lantern
- etcd
- Prometheus
- Influxdb
- Consul
- nsq
- beego
- ……

使用 Go 开发的组织：http://go-lang.cat-v.org/organizations-using-go
- 国外：Google、CloudFlare……
- 国内：阿里、腾讯、百度、京东、爱奇艺、小米、今日头条、滴滴、美团、饿了么、360、
七牛、B 站、盛大、搜狗……

## 安装
下载地址：
- https://golang.org/dl/
- https://golang.google.cn/dl/

### Window
### Linux
## 编辑器
- Visual Studio code
- Sublime Text
- Atom
- Eclipse+GoClipse
- LiteIDE
- Goland
- Vim
- Emacs
- ……


# 学习方法
1. 知识点分解
2. 刻意练习（练习不会的、不懂、或者难以理解的知识点）
3. 反馈
    1. 主动反馈
    2. 被动反馈

总结：多写（禁止复制+粘贴），多看，多问


# Go基础
## 第一个go程序

```go
package main 

import "fmt"

func main() {
    fmt.Println("hello world!")
}

执行：
PS E:\go-phase-two\go-course> go run test.go
hello world!

```

解读：
1. package
  go源文件开头必须使用package声明代码所属包，包是go代码分发的最基本单位。若程序需要运行包名必须为main。
2. import
  import用于导入程序以来所有的包。此程序依赖于fmt包
3. func
  func用于定义函数。main程序是程序的入口，若程序需要运行必须声明main函数，main函数无参数也无返回值。
4. fmt.Println
  调用fmt.Println函数将参数信息打印到控制台


编译&运行：
1. go build: 用于编译&链接程序或包
  go build -work -x -o helloworld.exe main.go
2. go run：用于直接运行程序
  go run -work -x main.go
3. go clean：清除编译文件
4. 常用参数：
- -x: 打印编译过程执行的命令，并完成编译或运行
- -n: 只打印编译过程执行命令
- -work：打印编译过程的临时目录
- -o: 指定编译结果文件

## 程序结构
go源文件以package声明开头，说源文件所属的包，截止使用import 导入以来的包，其次为包级别的变量、常量、类型和函数的声明和赋值。函数中可定义局部的变量和常量。

## 基本组成元素
### 标识符
标识符是编程时所使用的名字，用于给变量、常量、函数、类型、接口、包名进行命令，以建立名称和使用之间的关系。
go语言标识符命名规则：
1. 只能由非空字母（Unicode）、数字、下划线（_）组成。
2. 只能以字符或下划线开头。
3. 不能用go语言关键字，25个。
4. 避免使用go语言预定义标识符。
5. 建议使用驼峰式命名
6. 标识符区分大小写

go语言提供一些预先定义的标识符用来表示内置的常量、类型、函数，在自定义标识符时应避免使用：
- 内置常量：true、false、nil、iota
- 内置类型：bool、byte、rune、int、int8、int16、int32、int64、uint、uint8、uint16、uint32、uint64、uintptr、float32、float64、complex64、complex128、string、error
- 内置函数：make、len、cap、new、append、copy、close、delete、complex、real、imag、panic、recover
- 空白标识符：_

### 关键字
关键字用于特定的语法结构，Go语言定义25关键字：
声明：import、package
实体声明和定义：chan、const、func、interface、map、struct、type、var
流程控制：break、case、continue、default、defer、else、fallthrough、for、go、goto、if、range、return、select、switch。

### 字面量
字面量是值的表示方法，常用与对变量/常量进行初始化，主要分为：
- 标识计财处数据类型值的字面量，例如：0,1.1，true，3+4i，'a'，"我爱中国"
- 构造自定义的复合数据类型的类型字面量，例如：type Interval Int
- 用于表示符合数据类型值的复合字面量，用来构造array、slice、map、struct的值，例如：{1, 2, 3}

### 操作符
- 算术运算符：+、-、*、/、%、++、--
- 关系运算符：>、>=、<、<=、==、!=
- 逻辑运算符：&& 、||、!
- 位运算符：&、|、^、<<、>>、&^
- 赋值运算符：=、+=、-+、*=、/=、%=、&=、|=、^=、<<=、>>=
- 其他运算符：&（单目）、*（单目）、.（点）、-（单目）、...、<-

### 分割符
小括号() , 中括号[], 大括号{}， 分号; ，逗号, 

## 声明
声明语句用于定义程序的各种实体对象，主要有：
- 声明变量的var
- 声明常量的const
- 声明函数的func
- 声明类型的type

## 变量
变量是指对一块存储空间定义名称，通过名称对存储空间的内容进行访问或修改，使用var进行变量声明，常用的语法为：
- var 变量名 变量类型 = 指
  定义变量并进行初始化，例如：var name string = "slience"
- var 变量名 变量类型
  定义变量使用零值进行初始化，例如：var age int 
- var 变量名 = 值
  定义变量，变量类型通过值类型进行推导
  例如：var isBoy = true
- var 变量名1，变量名2，...，变量名n 变量类型
  定义多个相同类型的变量并使用零值进行初始化
  例如：var perfix, suffix string
- var 变量名1，变量名2， ...， 变量名n 变量类型 = 值1， 值2， ...， 值n
  定义多个相同类型的变量并使用对应的值进行初始化，例如var prev, next int = 3, 4 
- var 变量名1,变量名2,变量名n... = 值1, 值2, 值n
  定义多个变量并使用对应的值进行初始化，变量的类型使用值类型进行推掉，类型可以不相同，例如:var name, age = "slience", 30 

- 批量定义
  var (
    变量名1 变量类型1 = 值1
    变量名2 变量类型2 = 值2

  )
  定义多个变量并进行初始化， 批量复制中变量类型可以省略
  例如：
  var (
    name string = "slience"
    age int = 30
  )

  初始化表达式可以使用字面量、任何表达式、函数

示例：
```go
package main

import "fmt"

var v0 int
var v1 int = 1
var v2 = 2

var v3, v4 int = 3, 4

var v5, v6 = "5", 6

func main() {
	var (
		v7 int = 2 + 5
		v8 int = v2 + v6
	)
	fmt.Println(v0, v1, v2, v3, v4, v5, v6, v7, v8)
}

PS E:\go-phase-two\go-course> go run test.go
0 1 2 3 4 5 6 7 8
```

## 简短声明
在函数中可以通过简短声明语句声明并初始化变量，可通过简短声明同事声明和初始化多个变量，需要注意操作符左侧的变量至少有一个未定义过。
```go
package main

import "fmt"

func main() {
	n1, n2 := 1, 2
	fmt.Println(n1, n2)
	n2, n3 := 22, 3
	fmt.Println(n2, n3)
}

PS E:\go-phase-two\go-course> go run test.go
1 2
22 3
```

## 赋值
可以通过赋值运算=更新变量的值，go语言支持通过元组赋值同事更新多个变量的值，例如：n1, n2 = 1, 2, 可用于两个变量的交换x, y = y, x 

```go
package main

import "fmt"

func main() {
	n1, n2 := 1, 2
	fmt.Println(n1, n2)
	n1, n2 = n2, n1
	fmt.Println(n1, n2)
}


PS E:\go-phase-two\go-course> go run test.go
1 2
2 1
```

## 常量
常量用于定义不可被修改的值，需要在编译过程中进行计算，只能为基础的数据类型布尔、数值、字符串，使用const进行常量声明，常用语法：
- const 常量名 类型 = 值
  定义常量并进行初始化，例如：const pi float64 = 3.1416926
- const 常量名 = 值
  定义常量，类型通过值类型进行推导，例如：const e = 2.7182818 
- 批量定义
  const (
    常量名1 类型1 = 值1
    常量名2 类型2 = 值2

  )
  定义多个变量并进行初始化，批量复制中变量类型可省略，并且除了第一个常量值外其他常量可同时省略类型和值，标识使用前一个常量的初始化表达式
  例如：
  const (
    name string = "slience"
    age int = 30 
  )
  const (
    name string = "slience"
    desc
  )

  常量之间的运算，类型转换，以及对常量调用函数len、cap、real、imag、complex、unsafe.Sizeof得到的结果依然为常量
```go 
package main

import "fmt"

const c1 int = 1
const c2 = 3 - 1

const c3, c4 int = 3, 4

const c5, c6 = "5", c2 + c4

func main() {
	const (
		c7 int = 7
		c8
	)
	fmt.Println(c1, c2, c3, c4, c5, c6, c7, c8)
}

PS E:\go-phase-two\go-course> go run test.go
1 2 3 4 5 6 7 7
```

## 作用域
作用域指变量可以使用范围。go 语言使用大括号显示的标识作用域范围，大括号内包含一连串的语句，叫做语句块。语句块可以嵌套，语句块内定义的变量不能在语句块外使用。

常见隐式语句块：
- 全语句块
- 包语句块
- 文件语句块
- if、switch、for、select、case语句块

作用域内定义变量只能被声明一次且变量必须使用，否则编译错误。在不同作用域可定义相同的变量，此时局部将覆盖全局。
```go
package main

import "fmt"

func main() {
	outer := 1
	{
		inner := 2
		fmt.Println(outer, inner)
		outer := 3
		fmt.Println(outer, inner)
	}
	fmt.Println(outer)
	//fmt.Println(outer, inter)

}

PS E:\go-phase-two\go-course> go run test.go
1 2
3 2
1
```

## 注释
go支持两种注释方式，行注释和块注释
行注释：以// 开头， 例如： //我是行注释
块注释：以/* 开头， 以*/ 结尾， 例如/* 我是块注释 */

```go
package main

import "fmt"

func main() {
	// 我是行注释
	// 使用fmt包的Println函数将hello world字符串打印到控制台
	fmt.Println("Hello World!")

	/*
		我是块注释
		使用fmt包的Println函数将Hello World字符串打印到控制台

	*/
	fmt.Println("Hello World!")
}
```

## 问题跟踪
最基本的跟踪方式为打印日志：我们可以使用fmt包中提供的Println、Print、Printf函数用于将信息打印到控制台，帮助我们进行问题调试，基本使用方法：

```go
package main

import "fmt"

func main() {
	var name string = "slience"
	var age int = 30
	fmt.Println("name=", name)

	fmt.Printf("name=%v, age=%v\n", name, age)
	fmt.Printf("name=%T, age=%T", name, age)
}


PS E:\go-phase-two\go-course> go run test.go
name= slience
name=slience, age=30
name=string, age=int
```


# 基础数据类型
## 布尔类型
布尔类型用于表示真假，类型名为bool，只有两个值true和false，占用一个字节宽度，零值为false
```go
package main

import "fmt"

var (
	isBody bool = true
	isGirl bool = false
)
func main() {
}
```

常用操作：
1. 逻辑运算：
  1. 与（&&）
    只有左右表达式结果都为true，运算结果才为true。
    ```
    // 与
	  fmt.Println(isBody && isBody)
	  fmt.Println(isBody && isGirl)
	  fmt.Println(isGirl && isBody)
	  fmt.Println(isGirl && isGirl)

    PS E:\go-phase-two\go-course> go run test.go
    true
    false
    false
    false
    ```
  2. 或(||)
    只要左右表达式有一个为true，运算结果就为true
    ```
    fmt.Println(isBody || isBody)
    fmt.Println(isBody || isGirl)
    fmt.Println(isGirl || isBody)
    fmt.Println(isGirl || isGirl)


    PS E:\go-phase-two\go-course> go run test.go
    true
    true
    true
    false
    ```
  3. 非(!)
    右表达式true，运算结果为false；右表达式为false，运算结构为true。
    ```go
    fmt.Println(!isBody)
    fmt.Println(!isGirl)
    ```
2. 关系运算
  1. 等于(==)
  2. 不等于(!=)
    ```go
    //关系运算
    fmt.Println(isBody == isGirl)
    fmt.Println(isBody != isGirl)
    ```
  使用fmt.Printf进行格式化参数输出,占位符：
  - %t
  ```go
  	// 打印
	fmt.Println(isBody, isGirl)
	fmt.Println("isBody=%t, isGirl=%t\n", isBody, isGirl)
  ```

## 数值类型
### 整型
go语言提供了5种符号、5种无符号、1种指针、1种单字节、1种端个Unicode字符(unicode码点)，共13种证书类型，零值均为0 
|类型名 | 字节宽度 | 说明&取值范围 |
-- | -- | -- 
int | 与平台有关，32位系统4字节，64位系统8字节 | 有符号整型 | 
uint | 与平台有关，32位系统4字节，64位系统8字节 | 无符号整型 |
rune | 4字节 | Unicode码点,取值范围同uint32 |
int8 | 1字节 | 用8位标识的有符号整型，取值范围为:[-128, 127 ] |
int16 | 2字节 | 用16位标识有符号整型，取值范围为 [-32768, 32767] |
int32 | 4字节 | 用32位表示的有符号整型，取值范围为...
int64 | 8字节 | 用64位表示有符号整型，取值范围为: ...
uint8 | 1字节 | 用8位表示的无符号整型，取值范围为[0,255]
uint16 | 2字节 | 用16位表示的无符号整型，取值范围为：[0,65535]
uint32 | 4字节 | 用32位表示的无符号整型，取值范围为:[0,.....]
uint64 | 8字节 | 用64位表示的无符号整型，取值范围为：[0,....]
byte | 1字节 | 字节类型，取值范围同uint8 |
uintprt | 与平台有关, 32位系统4字节，64位系统8字节 | 指针值的无符号整型 |

字面量：
- 十进制表示法： 以10 为基数，采用0-9十个数字，逢10进位，例如：10 
- 八进制表示法：以8为基数，采用0-7八个数字，逢8进位，使用0开头表示八进制表示，例如010
- 十六进制表示法：以16位基数采用0-9十个数字和A-F六个字符，逢0X开头表示十六进制，例如：0X10 
```go
func main() {
	fmt.Println(10, 010, 0X10)
}

PS E:\go-phase-two\go-course> go run test.go
10 8 16
```

常用操作：
```go 
var n1, n2, n3 int = 1, 8, -2
var n4 uint = 2
```

1. 算术运算符： + 、 - 、 * 、/、%、++、-- 
  注意：针对/除数不能为0，且结果依然为整数
  ```go
  package main
  
  import "fmt"
  
  var n1, n2, n3 int = 1, 8, -2
  var n4 uint = 2
  
  func main() {
  	fmt.Println(n1 + n2)
  	fmt.Println(n1 - n2)
  	fmt.Println(n1 * n2)
  	fmt.Println(n1 / n2)
  	fmt.Println(n1 % n2)
  	n1++
  	n2--
  	fmt.Println(n1, n2)
  }

  ```

2. 关系运算符： > 、 >= 、<=、<、==、!=
```go
package main

import "fmt"

var n1, n2, n3 int = 1, 8, -2
var n4 uint = 2

func main() {
	fmt.Println(n1 > n2)
	fmt.Println(n1 >= n2)
	fmt.Println(n1 < n2)
	fmt.Println(n1 <= n2)
	fmt.Println(n1 == n2)
	fmt.Println(n1 != n2)
}

PS E:\go-phase-two\go-course> go run test.go
false
false
true
true
false
true
```
3. 位运算符：&、 | 、^、<<、>>/ &^
对于负整数在计算机中使用补码进行表示，对应正整数二进制数表示取反+1
针对左、右移的有右操作符必须有无符号整数
```go
//位运算
// 2 => 0010 
// 7 => 0111
// -2 => 1110
package main

import "fmt"
// 1 => 0001
// 8 => 1000
// -2 => 1110
// 2 => 0010


var n1, n2, n3 int = 1, 8, -2
var n4 uint = 2

func main() {
	fmt.Println(n1 & n2)
	fmt.Println(n1 | n2)
	fmt.Println(n1 ^ n2)
	fmt.Println(n1 &^ n2)
	fmt.Println(n2 << n4)
	fmt.Println(n2 >> n4)
	fmt.Println(n3 << n4)
	fmt.Println(n3 << n4)

}

```
4. 赋值运算符：= 、+= 、 -= 、*= 、/= 、 %=、&=、|=、^=、<<=、>>=
