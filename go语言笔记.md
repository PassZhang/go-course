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

```
// 赋值运算
package main

import "fmt"

var n1, n2, n3 int = 1, 8, -2
var n4 uint = 2

func main() {
	n1 += n2
	n2 -= n1
	n4 <<= n4

	fmt.Println(n1, n2, n4)
}

PS E:\go-phase-two\go-course> go run test.go
9 -1 8
```

6. 类型转换：
   Go不会自动对数据类型进行转换，因此左、右操作数类型必须一致或某个字面量，可通过类型名（数据）的语法将数据转为对应类型。需要注意值截断和值溢出问题

```
package main

import "fmt"

var n1, n2, n3 int = 1, 8, -2
var n4 uint = 2

func main() {
	fmt.Printf("%T %T %T %T %T %T\n", n1, uint(n1), n4, int(n4), uint8(n4), int64(n4))
}

PS E:\go-phase-two\go-course> go run test.go
int uint uint int uint8 int64
```

使用fmt.Printf进行格式化参数输出，占位符：
- %b: 二进制
- %c： 字符
- %d：十进制
  - %+d ： 表示对正整数带+符号
  - %nd ： 表示最小占位N个宽度且右对齐
  - %-nd : 表示最小占位N个宽度按且左对齐
  - %0nd ：表示最小占位n个宽度且右对齐，空字符使用0填充
- %o： 八进制，%#o带0的前缀
- %x、%X：十六进制，%#x(%#X)带0x(0X)的前缀
- %U：Unicode码点，%#U带字符的Unicode码点
- %q：带单引号的字符

```
package main

import "fmt"

var (
	n5 int  = 10
	n6 int  = -10
	c0 byte = 'a'
	u0 rune = '牛'
)

func main() {
	fmt.Printf("%d %b %o %x %X %#o %#x %#X\n", n5, n5, n5, n5, n5, n5, n5, n5)
	fmt.Printf("%d|%+d|%10d|%-10d|%010d|%+-10d|%+010d|\n", n5, n5, n5, n5, n5, n5, n5)
	fmt.Printf("%d|%d|%10d|%-10d|%010d|%+-10d|%+010d|\n", n5, n5, n5, n5, n5, n5, n5)
	fmt.Printf("%c %c %q %q\n", c0, u0, c0, u0)
	fmt.Printf("%U %#U\n", u0, u0)
}

PS E:\go-phase-two\go-course> go run test.go
10 1010 12 a A 012 0xa 0XA
10|+10|        10|10        |0000000010|+10       |+000000010|
10|10|        10|10        |0000000010|+10       |+000000010|
a 牛 'a' '牛'
U+725B U+725B '牛'
```

常用包：
- math
- math/rand


### 浮点型

浮点数用于表示带小数的数字，Go提供float32和float64两种浮点类型

字面量：
- 十进制表示法：3.1415926
- 科学计数法：1e-5

常用操作：
1. 算术运算符：+、-、*、/、++、--
  注意：针对/除数不能为 0


```go
package main

import "fmt"

var (
	f1 float32 = 3.1415926
	f2 float32 = 3E-3
	f3 float64 = 3.1E10
)

func main() {
	// 算术运算
	fmt.Println(f1 + f2)
	fmt.Println(f1 - f2)
	fmt.Println(f1 * f2)
	fmt.Println(f1 / f2)
	f1++
	f2--
	fmt.Println(f1, f2)
}

PS E:\go-phase-two\go-course> go run test.go
3.1445925
3.1385925
0.009424778
1047.1975
4.1415925 -0.997
```

2. 关系运算符：>、>=、<、<=
浮点型不能进行==或!=比较，可选择使用两个浮点数的差在一定区间内则认为相等

```go
package main

import "fmt"

var (
	f1 float32 = 3.1415926
	f2 float32 = 3E-3
	f3 float64 = 3.1E10
)

func main() {
	fmt.Println(f1 > f2)
	fmt.Println(f1 >= f2)
	fmt.Println(f1 < f2)
	fmt.Println(f1 <= f2)
}

PS E:\go-phase-two\go-course> go run test.go
true
true
false
false
```

3. 赋值运算符：=、+=、-=、*=、/=

```go
package main

import "fmt"

var (
	f1 float32 = 3.1415926
	f2 float32 = 3E-3
	f3 float64 = 3.1E10
)

func main() {
	f1 += f1
	fmt.Println(f1)
}

PS E:\go-phase-two\go-course> go run test.go
6.283185
```


4. 类型转换

  Go不会自动对数据类型进行转换，因此左、右操作数类型必须一致或某个字面量，可通过类型名(数据)的语法将数据转为对应类型。需要注意值截断和值溢出问题。

```go
package main

import "fmt"

var (
	f1 float32 = 3.1415926
	f2 float32 = 3E-3
	f3 float64 = 3.1E10
)

func main() {
	fmt.Println("%T %T\n", f1, float64(f1))

}
```

使用fmt.Printf进行格式化参数输出,占位符：
- %f、%F：十进制表示法
  - %n.mf表示最小占n个宽度并且保留m位小数
- %e，%E：科学计数法
- %g、%G：自动选择最紧凑的表示方法%e(%E)或%f(%F)

```go
package main

import "fmt"

var (
	f1 float32 = 3.1415926
	f2 float32 = 3E-3
	f3 float64 = 3.1E10
)

func main() {
	fmt.Printf("%f,%f,%F,%F\n", f1, f3, f1, f3)
	fmt.Printf("%e,%e,%E,%E\n", f1, f3, f1, f3)
	fmt.Printf("%g,%g,%G,%G\n", f1, f3, f1, f3)
	fmt.Printf("%5.2f,%f5.2f\n", f1, f3)

}

PS E:\go-phase-two\go-course> go run test.go
3.141593,31000000000.000000,3.141593,31000000000.000000
3.141593e+00,3.100000e+10,3.141593E+00,3.100000E+10
3.1415925,3.1e+10,3.1415925,3.1E+10
 3.14,31000000000.0000005.2f
```

常用包
- math
- math/rand

### 复数型

go提供complex64和complex128两种附中类型，针对complex64复数的实部和虚部均使用float32，针对complex128的实部和虚部均使用float64。

字面量：
- 十进制表示法:1 + 2i， , i*i = -1, 1为实部， 2为虚部

常用函数：
- complex： 工厂函数，通过两个参数创建一个复数
- real： 用于获取复数的实部
- imag： 用于获取复数的虚部
  
```go
package main

import "fmt"

var (
	c1 complex64 = 1 + 2i
	c2 complex64 = complex(3, 4)
)

func main() {
	fmt.Println(c1 + c2)
	fmt.Println(real(c1), imag(c1))
}

PS E:\go-phase-two\go-course> go run test.go
(4+6i)
1 2
```

常用包：
- math/cmplx

## 字符串类型

go语言内置了字符串类型，使用string表示

字面量：
- 可解析字符串，通过双引号（"）来创建，不能包含多行，支持特殊字符转义序列
- 原生字符串：通过反引号（`）来创建，可包含多行，不支持特殊字符转义序列

特殊字符：
- \\：反斜线
- \'：单引号
- \"：双引号
- \a：响铃
- \b：退格
- \f：换页
- \n：换行
- \r：回车
- \t：制表符
- \v：垂直制表符
- \ooo：3 个 8 位数字给定的八进制码点的 Unicode 字符（不能超过\377）  
- \uhhhh：4 个 16 位数字给定的十六进制码点的 Unicode 字符
- \Uhhhhhhhh：8 个 32 位数字给定的十六进制码点的 Unicode 字符
- \xhh：2 个 8 位数字给定的十六进制码点的 Unicode 字符

```go
package main

import "fmt"

var (
	s1 string = "abc\nbcd"
	s2 string = `abc\nbcd`
	s3 string = `abc
	bcd
	`
	s4 string = "avcsdjflsdjlgjdslkgjsldkfjlksd"
	s5 string = "我爱中国"
)

func main() {
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(s5)
}

PS E:\go-phase-two\go-course> go run test.go
abc
bcd
abc\nbcd
abc
        bcd

avcsdjflsdjlgjdslkgjsldkfjlksd
我爱中国
```

常用操作
a) 字符串连接：+
b) 关系运算符：>、>=、<、<=、==、!=
c) 赋值运算符：+=
d) 索引：s[index]，针对只包含 ascii 字符的字符串
e) 切片：s[start:end] ，针对只包含 ascii 字符的字符串

```go
package main

import "fmt"

var (
	s1 string = "abc\nbcd"
	s2 string = `abc\nbcd`
	s3 string = `abc
	bcd
	`
	s4 string = "avcsdjflsdjlgjdslkgjsldkfjlksd"
	s5 string = "我爱中国"
)

func main() {
	fmt.Println(s4 + "AAAAAAAAAAAAAAAAAAAAAAAA")
	s5 += "--来自Silence"
	fmt.Println(s5)

	fmt.Println(s1 == s2)
	fmt.Println(s1 != s3)
	fmt.Println("abc" > "abe")
	fmt.Println("abc" < "abe")
	fmt.Println("abc" >= "abe")
	fmt.Println("abc" <= "abe")
	fmt.Println(s4[0])
	fmt.Println(s4[:], s4[3:], s4[:8], s4[3:8])
}

PS E:\go-phase-two\go-course> go run test.go
avcsdjflsdjlgjdslkgjsldkfjlksdAAAAAAAAAAAAAAAAAAAAAAAA
我爱中国--来自Silence
false
true
false
true
false
true
97
avcsdjflsdjlgjdslkgjsldkfjlksd sdjflsdjlgjdslkgjsldkfjlksd avcsdjfl sdjfl
```

常用函数
1. len: 获取字符串长度(针对只包含ascii字符的字符串)
2. string： 将byte或者rune数组转换为字符串

使用fmt.Printf进行格式化参数输出，占位符
- %s

```go
package main

import "fmt"

var (
	s1 string = "abc\nbcd"
	s2 string = `abc\nbcd`
	s3 string = `abc
	bcd
	`
	s4 string = "avcsdjflsdjlgjdslkgjsldkfjlksd"
	s5 string = "我爱中国"
)

func main() {
	fmt.Printf("%s len %d\n", s1, len(s1))
	fmt.Printf("%b len %d\n", s5, len(s5))
}


PS E:\go-phase-two\go-course> go run test.go
abc
bcd len 7
我爱中国 len 12
```


常用包：
- fmt
- strings
- strconv
- unicode
- unicode 
- bytes
- regex


## 枚举类型

常使用 iota 生成器用于初始化一系列相同规则的常量，批量声明常量的第一个常量使用
iota 进行赋值，此时 iota 被重置为 0，其他常量省略类型和赋值，在每初始化一个常量则
加 1。

```go
package main

import "fmt"

const (
	c1 int = iota
	c2
	c3 int = iota
	c4
)

func main() {
	const (
		c5 = iota
		c6
		c7 = 7
		c8 = iota
		c9
		c10 = iota
		c11
		c12
	)
	fmt.Println(c1, c2, c3, c4, c5, c6, c7, c8, c9, c10, c11, c12)
}


PS E:\go-phase-two\go-course> go run test.go
0 1 2 3 0 1 7 3 4 5 6 7
```


## 指针类型

每个变量在内存中都有对应存储位置（内存地址），可以通过&运算符获取。指针是用来存储变量地址的变量。

1. 声明 

指针声明需要指定存储地址中对应数据的类型，并使用*作为类型前缀。指针变量声明后会被初始化为nil，表示空指针。

```go
package main

import "fmt"

func main() {
	var pointter01 *int
	var pointter02 *float64
	var pointter03 *string

	fmt.Printf("%T, %T, %T\n", pointter01, pointter02, pointter03)
	fmt.Println(pointter01, pointter02, pointter03)
	fmt.Printf("%d, %d, %d\n", pointter01, pointter02, pointter03)

}

PS E:\go-phase-two\go-course> go run test.go
*int, *float64, *string
<nil> <nil> <nil>
0, 0, 0
```


1. 初始化 
   1. 使用&运算符+变量初始化：&运算获取变量的存储位置来初始化指针变
   2. 使用new函数初始化：new函数根据数据类型申请内存空间并使用零值填充，并返回申请空间地址
```go
package main

import "fmt"

var (
	age   int     = 30
	heigh float64 = 1.68
	motto string  = "少年经不得顺境， 中年经不得闲境，晚年经不得逆境"
)

func main() {

	// 打印变量初始化
	var pointer01, pointer02, pointer03 = &age, &heigh, &motto
	pointer04, pointer05, pointer06 := &age, &heigh, &motto
	pointer07, pointer08, pointer09 := new(int), new(float64), new(string)

	// 打印变量地址
	fmt.Println(&age, &heigh, &motto)

	// 打印指针变量
	fmt.Println(pointer01, pointer02, pointer03)
	fmt.Println(pointer04, pointer05, pointer06)
	fmt.Println(pointer07, pointer08, pointer09)

	fmt.Println(age, heigh, motto)
	fmt.Println(*pointer01, *pointer02, *pointer03)
	fmt.Printf("%v, %v, %q\n", *pointer07, *pointer08, *pointer09)
}

PS E:\go-phase-two\go-course> go run test.go
0x565210 0x565218 0x5729d0
0x565210 0x565218 0x5729d0
0x565210 0x565218 0x5729d0
0xc0000100a0 0xc0000100a8 0xc0000481f0
30 1.68 少年经不得顺境， 中年经不得闲境，晚年经不得逆境
30 1.68 少年经不得顺境， 中年经不得闲境，晚年经不得逆境
0, 0, ""
```


3. 操作
   可通过 *运算符 + 指针变量名来访问和修改对应存储位置的值

```go
package main

import "fmt"

var (
	age   int     = 30
	heigh float64 = 1.68
	motto string  = "少年经不得顺境， 中年经不得闲境，晚年经不得逆境"
)

func main() {

	// 打印变量初始化
	var pointer01, pointer02, pointer03 = &age, &heigh, &motto
	pointer04, pointer05, pointer06 := &age, &heigh, &motto
	pointer07, pointer08, pointer09 := new(int), new(float64), new(string)

	// 打印变量地址
	fmt.Println(&age, &heigh, &motto)

	// 打印指针变量
	fmt.Println(pointer01, pointer02, pointer03)
	fmt.Println(pointer04, pointer05, pointer06)
	fmt.Println(pointer07, pointer08, pointer09)

	fmt.Println(age, heigh, motto)
	fmt.Println(*pointer01, *pointer02, *pointer03)
	fmt.Printf("%v, %v, %q\n", *pointer07, *pointer08, *pointer09)

	fmt.Println(age, heigh, motto)                  //打印变量值
	fmt.Println(*pointer01, *pointer02, *pointer03) //通过指针变量访问位置存储的值

	*pointer01 += 1
	*pointer02 = 1.70
	*pointer03 += "--曾国藩"

	fmt.Println(age, heigh, motto)                  //打印变量值
	fmt.Println(*pointer01, *pointer02, *pointer03) //通过指针变量访问位置存储的值

	fmt.Println(&age, &heigh, &motto)               // 打印变量地址
	fmt.Println(*pointer01, *pointer02, *pointer03) // 打印指针变量

	// 与赋值新变量对比
	age2, heigh2, motto2 := age, heigh, motto

	// 修改新变量值
	age2 += 1
	heigh2 = 1.72
	motto2 += "家书"

	// 打印变量
	fmt.Println(age, heigh, motto)    //打印变量值
	fmt.Println(age2, heigh2, motto2) //打印变量值

	fmt.Println(&age, &heigh, &motto)    //打印变量值
	fmt.Println(&age2, &heigh2, &motto2) //打印变量值
}


PS E:\go-phase-two\go-course> go run test.go
0x566210 0x566218 0x5739d0
0x566210 0x566218 0x5739d0
0x566210 0x566218 0x5739d0
0xc0000100a0 0xc0000100a8 0xc0000481f0
30 1.68 少年经不得顺境， 中年经不得闲境，晚年经不得逆境
30 1.68 少年经不得顺境， 中年经不得闲境，晚年经不得逆境
0, 0, ""
30 1.68 少年经不得顺境， 中年经不得闲境，晚年经不得逆境
30 1.68 少年经不得顺境， 中年经不得闲境，晚年经不得逆境
31 1.7 少年经不得顺境， 中年经不得闲境，晚年经不得逆境--曾国藩
31 1.7 少年经不得顺境， 中年经不得闲境，晚年经不得逆境--曾国藩
0x566210 0x566218 0x5739d0
31 1.7 少年经不得顺境， 中年经不得闲境，晚年经不得逆境--曾国藩
31 1.7 少年经不得顺境， 中年经不得闲境，晚年经不得逆境--曾国藩
32 1.72 少年经不得顺境， 中年经不得闲境，晚年经不得逆境--曾国藩家书
0x566210 0x566218 0x5739d0
0xc000010148 0xc000010150 0xc000048270
```


4. 指针的指针

  用来存储指针变量地址的变量叫做指针的指针

```go
package main

import "fmt"

var (
	age   int     = 30
	heigh float64 = 1.68
	motto string  = "少年经不得顺境， 中年经不得闲境，晚年经不得逆境"
)

func main() {

	// 打印变量初始化
	var pointer01, pointer02, pointer03 = &age, &heigh, &motto
	pointer04, pointer05, pointer06 := &age, &heigh, &motto
	pointer07, pointer08, pointer09 := new(int), new(float64), new(string)

	// 打印变量地址
	fmt.Println(&age, &heigh, &motto)

	// 打印指针变量
	fmt.Println(pointer01, pointer02, pointer03)
	fmt.Println(pointer04, pointer05, pointer06)
	fmt.Println(pointer07, pointer08, pointer09)

	fmt.Println(age, heigh, motto)
	fmt.Println(*pointer01, *pointer02, *pointer03)
	fmt.Printf("%v, %v, %q\n", *pointer07, *pointer08, *pointer09)

	fmt.Println(age, heigh, motto)                  //打印变量值
	fmt.Println(*pointer01, *pointer02, *pointer03) //通过指针变量访问位置存储的值

	*pointer01 += 1
	*pointer02 = 1.70
	*pointer03 += "--曾国藩"

	fmt.Println(age, heigh, motto)                  //打印变量值
	fmt.Println(*pointer01, *pointer02, *pointer03) //通过指针变量访问位置存储的值

	fmt.Println(&age, &heigh, &motto)               // 打印变量地址
	fmt.Println(*pointer01, *pointer02, *pointer03) // 打印指针变量

	// 与赋值新变量对比
	age2, heigh2, motto2 := age, heigh, motto

	// 修改新变量值
	age2 += 1
	heigh2 = 1.72
	motto2 += "家书"

	// 打印变量
	fmt.Println(age, heigh, motto)    //打印变量值
	fmt.Println(age2, heigh2, motto2) //打印变量值

	fmt.Println(&age, &heigh, &motto)    //打印变量值
	fmt.Println(&age2, &heigh2, &motto2) //打印变量值

	// 定义声明指针的指针
	var ppointer01 **int
	var ppointer02 **float64 = &pointer02
	ppointer03 := &pointer03

	fmt.Println(ppointer01, ppointer02, ppointer03)

	ppointer01 = &pointer01
	fmt.Println(ppointer01, ppointer02, ppointer03)

	// 通过指针的指针访问变量地址和变量值
	fmt.Println(*ppointer01, *ppointer02, *ppointer03)
	fmt.Println(**ppointer01, **ppointer02, **ppointer03)

	// 通过指针的指针修改和变量值
	**ppointer01 += 1
	**ppointer02 = 1.72
	**ppointer03 += "家属"

	fmt.Println(**ppointer01, **ppointer02, **ppointer03)
	fmt.Println(*pointer01, *pointer02, *pointer03)
	fmt.Println(age, heigh, motto)
}


PS E:\go-phase-two\go-course> go run test.go
0x567210 0x567218 0x5749d0
0x567210 0x567218 0x5749d0
0x567210 0x567218 0x5749d0
0xc000074068 0xc000074080 0xc00005c1e0
30 1.68 少年经不得顺境， 中年经不得闲境，晚年经不得逆境
30 1.68 少年经不得顺境， 中年经不得闲境，晚年经不得逆境
0, 0, ""
30 1.68 少年经不得顺境， 中年经不得闲境，晚年经不得逆境
30 1.68 少年经不得顺境， 中年经不得闲境，晚年经不得逆境
31 1.7 少年经不得顺境， 中年经不得闲境，晚年经不得逆境--曾国藩
31 1.7 少年经不得顺境， 中年经不得闲境，晚年经不得逆境--曾国藩
0x567210 0x567218 0x5749d0
31 1.7 少年经不得顺境， 中年经不得闲境，晚年经不得逆境--曾国藩
31 1.7 少年经不得顺境， 中年经不得闲境，晚年经不得逆境--曾国藩
32 1.72 少年经不得顺境， 中年经不得闲境，晚年经不得逆境--曾国藩家书
0x567210 0x567218 0x5749d0
0xc000074120 0xc000074128 0xc00005c260
<nil> 0xc00009e020 0xc00009e028
0xc00009e018 0xc00009e020 0xc00009e028
0x567210 0x567218 0x5749d0
31 1.7 少年经不得顺境， 中年经不得闲境，晚年经不得逆境--曾国藩
32 1.72 少年经不得顺境， 中年经不得闲境，晚年经不得逆境--曾国藩家属
32 1.72 少年经不得顺境， 中年经不得闲境，晚年经不得逆境--曾国藩家属
32 1.72 少年经不得顺境， 中年经不得闲境，晚年经不得逆境--曾国藩家属
```


# 流程控制

我们经常需要代码在满足一定条件时进行执行，或者需要重复执行代码多次，此时需要选择
条件语句(if-else if- else)或选择语句(switch case)及循环语句(for)。

## 条件语句

段子：老婆给当程序员的老公打电话：下班顺路买十个包子，如果看到卖西瓜的，买一个。
当晚老公手捧一个包子进了家门…老婆怒道：你怎么只买一个包子？！老公甚恐，喃喃道：
因为我真看到卖西瓜的了。
老婆：买十个包子，如果有卖西瓜的，买一个西瓜
老公：如果有卖西瓜的，买一个包子，否则买十个包子

### if语句

```go
package main

import "fmt"

func main() {
	// 老婆
	fmt.Println("老婆：")
	fmt.Println("有卖西瓜的：")
	has_watermelon := true
	fmt.Println("买10个包子")
	if has_watermelon {
		fmt.Println("买1个西瓜")
	}

	fmt.Println("没有卖西瓜的：")
	has_watermelon = false
	fmt.Println("买10个包子")
	if has_watermelon {
		fmt.Println("有卖西瓜的，买一个西瓜")
	}
}


PS E:\go-phase-two\go-course> go run test.go
老婆：
有卖西瓜的：
买10个包子
买1个西瓜
没有卖西瓜的：
买10个包子
```

当 if 表达式的结果为 true 则执行语句块内代码。


### if-else 语句 

```go
package main

import "fmt"

func main() {
	// // 老婆
	// fmt.Println("老婆：")
	// fmt.Println("有卖西瓜的：")
	// has_watermelon := true
	// fmt.Println("买10个包子")
	// if has_watermelon {
	// 	fmt.Println("买1个西瓜")
	// }

	// fmt.Println("没有卖西瓜的：")
	// has_watermelon = false
	// fmt.Println("买10个包子")
	// if has_watermelon {
	// 	fmt.Println("有卖西瓜的，买一个西瓜")
	// }

	// 老公
	fmt.Println("老公：")

	fmt.Println("有卖西瓜的：")
	has_watermelon := true
	if has_watermelon {
		fmt.Println("买1个包子")
	} else {
		fmt.Println("买10个包子")
	}

	fmt.Println("没有卖西瓜的")
	has_watermelon = false
	if has_watermelon {
		fmt.Println("买1个包子")
	} else {
		fmt.Println("买10个包子")
	}
}


PS E:\go-phase-two\go-course> go run test.go
老公：
有卖西瓜的：
买1个包子
没有卖西瓜的
买10个包子
```


当 if 表达式结果为 true，则执行 if 语句块内代码，否则执行 else 语句块内代码。


### else if

成绩评优: [90, 100]=>优秀(A), [80, 90) => 良好(B), [60, 80) => 及格(C), [0, 60) => 
不及格(D)
使用 if-else 语句实现

```go
package main

import "fmt"

func main() {
	score := 90

	if score >= 90 {
		fmt.Println("优秀")
	} else {
		if score >= 80 {
			fmt.Println("良好")
		} else {
			if score >= 60 {
				fmt.Println("及格")
			} else {
				fmt.Println("不及格")
			}
		}
	}
}



```


使用 if-else if-else 语句实现

```go
package main

import "fmt"

func main() {
	score := 90

	if score >= 90 {
		fmt.Println("优秀")
	} else if score >= 80 {
		fmt.Println("良好")
	} else if score >= 60 {
		fmt.Println("及格")
	} else {
		fmt.Println("不及格")
	}
}

```

当 if 表达式结果为 true，则执行 if 语句块内代码，否则依次从上到下判断 else if 表达
式结果，若结果为 true 则执行对应语句块内代码并退出 if-else if-else 语句，若 if 和
else if 表达式均为 false，则执行 else 语句块内代码


### 初始化子语句

```go

package main

import "fmt"

func main() {
	if flag := true; flag {
		fmt.Println("flag:", flag)
	}
	if flag := false; flag {
		fmt.Println("flag:", flag)
	} else {
		fmt.Println("flag else:", flag)
	}
}


PS E:\go-phase-two\go-course> go run test.go
flag: true
flag else: false
```

### 总结 
对于条件语句必须有 if 语句，可以有 0 个或多个 else if 语句，最多有 1 个 else 语句，语句从上到下执行，执行第一个条件表达式为 true 的语句块并退出，否则执行 else 语句块退出.


## 选择语句

根据输入条件的不同选择不同的语句块进行执行(多分支)

### switch-case 单值

```go
package main

import "fmt"

func main() {
	// 单值
	operate := "add"

	switch operate {
	case "add":
		fmt.Println("添加")
	case "delete":
		fmt.Println("删除")
	case "modify":
		fmt.Println("修改")
	case "query":
		fmt.Println("查询")
	default:
		fmt.Println("操作错误")
	}

	input := "y"

	switch input {
	case "yes", "y":
		fmt.Println("确认")
	case "no", "n":
		fmt.Println("取消")
	default:
		fmt.Println("操作错误")
	}

}

PS E:\go-phase-two\go-course> go run test.go
添加
确认
```

swtich 语句后面跟接变量，此时选择与变量相同的 case 语句块执行并退出，若所有 case 均不相同则执行 default 语句块，case 语句中可包含多个不同的值进行匹配。


### switch-case 表达式

```go
package main

import "fmt"

func main() {
	// 表达式
	score := 80

	switch {
	case score >= 90:
		fmt.Println("youxiu")
	case score >= 80:
		fmt.Println("lianghao")
	case score >= 60:
		fmt.Println("jige")
	default:
		fmt.Println("bujige")
	}
}



```

switch 后不跟接变量，此时自上到下选择第一个表达式为 true 的 case 语句块执行并退出，若所有case表达式均为false，则执行default语句块。


### 初始化子语句

可以在switch语句中初始化语句块内的局部变量，多个语句之间使用逗号（;）分割, 注意初始化表达式后面的逗号(;) 不能省略。

```go
package main

import "fmt"

func main() {

	// 初始化单值
	switch operate := "add"; operate {
	case "add":
		fmt.Println("添加")
	case "delete":
		fmt.Println("删除")
	case "modify":
		fmt.Println("修改")
	case "query":
		fmt.Println("查询")
	default:
		fmt.Println("操作错误")
	}

	// 初始化表达式
	switch score := 80; {
	case score >= 90:
		fmt.Println("youxiu")
	case score >= 80:
		fmt.Println("lianghao")
	case score >= 60:
		fmt.Println("jige")
	default:
		fmt.Println("bujige")
	}
}


```

### fallthrough

switch-case 默认执行case 语句后退出，需要继续执行下一个case 语句块，可以在case语句块中使用fullthrough 进行声明。

```go
package main

import "fmt"

func main() {

	// 初始化表达式
	switch score := 80; {
	case score >= 90:
		fmt.Println("youxiu")
	case score >= 80:
		fmt.Println("lianghao")
		fallthrough
	case score >= 60:
		fmt.Println("jige")
	default:
		fmt.Println("bujige")
	}
}


PS E:\go-phase-two\go-course> go run test.go
lianghao
jige

```

### 总结

对于选择语句可以有 0 个或多个 case 语句，最多有 1 个 default 语句，选择条件为 true的 case 语句块开始执行并退出，若所有条件为 false，则执行 default 语句块并退出。可以通过 fallthrough 修改执行退出行为，继续执行下一条的 case 或 default 语句块。



## 循环语句

问题：计算 1-100 所有整数的和

### for

```go
package main

import "fmt"

func main() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)
}


PS E:\go-phase-two\go-course> go run test.go
5050
```

for 语句后有3个子语句分别为： 初始化子语句，条件子语句和后置子语句
执行顺序为：
1. 初始化子语句
2. 条件子语句 
3. 语句块 
4. 后置子语句
5. b -> c -> d 
6. ...
7. 直到条件子语句为false 结束循环


### break 与 continue语句

```go
package main

import "fmt"

func main() {
	//continue
	for i := 1; i <= 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}
	fmt.Println("-----------------------")
	//break
	for i := 1; i <= 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println(i)
	}
}


PS E:\go-phase-two\go-course> go run test.go
1
2
4
5
-----------------------
1
2
```


break 用于跳出循环，当条件满足则结束循环
continue 用于跳过循环，当条件满足这跳过本次循环进行后置或条件子语句执行


### 类while 

for 子语句可以只保留条件子语句，此时类似于其他语言中的 while 循环.

```go
package main

import "fmt"

func main() {
	sum := 0
	i := 1
	for i <= 100 {
		sum += i
		i++
	}
	fmt.Println(sum)
}


PS E:\go-phase-two\go-course> go run test.go
5050
```


### 无限循环

for 子语句全部省略，则为无限循环（死循环），s常与break结合使用

```go
package main

import "fmt"

func main() {
	sum := 0
	i := 1
	for {
		if i > 100 {
			break
		}
		sum += i
		i++
	}
	fmt.Println(sum)
}


PS E:\go-phase-two\go-course> go run test.go
5050
```


可以使用ctrl+z 中止程序运行

```go
package main

import "fmt"

func main() {
	i := 0
	for {
		i++
		fmt.Println(i)
	}
}



```


### for-range

用于遍历可迭代对象的每个元素，例如字符串，数组，切片，映射，通道

```go
package main

import "fmt"

func main() {

	text := "我爱中国"
	for i, e := range text {
		fmt.Printf("%d: %c\n", i, e)

	}

	for _, e := range text {
		fmt.Printf("%c\n", e)
	}
}


PS E:\go-phase-two\go-course> go run test.go
0: 我
3: 爱
6: 中
9: 国
我
爱
中
国

```

针对包含Unicode字符的字符串遍历是需要使用for-range，range返回两个元素分别为字节索引和rune字符，可通过空白标识符需要接受的变量。

## label 与 goto 

可以通过 goto 语句任意跳转到当前函数指定的 label 位置

```go
package main

import "fmt"

func main() {

	sum := 0
	i := 1

START: // 定义START
	if i > 100 {
		goto END // 跳转到END标签

	} else {
		sum += i
		i++
		goto START //跳转到START标签
	}
END: // 定义END标签
	fmt.Println(sum)
}


PS E:\go-phase-two\go-course> go run test.go
5050
```


break 和 continue 后也可以指定 label 用于指定跳出或跳过指定 label 同层级的循环.


```go
package main

import "fmt"

func main() {

CI:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j == 3 {
				continue CI // 跳过外层循环
			}
			fmt.Println(i, j)
		}
	}

BI:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j == 3 {
				break BI // 跳出外层循环
			}
			fmt.Println(i, j)
		}
	}
}


PS E:\go-phase-two\go-course> go run test.go
0 0
0 1
0 2
1 0
1 1
1 2
2 0
2 1
2 2
3 0
3 1
3 2
4 0
4 1
4 2
0 0
0 1
0 2
```

## 练习 
### 乘法口诀

```go
package main

import "fmt"

func main() {
	for x := 1; x <= 9; x++ {
		for y := 1; y <= x; y++ {
			fmt.Printf("%d * %d = %2d ", x, y, x*y)
		}
		fmt.Println()
	}
}


PS E:\go-phase-two\go-course> go run test.go
1 * 1 =  1
2 * 1 =  2 2 * 2 =  4
3 * 1 =  3 3 * 2 =  6 3 * 3 =  9
4 * 1 =  4 4 * 2 =  8 4 * 3 = 12 4 * 4 = 16
5 * 1 =  5 5 * 2 = 10 5 * 3 = 15 5 * 4 = 20 5 * 5 = 25
6 * 1 =  6 6 * 2 = 12 6 * 3 = 18 6 * 4 = 24 6 * 5 = 30 6 * 6 = 36
7 * 1 =  7 7 * 2 = 14 7 * 3 = 21 7 * 4 = 28 7 * 5 = 35 7 * 6 = 42 7 * 7 = 49
8 * 1 =  8 8 * 2 = 16 8 * 3 = 24 8 * 4 = 32 8 * 5 = 40 8 * 6 = 48 8 * 7 = 56 8 * 8 = 64
9 * 1 =  9 9 * 2 = 18 9 * 3 = 27 9 * 4 = 36 9 * 5 = 45 9 * 6 = 54 9 * 7 = 63 9 * 8 = 72 9 * 9 = 81
```



### 猜数游戏

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

const max_Guess = 10

func main() {
	var answer int
	rand.Seed(time.Now().Unix())
	result := rand.Intn(100)

	for i := 1; i <= max_Guess; i++ {
		fmt.Print("请输入您要猜的数字：")

		if _, err := fmt.Scanln(&answer); err != nil {
			fmt.Print("请重新您要猜的数字：")
			continue
		}

		if answer > result {
			fmt.Printf("太大了，您还有%d次机会！\n", max_Guess-i)

		} else if answer < result {
			fmt.Printf("太小了，您还有%d次机会！\n", max_Guess-i)
		} else {
			fmt.Printf("猜对了，您一共猜了%d次！\n", i)
			break
		}
	}
}



```


# 复合数据类型
## 数组

数组是具有相同数据类型的数据组成一组长度固定的序列，每一个数据叫做数组的元素，数组的长度息是非负整数的常量，长度也是类型的一部分。

### 声明 

数组声明需要指定组成元素的类型以及存储元素的数量(长度)。在数组声明之后，其长度不可修改，数组的每个元素会根据对应的类型的零值进行初始化。

```go
package main

import (
	"fmt"
)

func main() {
	var names [10]string
	var scores [10]int

	fmt.Printf("%T, %T\n", names, scores)
	fmt.Printf("%q\n", names)
	fmt.Println(scores)
}


PS E:\go-phase-two\go-course> go run test.go
[10]string, [10]int
["" "" "" "" "" "" "" "" "" ""]
[0 0 0 0 0 0 0 0 0 0]
```


### 字面量 

1. 指定数组长度：[length]type{v1, v2, ..., vlength}
2. 使用初始化元素数量推到数组长度：[...]type{v1, v2, ..., vlength}
3. 对指定位置元素进行初始化： [length]type{im:vm, ..., sin:in}

```go
package main

import "fmt"

func main() {
	var users [4]string = [4]string{"kk", "qq", "aa"}
	bounds := [...]int{10, 20, 30, 40, 50}
	teachers := [5]string{"kk"}
	nums := [10]int{1: 10, 3: 30, 5: 50, 8: 80} //对指定元素的位置进行初始化设定,配置1:10 其实是对数组中第二个元素初始化成10，第一个元素是0，配置3:30 其实是对数组中第4个元素进行初始化。

	fmt.Printf("%q\n", users)
	fmt.Println(bounds)
	fmt.Printf("%q\n", teachers)
	fmt.Println(nums)
}


PS E:\go-phase-two\go-course> go run test.go
["kk" "qq" "aa" ""]
[10 20 30 40 50]
["kk" "" "" "" ""]
[0 10 0 30 0 50 0 0 80 0]
```

### 操作 

1. 关系运算  ==、!= 

```go
package main

import "fmt"

func main() {
	// var users [4]string = [4]string{"kk", "qq", "aa"}
	bounds := [...]int{10, 20, 30, 40, 50}
	// teachers := [5]string{"kk"}
	// nums := [10]int{1: 10, 3: 30, 5: 50, 8: 80} //对指定元素的位置进行初始化设定,配置1:10 其实是对数组中第二个元素初始化成10，第一个元素是0，配置3:30 其实是对数组中第4个元素进行初始化。

	fmt.Println(bounds == [...]int{10, 20, 30, 40, 50})
	fmt.Println(bounds != [...]int{20, 10, 30, 40, 50})
}


PS E:\go-phase-two\go-course> go run test.go
true
true
```

2. 获取数组长度
   使用len函数可获取数组的长度

```go
package main

import "fmt"

func main() {
	// var users [4]string = [4]string{"kk", "qq", "aa"}
	bounds := [...]int{10, 20, 30, 40, 50}
	// teachers := [5]string{"kk"}
	// nums := [10]int{1: 10, 3: 30, 5: 50, 8: 80} //对指定元素的位置进行初始化设定,配置1:10 其实是对数组中第二个元素初始化成10，第一个元素是0，配置3:30 其实是对数组中第4个元素进行初始化。

	fmt.Println(len(bounds))
}


PS E:\go-phase-two\go-course> go run test.go
5
```


3. 访问&修改

通过对编号对数组元素进行访问和修改，元素的编号从左到右依次为:0, 1, 2, …, n(n
为数组长度-1)

```go
package main

import "fmt"

func main() {
	// var users [4]string = [4]string{"kk", "qq", "aa"}
	// bounds := [...]int{10, 20, 30, 40, 50}
	// teachers := [5]string{"kk"}
	// nums := [10]int{1: 10, 3: 30, 5: 50, 8: 80} //对指定元素的位置进行初始化设定,配置1:10 其实是对数组中第二个元素初始化成10，第一个元素是0，配置3:30 其实是对数组中第4个元素进行初始化。

	langs := [...]string{"go", "python", "c#", "java", "c", "c++", "lua"}

	fmt.Println(langs[0], langs[5], langs[len(langs)-1])

	fmt.Printf("%q\n", langs)

	langs[0] = "Go"

	fmt.Printf("%q\n", langs)
}


PS E:\go-phase-two\go-course> go run test.go
go c++ lua
["go" "python" "c#" "java" "c" "c++" "lua"]
["Go" "python" "c#" "java" "c" "c++" "lua"]
```

4. 切片： array[start:end]/array[start:end:cap](end<=cap<=len)获取数组的一部分
元素做为切片

```go
package main

import "fmt"

func main() {
	// var users [4]string = [4]string{"kk", "qq", "aa"}
	// bounds := [...]int{10, 20, 30, 40, 50}
	// teachers := [5]string{"kk"}
	// nums := [10]int{1: 10, 3: 30, 5: 50, 8: 80} //对指定元素的位置进行初始化设定,配置1:10 其实是对数组中第二个元素初始化成10，第一个元素是0，配置3:30 其实是对数组中第4个元素进行初始化。

	langs := [...]string{"go", "python", "c#", "java", "c", "c++", "lua"}

	fmt.Printf("%T, %q\n", langs[1:3], langs[1:3])
	fmt.Printf("%T, %q\n", langs[1:3:3], langs[1:3:3])
}

PS E:\go-phase-two\go-course> go run test.go
[]string, ["python" "c#"]
[]string, ["python" "c#"]
```


5. 遍历
   可以通过for+len+访问方式或for-range方式对数组中元素进行遍历

```go
package main

import "fmt"

func main() {
	// var users [4]string = [4]string{"kk", "qq", "aa"}
	// bounds := [...]int{10, 20, 30, 40, 50}
	// teachers := [5]string{"kk"}
	// nums := [10]int{1: 10, 3: 30, 5: 50, 8: 80} //对指定元素的位置进行初始化设定,配置1:10 其实是对数组中第二个元素初始化成10，第一个元素是0，配置3:30 其实是对数组中第4个元素进行初始化。

	langs := [...]string{"go", "python", "c#", "java", "c", "c++", "lua"}

	for i := 0; i < len(langs); i++ {
		fmt.Printf("%d： %q\n", i, langs[i])

	}

	for i, v := range langs {
		fmt.Printf("%d: %q\n", i, v)
	}
}

使用 for-range 遍历数组，range 返回两个元素分别为数组元素索引和值
```


## 多维数组

数组的元素也可以是数组类型，此时称为多维数组

1. 声明&初始化

```go
package main

import "fmt"

func main() {
	// 声明&初始化
	var students00 [6][8]string
	students01 := [...][2]string{{"kk", "bb"}, {"aa", "cc"}, {"xx", "ll"}}
	students02 := [3][3]string{{"kk", "ww"}, {"dd", "cc"}, {"uu", "qq"}}
	students03 := [3][3]string{0: {"kk", "ww", "cc"}, 1: {"22", "33", "33"}}
	students04 := [3][3]string{0: {0: "kk", 2: "33"}, 2: {1: "ww", 2: "33"}}

	fmt.Printf("%T, %T, %T, %T, %T\n", students00, students01, students02, students03, students04)

	fmt.Printf("%q\n", students00)
	fmt.Printf("%q\n", students01)
	fmt.Printf("%q\n", students02)
	fmt.Printf("%q\n", students03)
	fmt.Printf("%q\n", students04)

}

PS E:\go-phase-two\go-course> go run test.go
[6][8]string, [3][2]string, [3][3]string, [3][3]string, [3][3]string
[["" "" "" "" "" "" "" ""] ["" "" "" "" "" "" "" ""] ["" "" "" "" "" "" "" ""] ["" "" "" "" "" "" "" ""] ["" 
"" "" "" "" "" "" ""] ["" "" "" "" "" "" "" ""]]
[["kk" "bb"] ["aa" "cc"] ["xx" "ll"]]
[["kk" "ww" ""] ["dd" "cc" ""] ["uu" "qq" ""]]
[["kk" "ww" "cc"] ["22" "33" "33"] ["" "" ""]]
[["kk" "" "33"] ["" "" ""] ["" "ww" "33"]]
PS E:\go-phase-two\go-course>

```


2. 访问&修改 

```go
package main

import "fmt"

func main() {
	// 声明&初始化
	var students00 [6][8]string
	students01 := [...][2]string{{"kk", "bb"}, {"aa", "cc"}, {"xx", "ll"}}
	students02 := [3][3]string{{"kk", "ww"}, {"dd", "cc"}, {"uu", "qq"}}
	students03 := [3][3]string{0: {"kk", "ww", "cc"}, 1: {"22", "33", "33"}}
	students04 := [3][3]string{0: {0: "kk", 2: "33"}, 2: {1: "ww", 2: "33"}}

	fmt.Printf("%T, %T, %T, %T, %T\n", students00, students01, students02, students03, students04)

	fmt.Printf("%q\n", students00)
	fmt.Printf("%q\n", students01)
	fmt.Printf("%q\n", students02)
	fmt.Printf("%q\n", students03)
	fmt.Printf("%q\n", students04)

	// 修改元素
	students01[1] = [2]string{"jj", "nn"}
	fmt.Printf("%q\n", students01)

PS E:\go-phase-two\go-course> go run test.go
[6][8]string, [3][2]string, [3][3]string, [3][3]string, [3][3]string
[["" "" "" "" "" "" "" ""] ["" "" "" "" "" "" "" ""] ["" "" "" "" "" "" "" ""] ["" "" "" "" "" "" "" ""] ["" 
"" "" "" "" "" "" ""] ["" "" "" "" "" "" "" ""]]
[["kk" "bb"] ["aa" "cc"] ["xx" "ll"]]
[["kk" "ww" ""] ["dd" "cc" ""] ["uu" "qq" ""]]
[["kk" "ww" "cc"] ["22" "33" "33"] ["" "" ""]]
[["kk" "" "33"] ["" "" ""] ["" "ww" "33"]]
[["kk" "bb"] ["jj" "nn"] ["xx" "ll"]]
```

3. 遍历

```go
package main

import "fmt"

func main() {
	// 声明&初始化
	var students00 [6][8]string
	students01 := [...][2]string{{"kk", "bb"}, {"aa", "cc"}, {"xx", "ll"}}
	students02 := [3][3]string{{"kk", "ww"}, {"dd", "cc"}, {"uu", "qq"}}
	students03 := [3][3]string{0: {"kk", "ww", "cc"}, 1: {"22", "33", "33"}}
	students04 := [3][3]string{0: {0: "kk", 2: "33"}, 2: {1: "ww", 2: "33"}}

	fmt.Printf("%T, %T, %T, %T, %T\n", students00, students01, students02, students03, students04)

	fmt.Printf("%q\n", students00)
	fmt.Printf("%q\n", students01)
	fmt.Printf("%q\n", students02)
	fmt.Printf("%q\n", students03)
	fmt.Printf("%q\n", students04)

	// 修改元素
	students01[1] = [2]string{"jj", "nn"}
	fmt.Printf("%q\n", students01)

	// 遍历元素
	for i := 0; i < len(students01); i++ {
		for j := 0; j < len(students01[i]); j++ {
			fmt.Printf("[%d, %d]: %q\n", i, j, students01[i][j])
		}
	}

	for i, line := range students01 {
		for j, v := range line {
			fmt.Printf("[%d, %d]: %q\n", i, j, v)
		}
	}
}


PS E:\go-phase-two\go-course> go run test.go
[6][8]string, [3][2]string, [3][3]string, [3][3]string, [3][3]string
[["" "" "" "" "" "" "" ""] ["" "" "" "" "" "" "" ""] ["" "" "" "" "" "" "" ""] ["" "" "" "" "" "" "" ""] ["" 
"" "" "" "" "" "" ""] ["" "" "" "" "" "" "" ""]]
[["kk" "bb"] ["aa" "cc"] ["xx" "ll"]]
[["kk" "ww" ""] ["dd" "cc" ""] ["uu" "qq" ""]]
[["kk" "ww" "cc"] ["22" "33" "33"] ["" "" ""]]
[["kk" "" "33"] ["" "" ""] ["" "ww" "33"]]
[["kk" "bb"] ["jj" "nn"] ["xx" "ll"]]
[0, 0]: "kk"
[0, 1]: "bb"
[1, 0]: "jj"
[1, 1]: "nn"
[2, 0]: "xx"
[2, 1]: "ll"
[0, 0]: "kk"
[0, 1]: "bb"
[1, 0]: "jj"
[1, 1]: "nn"
[2, 0]: "xx"
[2, 1]: "ll"
```




## 切片（slice）

切片是长度可变的数组（具有相同数据类型的数据项组成的一组长度可变的序列），切片有3个部分组成

- 指针：指向切片第一个元素指向的数组元素的地址
- 长度：切片元素的树立
- 容量：切片开始到结束为止元素的数量

### 声明

切片声明需要指定组成元素的类型，但不需要指定存储元素的数量（长度）。在切片声明后，会被初始化为nil，表示暂不存在的切片。

```go
package main

import "fmt"

func main() {
	var names []string
	fmt.Printf("%T, %t, %v\n", names, names == nil, names)

}



PS E:\go-phase-two\go-course> go run test.go
[]string, true, []
```


### 初始化
1. 使用字面量初始化：[]type{v1, v2, ..., vn}
2. 使用字面量初始化空切片: []type{}
3. 指定长度和容量字面量初始化: []type{im:vm, in,vn, ilength:vlength}
4. 使用make函数初始化
   make([]type, len)/make([]type, len, cap), 通过make函数创建长度为len，容量为cap的切片，len必须小于cap。
5. 使用数组切片操作初始化： array[start:end]/array[start:end:cap](end<=cap<=len)

```go

package main

import "fmt"

func main() {
	scores := [...]int{60, 70, 80, 95, 85}

	var scores00 []int = []int{80, 90, 95}

	scores01 := []int{}
	scores02 := []int{0: 80, 2: 95}
	scores03 := make([]int, 3)
	scores04 := make([]int, 3, 5)
	scores05 := scores[1:3]
	scores06 := scores[1:3:3]

	fmt.Printf("%T, %T, %T, %T, %T, %T\n", scores00, scores01, scores02, scores03, scores04, scores05, scores06)

	fmt.Println(scores00)
	fmt.Println(scores01)
	fmt.Println(scores02)
	fmt.Println(scores03)
	fmt.Println(scores04)
	fmt.Println(scores05)
	fmt.Println(scores06)
}


PS E:\go-phase-two\go-course> go run test.go
[]int, []int, []int, []int, []int, []int
%!(EXTRA []int=[70 80])[80 90 95]
[]
[80 0 95]
[0 0 0]
[0 0 0]
[70 80]
[70 80]
```


### 操作 
获取切片长度和容量
使用len 函数可以获取切片的长度，使用cap函数可以获取切片的容量

```go

```