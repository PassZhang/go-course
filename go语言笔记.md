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
- complex： 工厂函数，通过两个参数创建一个复数()
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
1. 获取切片长度和容量
使用len 函数可以获取切片的长度，使用cap函数可以获取切片的容量

```go
package main

import "fmt"

func main() {
	students := make([]string, 3, 5)
	fmt.Println(len(students), cap(students))

	fmt.Printf("%q\n", students)
}



PS E:\go-phase-two\go-course> go run test.go
3 5
["" "" ""]
```


2. 访问和修改
   通过对编号对切片元素进行访问和修改，元素的编号从左到右依次为：0， 1， 2， ..., n(n为切片长度-1)

```go
package main

import "fmt"

func main() {
	students := make([]string, 3, 5)

	fmt.Printf("%q, %q\n", students[0], students[1])
	students[0] = "KK"
	students[1] = "bb"
	students[2] = "aa"

	fmt.Printf("%q, %q, %q\n", students[0], students[1], students[2])
	fmt.Printf("%q\n", students)
}


PS E:\go-phase-two\go-course> go run test.go
"", ""
"KK", "bb", "aa"
["KK" "bb" "aa"]
```


3. 切片：slice[start:end]用于创建一个新的切片， end <= src_cap
   

```go
package main

import "fmt"

func main() {
	teachers := [...]string{"kk", "ww", "aa", "xx", "dd"}
	teachers00 := teachers[:]

	teachers01 := teachers[0:3]
	teachers02 := teachers[1:4]
	teachers03 := teachers00[1:3]

	fmt.Printf("%q\n", teachers)
	fmt.Printf("%d, %d, %q\n", len(teachers), cap(teachers01), teachers01)
	fmt.Printf("%d, %d, %q\n", len(teachers01), cap(teachers01), teachers01)
	fmt.Printf("%d, %d, %q\n", len(teachers02), cap(teachers02), teachers02)
	fmt.Printf("%d, %d, %q\n", len(teachers03), cap(teachers03), teachers03)

}


PS E:\go-phase-two\go-course> go run test.go
["kk" "ww" "aa" "xx" "dd"]
5, 5, ["kk" "ww" "aa"]
3, 5, ["kk" "ww" "aa"]
3, 4, ["ww" "aa" "xx"]
2, 4, ["ww" "aa"]
```
新创建切片长度和容量计算： len：end-start, cap: scr_cap-start

切片共享底层数组，若某个切片元素发生变化，则数组和其他有共享元素的切片也会发生变化

```go
package main

import "fmt"

func main() {
	teachers := [...]string{"kk", "ww", "aa", "xx", "dd"}
	teachers00 := teachers[:]

	teachers01 := teachers[0:3]
	teachers02 := teachers[1:4]
	teachers03 := teachers00[1:3]

	fmt.Printf("%q\n", teachers)

	teachers02[2] = "小林"

	fmt.Printf("%d, %d, %q\n", len(teachers), cap(teachers), teachers)

	fmt.Printf("%d, %d, %q\n", len(teachers), cap(teachers01), teachers01)
	fmt.Printf("%d, %d, %q\n", len(teachers01), cap(teachers01), teachers01)
	fmt.Printf("%d, %d, %q\n", len(teachers02), cap(teachers02), teachers02)
	fmt.Printf("%d, %d, %q\n", len(teachers03), cap(teachers03), teachers03)
}

PS E:\go-phase-two\go-course> go run test.go
["kk" "ww" "aa" "xx" "dd"]
5, 5, ["kk" "ww" "aa" "小林" "dd"]
5, 5, ["kk" "ww" "aa"]
3, 5, ["kk" "ww" "aa"]
3, 4, ["ww" "aa" "小林"]
2, 4, ["ww" "aa"]
```


slice[start: end: cap]可用于限制新切片的容量值， end<= cap <= src_cap

```go
package main

import "fmt"

func main() {
	teachers := [...]string{"kk", "ww", "aa", "xx", "dd"}
	teachers00 := teachers[:]

	teachers01 := teachers[0:3]
	teachers02 := teachers[1:4]
	teachers03 := teachers00[1:3]

	fmt.Printf("%q\n", teachers)

	teachers02[2] = "小林"

	fmt.Printf("%d, %d, %q\n", len(teachers), cap(teachers), teachers)

	fmt.Printf("%d, %d, %q\n", len(teachers), cap(teachers01), teachers01)
	fmt.Printf("%d, %d, %q\n", len(teachers01), cap(teachers01), teachers01)
	fmt.Printf("%d, %d, %q\n", len(teachers02), cap(teachers02), teachers02)
	fmt.Printf("%d, %d, %q\n", len(teachers03), cap(teachers03), teachers03)

	teachers04 := teachers[1:4:4]
	teachers05 := teachers00[1:3:3]

	fmt.Printf("%d, %d, %q\n", len(teachers02), cap(teachers02), teachers02)
	fmt.Printf("%d, %d, %q\n", len(teachers03), cap(teachers03), teachers03)
	fmt.Printf("%d, %d, %q\n", len(teachers04), cap(teachers04), teachers04)
	fmt.Printf("%d, %d, %q\n", len(teachers05), cap(teachers05), teachers05)

}



PS E:\go-phase-two\go-course> go run test.go
["kk" "ww" "aa" "xx" "dd"]
5, 5, ["kk" "ww" "aa" "小林" "dd"]
5, 5, ["kk" "ww" "aa"]
3, 5, ["kk" "ww" "aa"]
3, 4, ["ww" "aa" "小林"]
2, 4, ["ww" "aa"]
3, 4, ["ww" "aa" "小林"]
2, 4, ["ww" "aa"]
3, 3, ["ww" "aa" "小林"]
2, 2, ["ww" "aa"]
```



4. 遍历 ()
   可以通过for+len+访问方式或for-range方式对切片中元素进行遍历

```go
package main

import "fmt"

func main() {
	langs := []string{"go", "python", "c#", "c", "c++", "lua", "lisp", "php", "rust"}

	fmt.Printf("%q\n", langs[1:4:4])
	fmt.Printf("%q\n", langs[1:3:3])

	for i := 0; i < len(langs); i++ {
		fmt.Printf("%d, %q\n", i, langs[i])
	}

	for i, v := range langs {
		fmt.Printf("%d, %q\n", i, v)
	}
}


PS E:\go-phase-two\go-course> go run test.go
["python" "c#" "c"]
["python" "c#"]
0, "go"
1, "python"
2, "c#"
3, "c"
4, "c++"
5, "lua"
6, "lisp"
7, "php"
8, "rust"
0, "go"
1, "python"
2, "c#"
3, "c"
4, "c++"
5, "lua"
6, "lisp"
7, "php"
8, "rust"
```

使用for-range遍历切片，range返回两个元素分别为切片元素索引和值


5. 增加元素
   使用append对切片增加一个或多个元素并返回修改后切片，当长度在容量范围内时值增加长度，容量和底层数组不变。当长度常容量范围则会创建一个新的底层数组并对容量进行智能运算（元素数量<1024时,约按原容量1倍增加， >1024时约按原容量0.25倍增加）

```go
package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}

	nums2 := nums[:]
	fmt.Printf("%d, %d, %v\n", len(nums), cap(nums), nums)
	fmt.Printf("%d, %d, %v\n", len(nums2), cap(nums2), nums)

	nums2 = append(nums2, 6)
	fmt.Printf("%d, %d, %v\n", len(nums), cap(nums), nums)
	fmt.Printf("%d, %d, %v\n", len(nums2), cap(nums2), nums2)

	nums2[0] = 0
	fmt.Printf("%d, %d %v\n", len(nums), cap(nums), nums)
	fmt.Printf("%d, %d, %v\n", len(nums2), cap(nums2), nums2)

	nums2 = append(nums2, 0, 1, 3, 4, 5)
	fmt.Printf("%d, %d\n", len(nums2), cap(nums2))
}


PS E:\go-phase-two\go-course> go run test.go
5, 5, [1 2 3 4 5]
5, 5, [1 2 3 4 5]
5, 5, [1 2 3 4 5]
6, 10, [1 2 3 4 5 6]
5, 5 [1 2 3 4 5]
6, 10, [0 2 3 4 5 6]
11, 20
```

6. 复制切片到另一个切片
   复制元素数量为src元素数量和dest元素数量的最小值

```go
package main

import "fmt"

func main() {
	user01 := []string{"00", "01"}
	user02 := []string{"10", "11", "12"}
	user03 := []string{"20", "21", "22", "23"}

	fmt.Printf("%q, %q, %q\n", user01, user02, user03)
	copy(user01, user02)
	fmt.Printf("%q, %q, %q\n", user01, user02, user03)
	copy(user03, user02)
	fmt.Printf("%q, %q, %q\n", user01, user02, user03)

}

PS E:\go-phase-two\go-course> go run test.go
["00" "01"], ["10" "11" "12"], ["20" "21" "22" "23"]
["10" "11"], ["10" "11" "12"], ["20" "21" "22" "23"]
["10" "11"], ["10" "11" "12"], ["10" "11" "12" "23"]

```


### 使用 
1. 移除元素
```go

package main

import "fmt"

func main() {
	elements := []int{0, 1, 2, 3, 4, 5}
	copy(elements[3:], elements[4:])
	fmt.Println(elements)
	fmt.Println(elements[:len(elements)-1])
}


PS E:\go-phase-two\go-course> go run test.go
[0 1 2 4 5 5]
[0 1 2 4 5]

```

2. 队列
   先进先出

```go

package main

import "fmt"

func main() {
	queue := []int{}

	queue = append(queue, 1)
	queue = append(queue, 3)
	queue = append(queue, 2)

	fmt.Println(queue[0])
	queue = queue[1:]
	fmt.Println(queue[0])
	queue = queue[1:]
	fmt.Println(queue[0])

}


PS E:\go-phase-two\go-course> go run test.go
1
3
2
```




3. 堆栈

先进后出

```go
package main

import "fmt"

func main() {
	stack := []int{}
	stack = append(stack, 1)
	stack = append(stack, 3)
	stack = append(stack, 2)

	fmt.Println(stack[len(stack)-1])
	stack = stack[:len(stack)-1]
	fmt.Println(stack[len(stack)-1])
	stack = stack[:len(stack)-1]
	fmt.Println(stack[len(stack)-1])
}


PS E:\go-phase-two\go-course> go run test.go
2
3
1
```

## 多维切片

切片的元素也可以是切片类型，此时称为多维切片

```go
package main

import "fmt"

func main() {
	// 声明&初始化
	points := [][]int{{1, 1}, {1, 2, 3}}
	fmt.Printf("%T, %v, %V, %d\n", points, points, points[0], points[0][0])

	// 修改
	points[0] = []int{2, 2}
	points[1][1] = 3

	fmt.Println(points)

	// 切片
	fmt.Println(points[0:2])

	// 遍历
	for i := 0; i < len(points); i++ {
		for j := 0; j < len(points[i]); j++ {
			fmt.Printf("[%d, %d]: %v\n", i, j, points[i][j])
		}
	}

	for i, line := range points {
		for j, v := range line {
			fmt.Printf("[%d, %d]: %v\n", i, j, v)
		}
	}

	// append
	points = append(points, []int{2, 3, 1})
	points[0] = append(points[0], 1)
	fmt.Println(points)

	// copy
	points2 := [][]int{{}, {}}

	copy(points2, points)
	fmt.Println(points2)
}



PS E:\go-phase-two\go-course> go run test.go
[][]int, [[1 1] [1 2 3]], [%!V(int=1) %!V(int=1)], 1
[[2 2] [1 3 3]]
[[2 2] [1 3 3]]
[0, 0]: 2
[0, 1]: 2
[1, 0]: 1
[1, 1]: 3
[1, 2]: 3
[0, 0]: 2
[0, 1]: 2
[1, 0]: 1
[1, 1]: 3
[1, 2]: 3
[[2 2 1] [1 3 3] [2 3 1]]
[[2 2 1] [1 3 3]]
```



### 常用包 

sort


## 映射（map）

映射是存储一系列无序的 key/value 对，通过 key 来对 value 进行操作（增、删、改、查）。
映射的 key 只能为可使用==运算符的值类型（字符串、数字、布尔、数组），value 可以为任意类型

### 声明 

map 声明需要制定组成元素key 和value 的类型，在声明后，会被初始化为nil， 表示暂不存在的映射。

```go
package main

import "fmt"

func main() {
	var tels map[string]string
	var points map[[2]int]float64

	fmt.Printf("%T, %t, %v\n", tels, tels == nil, tels)
	fmt.Printf("%T, %t, %v\n", points, points == nil, points)

}


PS E:\go-phase-two\go-course> go run test.go
map[string]string, true, map[]
map[[2]int]float64, true, map[]
```


### 初始化 
1. 使用字面量初始化：map[ktype]vtype{k1:v1, k2:v2, ..., kn:vn}
2. 使用字面量初始化空映射： map[ktype]vtype{}
3. 使用make函数初始化
   make(map[ktype]vtype), 通过make函数创建映射

```go
package main

import "fmt"

func main() {
	var tels = map[string]string{"kk": "1520000000", "ww": "1859999999"}
	fmt.Printf("%q\n", tels)

	var points = map[[2]int]float64{{1, 2}: 3, {4, 5}: 6}
	fmt.Println(points)

	scores := map[string]int{"kk": 80, "ww": 90}
	fmt.Println(scores)

	heighs := map[string]float64{}
	fmt.Println(heighs)

	weights := make(map[string]float64)
	fmt.Println(weights)
}


PS E:\go-phase-two\go-course> go run test.go
map["kk":"1520000000" "ww":"1859999999"]
map[[1 2]:3 [4 5]:6]
map[kk:80 ww:90]
map[]
map[]
```

### 操作 
1. 获取元素的数量
   使用len函数获取映射元素的数量

```go
package main

import "fmt"

func main() {
	var tels = map[string]string{"kk": "1520000000", "ww": "1859999999"}
	fmt.Printf("%q\n", tels)

	var points = map[[2]int]float64{{1, 2}: 3, {4, 5}: 6}
	fmt.Println(points)

	scores := map[string]int{"kk": 80, "ww": 90}
	fmt.Println(scores)

	heighs := map[string]float64{}
	fmt.Println(heighs)

	weights := make(map[string]float64)
	fmt.Println(weights)

	fmt.Println(len(tels), len(points), len(scores), len(heighs), len(weights))

	
}


PS E:\go-phase-two\go-course> go run test.go
map["kk":"1520000000" "ww":"1859999999"]
map[[1 2]:3 [4 5]:6]
map[kk:80 ww:90]
map[]
map[]
2 2 2 0 0
```


2. 访问 

```go
package main

import "fmt"

func main() {
	students := map[int]string{1: "kk", 2: "ww"}
	students01 := map[int]map[string]string{1: map[string]string{"name": "kk", "tel": "152xxxxxxxx"}}

	fmt.Printf("%v, %q, %q\n", students, students[1], students[3])

	fmt.Printf("%v, %q, %q, %t\n", students01, students01[1], students01[3], students01[3] == nil)
}

当访问 key 存在与映射时则返回对应的值，否则返回值类型的零值.
```

3. 判断key 是否存在
   通过key 访问元素时可以接受两个值，第一个值为value， 第二个值为bool类型表示元素是否存在，若存在为true，否则为false

```go
package main

import "fmt"

func main() {
	students := map[int]string{1: "kk", 2: "ww"}
	students01 := map[int]map[string]string{1: map[string]string{"name": "kk", "tel": "152xxxxxxxx"}}

	fmt.Printf("%v, %q, %q\n", students, students[1], students[3])

	fmt.Printf("%v, %q, %q, %t\n", students01, students01[1], students01[3], students01[3] == nil)

	student, ok := students[1]
	fmt.Printf("%t, %v\n", ok, student)

	student, ok = students[2]
	fmt.Printf("%t, %v\n", ok, student)

	student01, ok := students01[1]
	fmt.Printf("%t, %v\n", ok, student01)

	student01, ok = students01[2]
	fmt.Printf("%t, %v\n", ok, student01)

}


PS E:\go-phase-two\go-course> go run test.go
map[1:kk 2:ww], "kk", ""
map[1:map[name:kk tel:152xxxxxxxx]], map["name":"kk" "tel":"152xxxxxxxx"], map[], true
true, kk
true, ww
true, map[name:kk tel:152xxxxxxxx]
false, map[]
```


4. 修改&增加
   使用key对映射赋值时当key 存在则修改key 对应的value， 若key 不存在则增加key 和value 

```go
package main

import "fmt"

func main() {
	students := map[int]string{1: "kk", 2: "ww"}
	students01 := map[int]map[string]string{1: map[string]string{"name": "kk", "tel": "152xxxxxxxx"}}

	fmt.Printf("%v, %q, %q\n", students, students[1], students[3])

	fmt.Printf("%v, %q, %q, %t\n", students01, students01[1], students01[3], students01[3] == nil)

	student, ok := students[1]
	fmt.Printf("%t, %v\n", ok, student)

	student, ok = students[2]
	fmt.Printf("%t, %v\n", ok, student)

	student01, ok := students01[1]
	fmt.Printf("%t, %v\n", ok, student01)

	student01, ok = students01[2]
	fmt.Printf("%t, %v\n", ok, student01)

	students[1] = "KK" // key 存在，修改
	students[3] = "WW" // key 不存在, 增加

	fmt.Println(students)

	students01[1]["tel"] = "1520000000"                                                 // key 存在，修改
	students01[1]["addr"] = "西安市"                                                       // key 不存在，增加
	students01[2] = map[string]string{"name": "www", "tel": "158000000", "addr": "北京市"} // key 不存在，增加
	fmt.Println(students01)

	students01[2] = map[string]string{"name": "xxx", "tel": "18888888888", "addr": "西安市"} // key存在， 修改
	fmt.Println(students01)

}


PS E:\go-phase-two\go-course> go run test.go
map[1:kk 2:ww], "kk", ""
map[1:map[name:kk tel:152xxxxxxxx]], map["name":"kk" "tel":"152xxxxxxxx"], map[], true
true, kk
true, ww
true, map[name:kk tel:152xxxxxxxx]
false, map[]
map[1:KK 2:ww 3:WW]
map[1:map[addr:西安市 name:kk tel:1520000000] 2:map[addr:北京市 name:www tel:158000000]]
map[1:map[addr:西安市 name:kk tel:1520000000] 2:map[addr:西安市 name:xxx tel:18888888888]]

```

5. 删除
   使用delete 函数删除映射中已经存在的key

```go
package main

import "fmt"

func main() {
	students := map[int]string{1: "kk", 2: "ww"}
	students01 := map[int]map[string]string{1: map[string]string{"name": "kk", "tel": "152xxxxxxxx"}}

	fmt.Printf("%v, %q, %q\n", students, students[1], students[3])

	fmt.Printf("%v, %q, %q, %t\n", students01, students01[1], students01[3], students01[3] == nil)

	student, ok := students[1]
	fmt.Printf("%t, %v\n", ok, student)

	student, ok = students[2]
	fmt.Printf("%t, %v\n", ok, student)

	student01, ok := students01[1]
	fmt.Printf("%t, %v\n", ok, student01)

	student01, ok = students01[2]
	fmt.Printf("%t, %v\n", ok, student01)

	students[1] = "KK" // key 存在，修改
	students[3] = "WW" // key 不存在, 增加

	fmt.Println(students)

	students01[1]["tel"] = "1520000000"                                                 // key 存在，修改
	students01[1]["addr"] = "西安市"                                                       // key 不存在，增加
	students01[2] = map[string]string{"name": "www", "tel": "158000000", "addr": "北京市"} // key 不存在，增加
	fmt.Println(students01)

	students01[2] = map[string]string{"name": "xxx", "tel": "18888888888", "addr": "西安市"} // key存在， 修改
	fmt.Println(students01)

	delete(students, 1)
	delete(students, 4)
	fmt.Println(students)
	delete(students01[1], "addr")
	delete(students01, 2)
	fmt.Println(students01)
}


PS E:\go-phase-two\go-course> go run test.go
map[1:kk 2:ww], "kk", ""
map[1:map[name:kk tel:152xxxxxxxx]], map["name":"kk" "tel":"152xxxxxxxx"], map[], true
true, kk
true, ww
true, map[name:kk tel:152xxxxxxxx]
false, map[]
map[1:KK 2:ww 3:WW]
map[1:map[addr:西安市 name:kk tel:1520000000] 2:map[addr:北京市 name:www tel:158000000]]
map[1:map[addr:西安市 name:kk tel:1520000000] 2:map[addr:西安市 name:xxx tel:18888888888]]
map[2:ww 3:WW]
map[1:map[name:kk tel:1520000000]]
```



6. 遍历 
   可通过for-range对映射中个元素进行遍历，range返回两个元素分别为映射的key 和 value

```go
package main

import "fmt"

func main() {
	students := map[int]string{1: "kk", 2: "ww"}
	students01 := map[int]map[string]string{1: map[string]string{"name": "kk", "tel": "152xxxxxxxx"}}

	fmt.Printf("%v, %q, %q\n", students, students[1], students[3])

	fmt.Printf("%v, %q, %q, %t\n", students01, students01[1], students01[3], students01[3] == nil)

	student, ok := students[1]
	fmt.Printf("%t, %v\n", ok, student)

	student, ok = students[2]
	fmt.Printf("%t, %v\n", ok, student)

	student01, ok := students01[1]
	fmt.Printf("%t, %v\n", ok, student01)

	student01, ok = students01[2]
	fmt.Printf("%t, %v\n", ok, student01)

	students[1] = "KK" // key 存在，修改
	students[3] = "WW" // key 不存在, 增加

	fmt.Println(students)

	students01[1]["tel"] = "1520000000"                                                 // key 存在，修改
	students01[1]["addr"] = "西安市"                                                       // key 不存在，增加
	students01[2] = map[string]string{"name": "www", "tel": "158000000", "addr": "北京市"} // key 不存在，增加
	fmt.Println(students01)

	students01[2] = map[string]string{"name": "xxx", "tel": "18888888888", "addr": "西安市"} // key存在， 修改
	fmt.Println(students01)

	for k, v := range students {
		fmt.Printf("%v: %v\n", k, v)
	}
}



PS E:\go-phase-two\go-course> go run test.go
map[1:kk 2:ww], "kk", ""
map[1:map[name:kk tel:152xxxxxxxx]], map["name":"kk" "tel":"152xxxxxxxx"], map[], true
true, kk
true, ww
true, map[name:kk tel:152xxxxxxxx]
false, map[]
map[1:KK 2:ww 3:WW]
map[1:map[addr:西安市 name:kk tel:1520000000] 2:map[addr:北京市 name:www tel:158000000]]
map[1:map[addr:西安市 name:kk tel:1520000000] 2:map[addr:西安市 name:xxx tel:18888888888]]
1: KK
2: ww
3: WW
```

### 使用

统计演讲稿中"我有一个梦想"中各个英文字符出现的次数。

```go
package main

import "fmt"

func main() {
	article := `but i think that whatever personal rule of life you may choose it should not, except in rare and heroic cases, be incompatible with happiness. if you look around at the men and women whom you can call happy, you will see that they all have certain things in common.`

	stats := make(map[rune]int)
	for _, ch := range article {
		if ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' {
			stats[ch]++

		}
	}

	for ch, cnt := range stats {
		fmt.Printf("%c: %v\n", ch, cnt)
	}
}


PS E:\go-phase-two\go-course> go run test.go
a: 19
e: 22
v: 2
d: 4
i: 15
h: 15
n: 16
k: 2
x: 1
l: 12
f: 3
c: 9
u: 8
p: 7
s: 9
o: 19
y: 7
m: 7
g: 1
b: 3
t: 17
w: 5
r: 8
```





# 函数 
## 定义&调用
函数用于代码块的逻辑封装，提供代码复用的最基本方式

### 定义 

函数包含函数名、行参列表、函数体和返回值列表，使用func进行声明，函数无参数或返回值时则形参列表和返回则列表省略。

```go
package main

import "fmt"

func main() {
	func name(parameters) returns {
		body
	}
}



```

形参列表需要描述参数名及参数类型， 所有形参为函数块局部变量。返回值需要描述返回值类型。

举例：
1. 无参&无返回值

```go
package main

import "fmt"

func main() {
}

func sayHello() {
	fmt.Println("Hello World")
}

```

2. 有参&无返回值

```go
package main

import "fmt"

func main() {
	sayHi("kk")
}

func sayHi(name string) {
	fmt.Printf("Hi, %s\n", name)
}

```


3. 有参&有返回值

```go
package main

import "fmt"

func main() {

	n1, n2 := 1, 2
	fmt.Printf("%d + %d = %d\n", n1, n2, add(n1, n2))
}
func add(n1 int, n2 int) int {
	return n1 + n2
}


PS E:\go-phase-two\go-course> go run test.go
1 + 2 = 3
```



### 调用 
函数通过函数名（实参列表）， 在调用过程中实参的每个数据会赋值给形参中的每个变量，因此实参列表类型和数量需要函数定义的形参一一对应。 针对函数返回值可通过变量赋值的方式接受。

```go
package main

import "fmt"

func main() {
	// 调用无参无返回值函数
	sayHello()

	// 调用有参无返回值函数
	sayHi("KK")

	// 调用有参有返回值函数
	n1, n2 := 1, 2
	fmt.Printf("%d + %d = %d\n", n1, n2, add(n1, n2))

	n3 := add(4, 5)
	fmt.Println(n3)

	// 忽略函数返回值
	add(3, 4)
}

func add(n1 int, n2 int) int {
	return n1 + n2
}

func sayHello() {
	fmt.Println("Hello World")
}

func sayHi(name string) {
	fmt.Printf("Hi, %s\n", name)

}


Hello World
Hi, KK
1 + 2 = 3
9
```


## 参数 

### 类型合并

在声明函数中若存在多个连续形参类型相同可以只保留最后一个参数类型名

```go
package main

import "fmt"

func main() {
	mergeFuncArgsType(1, 2, "3", "4", "5", true)

}

// 合并相同类型参数类型名
func mergeFuncArgsType(n1, n2 int, s1, s2, s3 string, b1 bool) {
	fmt.Printf("%T, %T, %T, %T, %T, %T\n", n1, n2, s1, s2, s3, b1)
	fmt.Println(n1, n2, s1, s2, s3, b1)
}



PS E:\go-phase-two\go-course> go run test.go
int, int, string, string, string, bool
1 2 3 4 5 true
```



### 可变参数 

某些情况下函数需要处理形参数量可变，需要运算符…声明可变参数函数或在调用时传递可变参数

1. 定义 
   可变参数只能定义一个且只能在参数列表末端。在调用函数后，可变参数则被初始化为对应类型的切片。 

```go
package main

import "fmt"

func main() {
	// 调用可变参数函数
	printArgs(1, 2)
	printArgs(1, 2, "3", "4", "5")
}

// 定义可变参数列表函数，至少有2个参数
// 打印所有参数到控制台

func printArgs(n1, n2 int, args ...string) {
	fmt.Printf("%T, %T, %T\n", n1, n2, args)
	fmt.Println(n1, n2, args)
}

PS E:\go-phase-two\go-course> go run test.go
int, int, []string
1 2 []
int, int, []string
1 2 [3 4 5]
```


2. 传递
   在调用函数时， 也可以使用运算符...将切片解包传递到可变参数函数中

```go
package main

import "fmt"

func main() {
	// 调用可变参数函数
	printArgs(1, 2)
	printArgs(1, 2, "3", "4", "5")

	printArgs(1, 2, []string{"3", "4", "5", "6"}...)
	args := []string{"3", "4", "5", "6", "8"}
	printArgs(1, 2, args...)
	printArgs(1, 2, args[:3]...)
}

// 定义可变参数列表函数，至少有2个参数
// 打印所有参数到控制台

func printArgs(n1, n2 int, args ...string) {
	fmt.Printf("%T, %T, %T\n", n1, n2, args)
	fmt.Println(n1, n2, args)
}


PS E:\go-phase-two\go-course> go run test.go
int, int, []string
1 2 []
int, int, []string
1 2 [3 4 5]
int, int, []string
1 2 [3 4 5 6]
int, int, []string
1 2 [3 4 5 6 8]
int, int, []string
1 2 [3 4 5]
```


## 返回值 
在函数提中可以使用return 关键字为函数调用这提供函数计算结果

### 多返回值
go 语言支持函数有多个返回值，在声明函数时使用括号包含多有返回值类型，并使用return返回对应数量的用逗号分隔数据。

```go
package main

import "fmt"

func main() {
	fmt.Println(calc(1, 2))

}

// 定义有多个返回值的函数
// 计算两个参数的四则运算结果并返回

func calc(n1, n2 int) (int, int, int, int) {
	return n1 + n2, n1 - n2, n1 * n2, n1 / n2
}


PS E:\go-phase-two\go-course> go run test.go
3 -1 2 0
```

### 命名返回值 
在函数返回值中可指定变量名，变量在调用时会根据类型使用零值进行初始化，在函数体重可进行赋值，同时在调用return时不需要添加返回值，go 语言自动将变量的最终结果进行返回

在使用命名返回值时，当声明函数中存在多个练习返回值类型相同可只保留最后一个返回值类型名。



```go
package main

import "fmt"

func main() {
	fmt.Println(calcReturnNamecalc(5, 2))
}

// 定义命名返回值函数

func calcReturnNamecalc(n1, n2 int) (sum, difference, product, quotient int) {
	sum, difference, product, quotient = n1+n2, n1-n2, n1*n2, n1/n2
	return
}


PS E:\go-phase-two\go-course> go run test.go
7 3 10 2
```

## 递归

递归是指函数直接或间接调用自己，递归常用于解决分治问题，将大问题分解为相同的小问题进行解决，需要关注中止条件。

1. 计算n阶乘

```go
package main

import "fmt"

func main() {
	fmt.Println(factorial(10))
}

/*
计算n的阶乘
n < 0 ： 错误的
n = 0 或 1： 返回1
n > 0 : n! = n * n - 1!
*/
func factorial(n int) int {
	if n < 0 {
		return -1

	} else if n == 0 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}

PS E:\go-phase-two\go-course> go run test.go
3628800

```

```go
package main

import "fmt"

func main() {
	fmt.Println(factorial(3))
}

/*
计算n的阶乘
n < 0 错误
n = 0 返回1
n > 0: n! = n *  (n - 1)!
*/

func factorial(n int) int {
	if n < 0 {
		return 0
	} else if n == 0 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}


PS E:\go-phase-two\go-course> go run test.go
6
```


2. 汉诺塔小游戏

基本版本：
```go
package main

import "fmt"

func main() {
	tower("塔1", "塔2", "塔3", 3)

}

/*
汉诺塔小游戏
将所有a柱上的圆盘借助b柱移动到c柱，在移动过程中保证每个柱子上面的圆盘比下面的圆盘小
n : a -> c(b)
n=1 : a -> c
n>1: n-1(a -> b(c); a -> c; n-1(b -> c((a))
*/

func tower(x, y, z string, n int) {
	if n <= 0 {
		return
	}

	if n == 1 {
		fmt.Printf("%s - > %s\n", x, z)
		return
	}

	tower(x, z, y, n-1)
	fmt.Printf("%s - > %s\n", x, z)
	tower(y, x, z, n-1)
}


PS E:\go-phase-two\go-course> go run test.go
塔1 - > 塔3
塔3 - > 塔2
塔1 - > 塔3
塔2 - > 塔1
塔2 - > 塔3
塔1 - > 塔3
```


升级版本:
```go
package main

import "fmt"

func main() {
	tower("塔1", "塔2", "塔3", 3)
	mv("")
}

/*
汉诺塔小游戏
将所有a柱上的圆盘借助b柱移动到c柱，在移动过程中保证每个柱子上面的圆盘比下面的圆盘小
n : a -> c(b)
n=1 : a -> c
n>1: n-1(a -> b(c); a -> c; n-1(b -> c((a))
*/

func tower(x, y, z string, n int) {
	if n <= 0 {
		return
	}

	if n == 1 {
		fmt.Printf("%s - > %s\n", x, z)
		return
	}

	tower(x, z, y, n-1)
	fmt.Printf("%s - > %s\n", x, z)
	tower(y, x, z, n-1)
}

func mv(N, M string, disks int) {
	fmt.Printf("第 %d 次移动 : 把 %d 号圆盘从 %s ->移到->  %s\n", n, disks, N, M)
}

```


## 函数类型
函数也可以赋值给变量，存储在数组、切片、映射中， 也可以作为参数传递给函数或作为函数返回值进行返回。

### 声明&初始化&调用

```go
package main

import "fmt"

func main() {

	// 定义函数类型变量，并使用零值nil进行初始化
	var callback func(n1, n2 int) (r1, r2, r3, r4 int)
	fmt.Printf("%T, %v\n", callback, callback)

	callback = calc             // 赋值为函数calc
	fmt.Println(callback(5, 2)) // 调用calc函数

	callback = calcReturnNamecalc // 赋值函数为calcReturnNamecalc
	fmt.Println(callback(8, 2))
}

// 定义有多个返回值的函数
// 计算两个参数的四则运算结果并返回
func calc(n1, n2 int) (int, int, int, int) {
	return n1 + n2, n1 - n2, n1 * n2, n1 / n2
}

// 定义命名返回值函数
func calcReturnNamecalc(n1, n2 int) (sum, difference, product, quotient int) {
	sum, difference, product, quotient = n1+n2, n1-n2, n1*n2, n1/n2
	return
}


PS E:\go-phase-two\go-course> go run test.go
func(int, int) (int, int, int, int), <nil>
7 3 10 2
10 6 16 4
```


### 声明&调用参数类型为函数的函数

```go
package main

import "fmt"

func main() {
	names := []string{"kk", "ww", "aa"}
	pringResult(line, names...)
	pringResult(colume, names...)
}

/*
定义接收函数类型作为参数的函数
*/

func pringResult(pf func(...string), list ...string) {
	pf(list...)
}

func line(list ...string) {
	fmt.Print("|")
	for _, e := range list {
		fmt.Print(e)
		fmt.Print("\t|")
	}
	fmt.Println()
	fmt.Println()
}

func colume(list ...string) {
	for _, e := range list {
		fmt.Println(e)
	}
	fmt.Println()
}


PS E:\go-phase-two\go-course> go run test.go
|kk     |ww     |aa     |

kk
ww
aa
```

## 匿名函数与闭包

### 匿名函数 

不需要定义名字的函数叫做匿名函数，常用作帮助函数在局部代码中使用或作为其他函数的参数。

```go
// 定义匿名函数并赋值给hi 
hi := func(name string) {
	fmt.Printf("Hi, %s\n", name)
}

hi("kk")
hi("ww")

//定义匿名函数并进行调用
func() {
	fmt.Println("我是匿名函数，我在使用")
}()

// 使用匿名函数作为pringResult的参数
printResult(func(list ...string) {
	for i, v := range list {
		fmt.Printf("%d: %s\n", i, v)
	}
}, names...)



```


### 闭包

匿名函数又叫闭包，是指在函数内定义的匿名函数引用外部函数的变量，只要匿名函数继续使用外部函数赋值的变量不会被自动销毁。

```go
package main

import "fmt"

func main() {
	// 使用闭包函数
	base2 := addBase(2)
	base10 := addBase(10)

	fmt.Println(base2(1), base10(1))
	fmt.Println(base2(5), base10(5))
	fmt.Println(base2(10), base10(10))
}

// 定义闭包函数，返回一个匿名函数用于计算与base元素的和

func addBase(base int) func(int) int {
	return func(num int) int {
		return base + num
	}

}


PS E:\go-phase-two\go-course> go run test.go
3 11
7 15
12 20
```


## 值类型和引用类型

值类型和引用类型的差异在于赋值同类型新变量后，对新变量进行修改是否能够影响原来的变量，若不能则为值类型，若能则为引用类型

```go
package main

import "fmt"

func main() {

	name, age, heigh, isBoy := "Silence", 30, 1.58, false // 定义字符串，数值，布尔类型
	pointer := new(int)                                   // 定义指针类型
	scores := [...]int{1, 2, 3}                           // 定义数组类型
	names := make([]string, 1, 3)                         // 定义切片类型
	user := make(map[int]string)                          // 定义映射类型

	name2, age2, heigh2, isBoy2, pointer2, scores2, names2, user2 := name, age, heigh, isBoy, pointer, scores, names, user

	name2 = "kk"
	age2 = 31
	heigh2 = 1.30
	isBoy2 = true
	scores2[0] = 0
	pointer2 = &age
	names2[0] = "kk"
	user2[1] = "kk"

	fmt.Println(name, age, heigh, isBoy, scores, pointer, names, user)
	fmt.Println(name2, age2, heigh2, isBoy2, scores2, pointer2, names2, user2)
}

PS E:\go-phase-two\go-course> go run test.go
Silence 30 1.58 false [1 2 3] 0xc0000100a8 [kk] map[1:kk]
kk 31 1.3 true [0 2 3] 0xc0000100a0 [kk] map[1:kk]

```


- 值类型： 数值， 布尔， 字符串， 指针， 数组， 结构体
- 引用类型： 切片， 映射， 接口等

针对值类型可以借助指针修改原值

```go
package main

import "fmt"

func main() {

	name, age, heigh, isBoy := "Silence", 30, 1.58, false // 定义字符串，数值，布尔类型
	pointer := new(int)                                   // 定义指针类型
	scores := [...]int{1, 2, 3}                           // 定义数组类型
	names := make([]string, 1, 3)                         // 定义切片类型
	user := make(map[int]string)                          // 定义映射类型

	name2, age2, heigh2, isBoy2, pointer2, scores2, names2, user2 := name, age, heigh, isBoy, pointer, scores, names, user

	name2 = "kk"
	age2 = 31
	heigh2 = 1.30
	isBoy2 = true
	scores2[0] = 0
	pointer2 = &age
	names2[0] = "kk"
	user2[1] = "kk"

	fmt.Println(name, age, heigh, isBoy, scores, pointer, names, user)
	fmt.Println(name2, age2, heigh2, isBoy2, scores2, pointer2, names2, user2)

	fmt.Printf("%p, %p, %p, %p, %p, %p, %p %p\n", &name, &age, &heigh, &isBoy, &scores, &pointer, &names, &user)
	fmt.Printf("%p, %p, %p, %p, %p, %p, %p %p\n", &name2, &age2, &heigh2, &isBoy2, &scores2, &pointer2, &names2, &user2)
}


package main

import "fmt"

func main() {

	name, age, heigh, isBoy := "Silence", 30, 1.58, false // 定义字符串，数值，布尔类型
	pointer := new(int)                                   // 定义指针类型
	scores := [...]int{1, 2, 3}                           // 定义数组类型
	names := make([]string, 1, 3)                         // 定义切片类型
	user := make(map[int]string)                          // 定义映射类型

	name2, age2, heigh2, isBoy2, pointer2, scores2, names2, user2 := name, age, heigh, isBoy, pointer, scores, names, user

	name2 = "kk"
	age2 = 31
	heigh2 = 1.30
	isBoy2 = true
	scores2[0] = 0
	pointer2 = &age
	names2[0] = "kk"
	user2[1] = "kk"

	fmt.Println(name, age, heigh, isBoy, scores, pointer, names, user)
	fmt.Println(name2, age2, heigh2, isBoy2, scores2, pointer2, names2, user2)

	fmt.Printf("%p, %p, %p, %p, %p, %p, %p %p\n", &name, &age, &heigh, &isBoy, &scores, &pointer, &names, &user)
	fmt.Printf("%p, %p, %p, %p, %p, %p, %p %p\n", &name2, &age2, &heigh2, &isBoy2, &scores2, &pointer2, &names2, &user2)
}

```


针对值类型和引用类型在赋值后新旧变量的地址并不相同，只是引用类型在底层共享数据结构(其中包含指针类型元素)

## 值传递和引用传递

### 值传递 

在 Go 语言中参数传递默认均为值传递（形参为实参变量的副本），对于引用类型数据因其底层共享数据结构，所以在函数内可对引用类型数据修改从而影响函数外的原变量信息


```go
package main

import "fmt"

func main() {
	e1, e2 := "kk", []string{"kk", "silence"}

	// 值传递
	// 在函数内修改值类型
	fmt.Printf("e1: %p %v\n", &e1, e1)
	func(e string) {
		fmt.Printf("e: %p %v\n", &e, e)
		e = "silence"
	}(e1)

	// 在函数内修改引用类型
	fmt.Printf("e2: %p %v\n", &e2, e2)
	func (e []string) {
		fmt.Printf("e: %p %v\n", &e, e)
		e[1] = "woniu"
	}(e2)

	fmt.Println(e1, e2)
}


PS E:\go-phase-two\go-course> go run test.go
e1: 0xc0000461f0 kk
e: 0xc000046210 kk
e2: 0xc0000044a0 [kk silence]
e: 0xc000004520 [kk silence]
kk [kk woniu]
```


### 应用传递

可以通过将变量的地址通过指定类型传递给函数，此时可通过指针对函数外的原变量进行修改。

```go
package main

import "fmt"

func main() {
	e1, e2 := "kk", []string{"kk", "silence"}

	// 值传递
	// 在函数内修改值类型
	fmt.Printf("e1: %p %v\n", &e1, e1)
	func(e *string) {
		fmt.Printf("e: %p %v\n", &e, *e)
		*e = "silence"
	}(&e1)

	// 在函数内修改引用类型
	fmt.Printf("e2: %p %v\n", &e2, e2)
	func(e *[]string) {
		fmt.Printf("e: %p %v\n", &e, *e)
		(*e)[1] = "woniu"
	}(&e2)

	fmt.Println(e1, e2)

}


PS E:\go-phase-two\go-course> go run test.go
e1: 0xc00005a1e0 kk
e: 0xc00009c020 kk
e2: 0xc000068440 [kk silence]
e: 0xc00009c028 [kk silence]
silence [kk woniu]
```


## 错误处理 
### error 接口 

Go 语言通过 error 接口实现错误处理的标准模式，通过使用函数返回值列表中的最后一个值返回错误信息，将错误的处理交由程序员主动进行处理。

```go

// 定义除法函数，若除数为0则使用error返回错误信息
func division(n1, n2 int) (int, error) {
	if n2 == 0 {
		return 0, errors.New("除数为0")

	}
	return n1 / n2, nil
}

// 处理函数返回的错误
for _, v := range [...]int{0, 3} {
	if r, err := division(6, v); err != nil {
		fmt.Println(err)

	} else {
		fmt.Println(r)
	}
}
```

error 接口的初始化方法

1. 通过errors 包的New 方法创建

2. 通过fmt.Errorf 方法创建 
```go
package main

import (
	"errors"
	"fmt"
)

func main() {

	err1, err2 := errors.New("error: 1"), fmt.Errorf("error: %d", 2)
	fmt.Printf("%T, %T, %v, %v", err1, err2, err1, err2)
}
```

### defer 

defer 关键字用户声明函数，不论函数是否发生错误都在函数执行最后执行(return 之前)，若使用 defer 声明多个函数，则按照声明的顺序，先声明后执行（堆） ，常用来做资源释放，记录日志等工作

```go
package main

import (
	"fmt"
)

func main() {

	defer func() {
		fmt.Println("defer 01")
	}()

	defer func() {
		fmt.Println("defer 02")
	}()

	defer func() {
		fmt.Println("defer 03")
	}()

	fmt.Println("over")
}

```



### panic 与 recover 函数 

go 语言提供panic 与 recover 函数用于处理运行时错误， 当调用panic 抛出错误，中断原有的控制流程，常用于不可修复性错误。recover 函数用于中止错误处理流程，仅在defer 语句中的函数中有效，用于截取错误处理流程，recover只能捕获到最后一个错误。

1. panic 
```go
package main

import (
	"fmt"
)

func main() {
	defer func() {
		fmt.Println("defer 01")

	}()

	panic("error 00")
}

defer 01
panic: error 00

goroutine 1 [running]:
main.main()
        E:/go-phase-two/go-course/test.go:13 +0x66
exit status 2
```

2. recover 

```go
package main

import (
	"fmt"
)

func init() {
	fmt.Println("init")
}

func main() {
	fmt.Println("main")
}

// 当未发生panic则recover函数得到的结果为nil

func success() {
	defer func() {
		fmt.Println(recover())
	}()
	fmt.Println("success")
}

// 当未发生panic 则recover 函数的得到的结果为panic传递的参数
func failure() {
	defer func() {
		fmt.Println(recover())
	}()
	fmt.Println("failure")
	panic("error")
}

// recover 只能获取最后依次的panic信息

func failure2() {
	defer func() {
		fmt.Println(recover())
	}()

	defer func() {
		fmt.Println("failure 02")
		panic("error 02")
	}()

	fmt.Println("failure")
	panic("error")
}



PS E:\go-phase-two\go-course> go run test.go
init
main
```



# 包 

包是函数和数据的集合，将有相关特性的函数和数据放在统一的文件/目录进行管理，每个包都可以作为独立的单元维护并提供给其他项目使用。

## 声明 

go 源文件都需要在开头使用package 声明所在包，包名告知编译器哪些包的源代码用于编译库文件，其次包名用于限制包内成员对外的可见性，最后包名用于在包外对外公开成员的访问。

包名使用简短的小写字母，常与所在目录名保持一致，一个包中可以由多个 Go 源文件，但必须使用相同包名


![](https://raw.githubusercontent.com/PassZhang/passzhang.github.io/images-picgo/20200910155207.png)

声明两个包，路径分别为 gpkgname/pkg01 和 gpkgname/pkg02

## 导入&调用 

在使用包时需要使用 import 将包按路径名导入，再通过包名调用成员

可通过 import 每行导入一个包，也可使用括号包含所有包并使用一个 import 导入

![](https://raw.githubusercontent.com/PassZhang/passzhang.github.io/images-picgo/20200910155513.png)

工作目录结构说明：
- bin: 用于放置发布的二进制程序
- pkg： 用于放置发布的库文件
- src： 用于放置源代码


**运行:**
a) 将 E:\go-phase-two\go-course\good_print 目录添加到 GOPATH 环境变量中
```go
# bash是 powershell

$env:GOPATH="E:\go-phase-two\go-course\good_print"
$env:GOPATH
ls env:
```

b) 编译&运行
- 使用 go build 编译二进制文件
命令：go build gpkgmain
说明：编译路径 gpkgmain 下的包，main 包，则在当前目录产生以 pkgmain 命名的二进制程序

- 使用 go run 运行二进制文件
命令：go run gpkgmain

- 使用 go install 编译并发布二进制文件
命令：go install gpkgmain
说明：编译并发布路径 gpkgmain 下的包，main 包，则在将编译后的以 pkgmain 命名的二进制程序拷贝到 bin 目录

- 使用 go install 编译发布库文件
命令：go install gpkgname/pkg01
说明：编译并发布路径 gpkgname/pkg01 下的包，非 main 包，则在将编译的以包名命名的库文件拷贝到 pkg/GOOS_GOARCH 目录下 ⚫ 使用 go install 编译发布所有二进制和库文件

命令：go install ./…
说明：编译并发布当前路径下的所有二进制程序和库文件
注意：Go 语言不允许交叉导入包

## 导入形式

1) 绝对路径导入
在 GOPATH 目录中查找包
示例：
	- import "fmt"
	- import "gpkgname/pkg01"
	
2) 相对路径导入
在当前文件所在的目录查找
示例：import "./gpkgname/pkg02"

3) 点导入
在调用点导入包中的成员时可以直接使用成员名称进行调用（省略包名）

```go
package main

import . "fmt"


func main()  {
	Println("Hello World")
}
```

4) 别名导入
当导入不同路径的相同包名时，可以别名导入为包重命名，避免冲突

```go
package main

import f "fmt"


func main()  {
	f.Println("Hello World")
}
```

5) 下划线导入
Go 不允许包导入但未使用，在某些情况下需要初始化包，使用空白符作为别名进行导入，从而使得包中的初始化函数可以执行

```go
package main

import _ "fmt"


func main()  {
	Println("Hello World")
}
```


## 成员可见性

Go 语言使用名称首字母大小写来判断对象(常量、变量、函数、类型、结构体、方法等)的访问权限，首字母大写标识包外可见(公开的)，否者仅包内可访问(内部的)

```go
package pkg01

// public
var Name string = "good print pkg01"


//private
var version string = "1.0.0"
```

## main包和main函数

main 包用于声明告知编译器将包编译为二进制可执行文件，在main包中的main函数是程序的入口，无返回值，无参数。

## init 函数 

init 函数时初始化包使用，无返回值，无参数。 建议每个包只定义一个。 init 函数在import 包时自动被调用（const-> var -> init）

```go
E:\go-phase-two\go-course\good_print\src\gpkgname\pkg01\module.go

package pkg01

import "fmt"

// public
var Name string = "good print pkg01"

//private
var version string = "1.0.0"

func init() {
	fmt.Println("Version: ", version)
}

```

```go
E:\go-phase-two\go-course\good_print\src\gpkgmain\main.go

package main

import (
	"fmt"
	"gpkgname/pkg01" // 导入包pkg01， 路径gpkgname/pkg01
	"gpkgname/pkg02" // 导入包pkg02， 路径gpkgname/pkg02
)

func init() {
	fmt.Println("main")
}

func main() {
	fmt.Println("gpkgmain")
	fmt.Println(pkg01.Name) // 调用pkg01包中的成员Name
	fmt.Println(pkg02.Name) // 调用pkg02包中的成员Name
}


PS E:\go-phase-two\go-course\good_print> go run gpkgmain
Version:  1.0.0
main
gpkgmain
good print pkg01
good print pkg02
```


## go包管理

### 介绍

Go1.11 版本提供 Go modules 机制对包进行管理，同时保留 GOPATH 和 vendor 机制，使用临时环境变量 

GO111MODULE 进行控制，GO111MODULE 有三个可选值：
a) 当 GO111MODULE 为 off 时，构建项目始终在 GOPATH 和 vendor 目录搜索目标程序依赖包

b) 当 GO111MODULE 为 on 时，构建项目则始终使用 Go modules 机制，在 GOPATH/pkg/mod目录搜索目标程序依赖包

c) 当 GO111MODULE 为 auto(默认)时,当构建源代码不在 GOPATH/src 的子目录且包含go.mod 文件，则使用 Go modules 机制，否则使用 GOPATH 和 vendor 机制

### GOPATH+vendor机制

a) vendor
将项目依赖包拷贝到项目下的 vendor 目录，在编译时使用项目下 vendor 目录中的包进行编译
解决问题：
- 依赖外部包过多，在使用第三方包时需要使用 go get 进行下载
- 第三方包在 go get 下载后不能保证开发和编译时版本的兼容性


b) 包搜索顺序
- 在当前包下的 vendor 目录查找
- 向上级目录查找，直到 GOPATH/src/vendor 目录
- 在 GOPATH 目录查找
- 在 GOROOT 目录查找


c) 第三方包
可以借助 go get 工具下载和安装第三方包及其依赖，需要安装与第三方包匹配的代码
管理工具，比如 git、svn 等


![](https://raw.githubusercontent.com/PassZhang/passzhang.github.io/images-picgo/20200910192446.png

go get 常用参数
-  -d：仅下载依赖包
-  -u：更新包并安装
-  -x：打印执行的命令
-  -v：打印构建的包
-  -insecure：允许使用 http 协议下载包


第三方包查找地址：
-  https://godoc.org
-  https://gowalker.org


### go module 机制

a) 优势：
- 不用设置 GOPATH，代码可任意放置
- 自动下载依赖管理
- 版本控制
- 不允许使用相对导入
- replace 机制

b) 初始化模块
命令：go mod init modname 

```go
go mod init github.com/PassZhang/go-course/go_mod_test
```




c) 当前模块下的包
对于当前模块下的包导入时需要使用 modname+packagename

```go

package main

import (
	"fmt"

	"github.com/PassZhang/go-course/go_mod_test/src/gpkgname/pkg01" // 导入包pkg01， 路径gpkgname/pkg01
	"github.com/PassZhang/go-course/go_mod_test/src/gpkgname/pkg02" // 导入包pkg02， 路径gpkgname/pkg02
)

func main() {
	fmt.Println("go_mod")
	fmt.Println(pkg01.Name) // 调用pkg01包中的成员Name
	fmt.Println(pkg02.Name) // 调用pkg02包中的成员Name
}

```


第三方包
在使用 go mod tidy、go build、go test、go list 命令会自动将第三方依赖包写入到go.mod 文件中同时下载第三方依赖包到 GOPATH/pkg/mod/cache 目录，并在当前模块目录生成一个构建状态跟踪文件 go.sum,文件中记录当前 module 所有的顶层和间接依赖，以及这些依赖的校验和.

```go
package main

import "github.com/astaxie/beego"

func main() {
	beego.Run()
}


PS E:\go-phase-two\go-course\go_beego_web> .\web.exe
2020/09/10 20:07:31.841 [I]  http server Running on http://:8080

```

e) 常用命令
- go mod tidy:整理依赖模块（添加新增的，删除未使用的）
- go mod vendor: 将依赖模块拷贝到模块中的 vendor 目录
- go build: 编译当前模块
- go build ./...: 编译当前目录下的所有模块
- go build -mod=vendor:使用当前模块下的 vendor 目录中的包进行编译
- go mod download: 仅下载第三方模块
- go mod grapha: 打印所有第三方模块
- go list -m -json all：显示所有模块信息
- go mod edit: 修改 go.mod 文件
-require=pakcage@version
-replace=old_package@version=new_package@version
可以使用-replace 功能将包替换为本地包，实现相对导入


### 标准包 
Go 提供了大量标准包，可查看：https://golang.google.cn/pkg

#### godoc 工具 
使用 godoc 命令可以在本地启动 golang 网站，用于本地查看帮助手册

#### 帮助 
a) go list std：查看所有标准包

b) go doc packagename：查看包的帮助信息

c) go doc packagename.element：查看包内成员帮助信息

# 结构体 

## 自定义类型

在 go 语言中使用 type 声明一种新的类型，语法格式为：
```go
type TypeName Formatter
```
Format 可以时任意内置类型、函数签名、结构体、接口

```go
package main

import (
	"fmt"
)

func main() {
	// 使用自定义结构
	var counter Counter

	fmt.Printf("%T:%v", counter, counter)
	counter++
	counter++

	fmt.Println(counter)

	// 使用自定义结构体User
	var me User = make(User)
	me["name"] = "kk"
	me["addr"] = "西安"

	fmt.Printf("%T: %#v\n", me, me)

	names := []string{"kk", "ww", "ada"}
	// 将函数column传递给printResult的pf(Callback类型)参数
	printResult(column, names...)
}

// 自定义int 类型
type Counter int

// 自定义map[string] string 类型
type User map[string]string

// 自定义函数类型
type Collback func(...string)

// 定义接受自定义类型（函数类型）为参数的函数

func printResult(pf Collback, list ...string) {
	pf(list...)
}

func column(list ...string) {
	for _, e := range list {
		fmt.Println(e)
	}
	fmt.Println()
}


PS E:\go-phase-two\go-course> go run .\1test.go
main.Counter:02
main.User: main.User{"addr":"西安", "name":"kk"}
kk
ww
ada
```

## 定义

结构体定义使用struct 标识，需要指定其包含的属性(名和类型)， 在定义结构体时可以为结构体指定结构体名（命名结构体），用于后续声明结构体变量使用。

```go
package main

import (
	"fmt"
)

func main() {
	var kk Author
	fmt.Printf("%T: %#v\n", kk, kk)

	var pkk *Author
	fmt.Printf("%T: %#v\n", pkk, pkk)
}

// 定义结构体
type Author struct {
	id   int
	name string
	addr string
	tel  string
	desc string
}


PS E:\go-phase-two\go-course> go run .\1test.go
main.Author: main.Author{id:0, name:"", addr:"", tel:"", desc:""}
*main.Author: (*main.Author)(nil)

```


## 声明 
声明结构体变量只需要定义变量类型为结构体名，变量中的每个属性被初始化为对应类型的零值。也可以声明结构体指针变量，此时变量被初始化为nil。


```go
package main

import (
	"fmt"
)

// 定义结构体
type Author struct {
	id   int
	name string
	addr string
	tel  string
	desc string
}


func main() {
	var kk Author
	fmt.Printf("%T: %#v\n", kk, kk)

	var pkk *Author
	fmt.Printf("%T: %#v\n", pkk, pkk)
}



PS E:\go-phase-two\go-course> go run .\1test.go
main.Author: main.Author{id:0, name:"", addr:"", tel:"", desc:""}
*main.Author: (*main.Author)(nil)

```


## 初始化 

使用结构体创建的变量叫做对应结构体的实例或者对象
1. 使用结构体零值初始化结构体值对象
   ```go
   var kk Author = Author
   ```

2. 使用结构体字面量初始化结构体对象

```go
package main

import "fmt"

// type person struct {
// 	name string
// 	age  int
// }

type Author struct {
	id   int
	name string
	addr string
	tel  string
	desc string
}

func main() {

	// 按照结构体属性定义顺序初始化变量
	var kk01 Author = Author{
		1001,
		"kk",
		"西安",
		"101000111",
		"天天学习", // 结尾逗号不能省略
	}

	fmt.Printf("%T, %#v\n", kk01, kk01)

	// 只初始化一部分属性或不按照定义顺序初始化使用命名方式
	kk02 := Author{name: "kk2", id: 1002}
	fmt.Printf("%T, %#v\n", kk01, kk01)
	fmt.Printf("%T, %#v\n", kk02, kk02)

	// 使用new函数进行初始化结构体指针
	var kk03 *Author = new(Author)
	fmt.Printf("%T: %#v, %#v\n", kk03, kk03, *kk03)

	// // 用type 关键字定义一个结构体

	// // 实例化一个结构体
	// // 第一种情况
	// p1 := person{}
	// p1.name = "小麦"
	// p1.age = 666
	// fmt.Println(p1.name)
	// fmt.Println(p1.age)

	// // 第二种情况
	// p2 := person{
	// 	name: "小麦",
	// 	age:  666,
	// }
	// // %#v 打印详细信息
	// fmt.Printf("%#v\n", p2)
}

PS E:\go-phase-two\go-course> go run .\1test.go
main.Author, main.Author{id:1001, name:"kk", addr:"西安", tel:"101000111", desc:"天天学习"}
main.Author, main.Author{id:1001, name:"kk", addr:"西安", tel:"101000111", desc:"天天学习"}
main.Author, main.Author{id:1002, name:"kk2", addr:"", tel:"", desc:""}
*main.Author: &main.Author{id:0, name:"", addr:"", tel:"", desc:""}, main.Author{id:0, name:"", addr:"", 
tel:"", desc:""}


```

3. 使用new函数进行初始化结构体指针对象

```go
package main

import "fmt"

// type person struct {
// 	name string
// 	age  int
// }

type Author struct {
	id   int
	name string
	addr string
	tel  string
	desc string
}

func main() {

	// 按照结构体属性定义顺序初始化变量
	var kk01 Author = Author{
		1001,
		"kk",
		"西安",
		"101000111",
		"天天学习", // 结尾逗号不能省略
	}

	fmt.Printf("%T, %#v\n", kk01, kk01)

	// 只初始化一部分属性或不按照定义顺序初始化使用命名方式
	kk02 := Author{name: "kk2", id: 1002}
	fmt.Printf("%T, %#v\n", kk01, kk01)
	fmt.Printf("%T, %#v\n", kk02, kk02)

	// 使用new函数进行初始化结构体指针
	var kk03 *Author = new(Author)
	fmt.Printf("%T: %#v, %#v\n", kk03, kk03, *kk03)

	// // 用type 关键字定义一个结构体

	// // 实例化一个结构体
	// // 第一种情况
	// p1 := person{}
	// p1.name = "小麦"
	// p1.age = 666
	// fmt.Println(p1.name)
	// fmt.Println(p1.age)

	// // 第二种情况
	// p2 := person{
	// 	name: "小麦",
	// 	age:  666,
	// }
	// // %#v 打印详细信息
	// fmt.Printf("%#v\n", p2)
}

PS E:\go-phase-two\go-course> go run .\1test.go
main.Author, main.Author{id:1001, name:"kk", addr:"西安", tel:"101000111", desc:"天天学习"}
main.Author, main.Author{id:1001, name:"kk", addr:"西安", tel:"101000111", desc:"天天学习"}
main.Author, main.Author{id:1002, name:"kk2", addr:"", tel:"", desc:""}
*main.Author: &main.Author{id:0, name:"", addr:"", tel:"", desc:""}, main.Author{id:0, name:"", addr:"", 
tel:"", desc:""}


```

4. 使用结构体字面量初始化结构体指针对象

```go
package main

import "fmt"

type Author struct {
	id   int
	name string
	addr string
	tel  string
	desc string
}

// type person struct {
// 	name string
// 	age  int
// }

func main() {

	// 按照结构体属性定义顺序初始化变量
	var kk01 Author = Author{
		1001,
		"kk",
		"西安",
		"101000111",
		"天天学习", // 结尾逗号不能省略
	}

	fmt.Printf("%T, %#v\n", kk01, kk01)

	// 只初始化一部分属性或不按照定义顺序初始化使用命名方式
	kk02 := Author{name: "kk2", id: 1002}
	fmt.Printf("%T, %#v\n", kk01, kk01)
	fmt.Printf("%T, %#v\n", kk02, kk02)

	// 使用new函数进行初始化结构体指针
	var kk03 *Author = new(Author)
	fmt.Printf("%T: %#v, %#v\n", kk03, kk03, *kk03)

	// 使用结构体字面量初始化结构体指针对象
	kk04 := &Author{id: 1004, name: "kk4"}
	fmt.Printf("%T: %#v, %#v\n", kk04, kk04, *kk04)

	// // 用type 关键字定义一个结构体

	// // 实例化一个结构体
	// // 第一种情况
	// p1 := person{}
	// p1.name = "小麦"
	// p1.age = 666
	// fmt.Println(p1.name)
	// fmt.Println(p1.age)

	// // 第二种情况
	// p2 := person{
	// 	name: "小麦",
	// 	age:  666,
	// }
	// // %#v 打印详细信息
	// fmt.Printf("%#v\n", p2)
}


PS E:\go-phase-two\go-course> go run .\1test.go
main.Author, main.Author{id:1001, name:"kk", addr:"西安", tel:"101000111", desc:"天天学习"}
main.Author, main.Author{id:1001, name:"kk", addr:"西安", tel:"101000111", desc:"天天学习"}
main.Author, main.Author{id:1002, name:"kk2", addr:"", tel:"", desc:""}
*main.Author: &main.Author{id:0, name:"", addr:"", tel:"", desc:""}, main.Author{id:0, name:"", addr:"", 
tel:"", desc:""}
*main.Author: &main.Author{id:1001, name:"kk4", addr:"", tel:"", desc:""}, main.Author{id:1001, name:"kk4", addr:"", tel:"", desc:""}

```

## New函数 

Go 语言中常定义N(n)ew+结构体名命名的函数用于创建对应的结构体值对象或指针对象

```go
package main

import "fmt"

type Author struct {
	id   int
	name string
	addr string
	tel  string
	desc string
}

func main() {

	// 按照结构体属性定义顺序初始化变量
	var kk01 Author = Author{
		1001,
		"kk",
		"西安",
		"101000111",
		"天天学习", // 结尾逗号不能省略
	}

	fmt.Printf("%T, %#v\n", kk01, kk01)

	// 只初始化一部分属性或不按照定义顺序初始化使用命名方式
	kk02 := Author{name: "kk2", id: 1002}
	fmt.Printf("%T, %#v\n", kk01, kk01)
	fmt.Printf("%T, %#v\n", kk02, kk02)

	// 使用new函数进行初始化结构体指针
	var kk03 *Author = new(Author)
	fmt.Printf("%T: %#v, %#v\n", kk03, kk03, *kk03)

	// 使用结构体字面量初始化结构体指针对象
	kk04 := &Author{id: 1004, name: "kk4"}
	fmt.Printf("%T: %#v, %#v\n", kk04, kk04, *kk04)

	kk05 := NewAuthor(1004, "kk05", "西安市", "15300000000", "看山是山")
	fmt.Printf("%T: %#v\n", kk05, kk05)
}

func NewAuthor(id int, name, addr, tel, desc string) *Author {
	return &Author{id, name, addr, tel, desc}
}


PS E:\go-phase-two\go-course> go run .\1test.go
main.Author, main.Author{id:1001, name:"kk", addr:"西安", tel:"101000111", desc:"天天学习"}
main.Author, main.Author{id:1001, name:"kk", addr:"西安", tel:"101000111", desc:"天天学习"}
main.Author, main.Author{id:1002, name:"kk2", addr:"", tel:"", desc:""}
*main.Author: &main.Author{id:0, name:"", addr:"", tel:"", desc:""}, main.Author{id:0, name:"", addr:"", 
tel:"", desc:""}
*main.Author: &main.Author{id:1004, name:"kk4", addr:"", tel:"", desc:""}, main.Author{id:1004, name:"kk4", addr:"", tel:"", desc:""}
*main.Author: &main.Author{id:1004, name:"kk05", addr:"西安市", tel:"15300000000", desc:"看山是山"}  
```


## 属性的访问和修改

通过结构体对象名/结构体指针对象. 属性名的方式来访问和修改对象的属性值 

```go

package main

import "fmt"

type Author struct {
	id   int
	name string
	addr string
	tel  string
	desc string
}

func main() {

	var kk Author = Author{
		1000,
		"kk",
		"北京",
		"101000111",
		"天天开心", // 结尾逗号不能省略
	}

	// 按照结构体属性定义顺序初始化变量
	var kk01 Author = Author{
		1001,
		"kk1",
		"西安",
		"101000111",
		"天天学习", // 结尾逗号不能省略
	}

	fmt.Printf("%T, %#v\n", kk01, kk01)

	// 只初始化一部分属性或不按照定义顺序初始化使用命名方式
	kk02 := Author{name: "kk2", id: 1002}
	fmt.Printf("%T, %#v\n", kk01, kk01)
	fmt.Printf("%T, %#v\n", kk02, kk02)

	// 使用new函数进行初始化结构体指针
	var kk03 *Author = new(Author)
	fmt.Printf("%T: %#v, %#v\n", kk03, kk03, *kk03)

	// 使用结构体字面量初始化结构体指针对象
	kk04 := &Author{id: 1004, name: "kk4"}
	fmt.Printf("%T: %#v, %#v\n", kk04, kk04, *kk04)

	kk05 := NewAuthor(1004, "kk05", "西安市", "15300000000", "看山是山")
	fmt.Printf("%T: %#v\n", kk05, kk05)

	// 使用结构体对象访问和修改属性
	fmt.Printf("%d: %q\n", kk.id, kk.name)
	fmt.Printf("%d: %q\n", kk01.id, kk01.name)
	fmt.Printf("%d: %q\n", kk02.id, kk02.name)

	kk02.addr = "西安市"
	fmt.Println(kk02)

	// 使用结构体指针变量访问属性
	fmt.Println((*kk03).name)
	fmt.Println(kk04.name) // 使用指针直接访问

	// 使用结构体指针变量修改属性
	(*kk03).name = "kk3"
	kk04.addr = "上海市" // 使用指针直接修改
	fmt.Println(kk03, kk04)
}

func NewAuthor(id int, name, addr, tel, desc string) *Author {
	return &Author{id, name, addr, tel, desc}
}


PS E:\go-phase-two\go-course> go run .\1test.go
main.Author, main.Author{id:1001, name:"kk1", addr:"西安", tel:"101000111", desc:"天天学习"}
main.Author, main.Author{id:1001, name:"kk1", addr:"西安", tel:"101000111", desc:"天天学习"}
main.Author, main.Author{id:1002, name:"kk2", addr:"", tel:"", desc:""}
*main.Author: &main.Author{id:0, name:"", addr:"", tel:"", desc:""}, main.Author{id:0, name:"", addr:"", 
*main.Author: &main.Author{id:1004, name:"kk4", addr:"", tel:"", desc:""}, main.Author{id:1004, name:"kk4", addr:"", tel:"", desc:""}
*main.Author: &main.Author{id:1004, name:"kk05", addr:"西安市", tel:"15300000000", desc:"看山是山"}      
1000: "kk"
1001: "kk1"
1002: "kk2"
{1002 kk2 西安市  }

kk4
&{0 kk3   } &{1004 kk4 上海市  }
```

可以通过结构体指针对象的点操作直接对对象的属性值进行修改和访问。


## 匿名结构体

在定义变量时将类型指定为结构体的结构，此时叫匿名结构体。匿名结构体常用于初始化一次结构体变量的场景，例如项目配置

```go
package main

import "fmt"

func main() {

	var coordinate struct {
		longitude float64
		latitude  float64
	}

	me := struct {
		name string
		age  int
		addr string
	}{"silence", 30, "西安"}

	// 匿名结构体数组
	studs := []struct {
		name string
		addr string
	}{{"silence", "西安"}, {"ww", "北京"}}

	fmt.Println(coordinate, coordinate.longitude, coordinate.latitude)

	fmt.Println(me, me.name, me.age, me.addr)

	fmt.Println(studs)

	for i, stud := range studs {
		fmt.Println(i, stud, stud.name, stud.addr)
	}
}

PS E:\go-phase-two\go-course> go run .\1test.go
{0 0} 0 0
{silence 30 西安} silence 30 西安
[{silence 西安} {ww 北京}]
0 {silence 西安} silence 西安
1 {ww 北京} ww 北京
```


## 命名嵌入

结构体命名嵌入是指结构体中的属性对应的类型也是结构体

### 定义 

```go
package main

import "fmt"

func main() {

	// 定义收货人结构体（区域，街道，门牌号）
	type Address struct {
		region string
		street string
		no     string
	}

	// 定义收货人结构体（名字，电话，地址）
	type User struct {
		name string
		tel  string
		addr Address
	}

	// 声明& 初始化组合对象
	var u1 User

	fmt.Printf("%T, %#v\n", u1, u1)

	var a1 Address = Address{
		"陕西省西安市",
		"开心路",
		"001",
	}

	var u2 User = User{
		"kk",
		"15000000000",
		a1,
	}

	u3 := User{
		name: "ww",
		tel:  "100000000001",
		addr: Address{
			"北京市",
			"海淀路",
			"001",
		},
	}

	fmt.Println(u1, u2, u3)

	// 属性的访问
	fmt.Println(u2.addr)
	fmt.Println(u2.addr.no)

	// 属性的修改
	u2.addr.no = "002"
	fmt.Println(u2)
}

PS E:\go-phase-two\go-course> go run .\1test.go
main.User, main.User{name:"", tel:"", addr:main.Address{region:"", street:"", no:""}}
{  {  }} {kk 15000000000 {陕西省西安市 开心路 001}} {ww 100000000001 {北京市 海淀路 001}}
{陕西省西安市 开心路 001}
001
{kk 15000000000 {陕西省西安市 开心路 002}}
```
### 声明和初始化


```go
package main

import "fmt"

func main() {

	// 定义收货人结构体（区域，街道，门牌号）
	type Address struct {
		region string
		street string
		no     string
	}

	// 定义收货人结构体（名字，电话，地址）
	type User struct {
		name string
		tel  string
		addr Address
	}

	// 声明& 初始化组合对象
	var u1 User

	fmt.Printf("%T, %#v\n", u1, u1)

	var a1 Address = Address{
		"陕西省西安市",
		"开心路",
		"001",
	}

	var u2 User = User{
		"kk",
		"15000000000",
		a1,
	}

	u3 := User{
		name: "ww",
		tel:  "100000000001",
		addr: Address{
			"北京市",
			"海淀路",
			"001",
		},
	}

	fmt.Println(u1, u2, u3)

	// 属性的访问
	fmt.Println(u2.addr)
	fmt.Println(u2.addr.no)

	// 属性的修改
	u2.addr.no = "002"
	fmt.Println(u2)
}

PS E:\go-phase-two\go-course> go run .\1test.go
main.User, main.User{name:"", tel:"", addr:main.Address{region:"", street:"", no:""}}
{  {  }} {kk 15000000000 {陕西省西安市 开心路 001}} {ww 100000000001 {北京市 海淀路 001}}
{陕西省西安市 开心路 001}
001
{kk 15000000000 {陕西省西安市 开心路 002}}
```

### 属性的访问和修改


```go
package main

import "fmt"

func main() {

	// 定义收货人结构体（区域，街道，门牌号）
	type Address struct {
		region string
		street string
		no     string
	}

	// 定义收货人结构体（名字，电话，地址）
	type User struct {
		name string
		tel  string
		addr Address
	}

	// 声明& 初始化组合对象
	var u1 User

	fmt.Printf("%T, %#v\n", u1, u1)

	var a1 Address = Address{
		"陕西省西安市",
		"开心路",
		"001",
	}

	var u2 User = User{
		"kk",
		"15000000000",
		a1,
	}

	u3 := User{
		name: "ww",
		tel:  "100000000001",
		addr: Address{
			"北京市",
			"海淀路",
			"001",
		},
	}

	fmt.Println(u1, u2, u3)

	// 属性的访问
	fmt.Println(u2.addr)
	fmt.Println(u2.addr.no)

	// 属性的修改
	u2.addr.no = "002"
	fmt.Println(u2)
}

PS E:\go-phase-two\go-course> go run .\1test.go
main.User, main.User{name:"", tel:"", addr:main.Address{region:"", street:"", no:""}}
{  {  }} {kk 15000000000 {陕西省西安市 开心路 001}} {ww 100000000001 {北京市 海淀路 001}}
{陕西省西安市 开心路 001}
001
{kk 15000000000 {陕西省西安市 开心路 002}}
```


## 匿名嵌入

结构体匿名嵌入是指将已定义的结构体名直接声明在新的结构体中，从而实现对以后已有类型的扩展和修改。

### 定义
```go
package main

import (
	"fmt"
)

// 定义地址结构体（区域，街道，门牌号）
type Address struct {
	region string
	street string
	no     string
}

// 定义收货人结构体（名字，电话，地址）
type User struct {
	name string
	tel  string
	addr Address
}

// 定义员工，匿名嵌入user结构体
type Employee struct {
	User
	salary float64
	title  string
}

func main() {

	var ekk Employee
	fmt.Printf("%T: %#v\n", ekk, ekk)

	// 使用已有变量初始化

	ekk02 := Employee{
		User{
			"kk",
			"152000022222",
			Address{
				"陕西省西安市",
				"锦业路",
				"002",
			},
		},
		22000,
		"安全开发工程师",
	}

	ekk03 := Employee{
		User{
			"kk",
			"153000033333",
			Address{
				"陕西省西安市",
				"锦业路",
				"003",
			},
		},
		23000,
		"架构工程师",
	}


	fmt.Println(ekk02)
	fmt.Println(ekk03)
}


PS E:\go-phase-two\go-course> go run E:\go-phase-two\go-course\1test.go
main.Employee: main.Employee{User:main.User{name:"", tel:"", addr:main.Address{region:"", street:"", no:""}}, salary:0, title:""}
{{kk 152000022222 {陕西省西安市 锦业路 002}} 22000 安全开发工程师}
{{kk 153000033333 {陕西省西安市 锦业路 003}} 23000 架构工程师}
```

### 声明&初始化

```go
package main

import (
	"fmt"
)

// 定义地址结构体（区域，街道，门牌号）
type Address struct {
	region string
	street string
	no     string
}

// 定义收货人结构体（名字，电话，地址）
type User struct {
	name string
	tel  string
	addr Address
}

// 定义员工，匿名嵌入user结构体
type Employee struct {
	User
	salary float64
	title  string
}

func main() {

	var ekk Employee
	fmt.Printf("%T: %#v\n", ekk, ekk)

	// 使用已有变量初始化

	ekk02 := Employee{
		User{
			"kk",
			"152000022222",
			Address{
				"陕西省西安市",
				"锦业路",
				"002",
			},
		},
		22000,
		"安全开发工程师",
	}

	ekk03 := Employee{
		User{
			"kk",
			"153000033333",
			Address{
				"陕西省西安市",
				"锦业路",
				"003",
			},
		},
		23000,
		"架构工程师",
	}


	fmt.Println(ekk02)
	fmt.Println(ekk03)
}


PS E:\go-phase-two\go-course> go run E:\go-phase-two\go-course\1test.go
main.Employee: main.Employee{User:main.User{name:"", tel:"", addr:main.Address{region:"", street:"", no:""}}, salary:0, title:""}
{{kk 152000022222 {陕西省西安市 锦业路 002}} 22000 安全开发工程师}
{{kk 153000033333 {陕西省西安市 锦业路 003}} 23000 架构工程师}
```


在初始化匿名嵌入的结构体对象时，需要遵循树状声明的结构，对于匿名嵌入的结构体可以使用结构体名来进行确定来指定初始化参数。

### 属性的访问和修改

```go
package main

import (
	"fmt"
)

// 定义地址结构体（区域，街道，门牌号）
type Address struct {
	region string
	street string
	no     string
}

// 定义收货人结构体（名字，电话，地址）
type User struct {
	name string
	tel  string
	addr Address
}

// 定义员工，匿名嵌入user结构体
type Employee struct {
	User
	salary float64
	title  string
}

func main() {

	var ekk Employee
	fmt.Printf("%T: %#v\n", ekk, ekk)

	// 使用已有变量初始化

	ekk02 := Employee{
		User{
			"kk",
			"152000022222",
			Address{
				"陕西省西安市",
				"锦业路",
				"002",
			},
		},
		22000,
		"安全开发工程师",
	}

	ekk03 := Employee{
		User{
			"kk",
			"152000022222",
			Address{
				"陕西省西安市",
				"锦业路",
				"003",
			},
		},
		23000,
		"架构工程师",
	}

	fmt.Println(ekk02)
	fmt.Println(ekk03)

	// 针对匿名嵌入结构体属性，可以通过对象，结构体名称，属性名范围和修改属性值

	fmt.Println(ekk03.User.name, ekk03.User.tel, ekk03.User.addr, ekk03.salary, ekk03.title)
	ekk03.User.tel = "15200003333"
	fmt.Println(ekk03)

	// 针对匿名嵌入结构体属性值访问和修改时间可省略结构体名称
	fmt.Println(ekk03.name, ekk03.tel, ekk03.addr, ekk03.salary, ekk03.title)

	ekk03.tel = "15200004444"
	fmt.Println(ekk03)
}

PS E:\go-phase-two\go-course> go run E:\go-phase-two\go-course\1test.go
main.Employee: main.Employee{User:main.User{name:"", tel:"", addr:main.Address{region:"", street:"", no:""}}, salary:0, title:""}
{{kk 152000022222 {陕西省西安市 锦业路 002}} 22000 安全开发工程师}
{{kk 152000022222 {陕西省西安市 锦业路 003}} 23000 架构工程师}
kk 152000022222 {陕西省西安市 锦业路 003} 23000 架构工程师
{{kk 15200003333 {陕西省西安市 锦业路 003}} 23000 架构工程师}
kk 15200003333 {陕西省西安市 锦业路 003} 23000 架构工程师
{{kk 15200004444 {陕西省西安市 锦业路 003}} 23000 架构工程师}
```


在访问和修改嵌入结构体的属性值时，可以通过对象名.结构体名称.属性名的方式进行访问和修改，结构体名称可以省略（匿名成员有一个隐式的名称），因此不能嵌套两个相同名称的结构体。当被嵌入结构体和嵌入结构体有相同的属性名时，在访问和修改嵌入结构体成员的属性值时不能省略结构体名称.


```go
package main

import (
	"fmt"
)

// 定义地址结构体（区域，街道，门牌号）
type Address struct {
	region string
	street string
	no     string
}

// 定义收货人结构体（名字，电话，地址）
type User struct {
	name string
	tel  string
	addr Address
}

// 定义员工，匿名嵌入user结构体
type Employee struct {
	User
	salary float64
	title  string
}

// 定义员工，匿名嵌入User结构体，并于User结构体具有相同的属性名

type Employee2 struct {
	User
	salary float64
	title  string
	addr   string
}

func main() {

	var ekk Employee
	fmt.Printf("%T: %#v\n", ekk, ekk)

	// 使用已有变量初始化

	ekk02 := Employee{
		User{
			"kk",
			"152000022222",
			Address{
				"陕西省西安市",
				"锦业路",
				"002",
			},
		},
		22000,
		"安全开发工程师",
	}

	ekk03 := Employee{
		User{
			"kk",
			"152000022222",
			Address{
				"陕西省西安市",
				"锦业路",
				"003",
			},
		},
		23000,
		"架构工程师",
	}

	fmt.Println(ekk02)
	fmt.Println(ekk03)

	// 针对匿名嵌入结构体属性，可以通过对象，结构体名称，属性名范围和修改属性值

	fmt.Println(ekk03.User.name, ekk03.User.tel, ekk03.User.addr, ekk03.salary, ekk03.title)
	ekk03.User.tel = "15200003333"
	fmt.Println(ekk03)

	// 针对匿名嵌入结构体属性值访问和修改时间可省略结构体名称
	fmt.Println(ekk03.name, ekk03.tel, ekk03.addr, ekk03.salary, ekk03.title)

	ekk03.tel = "15200004444"
	fmt.Println(ekk03)

	// 使用执行名称初始化
	ekk04 := Employee2{
		User: User{
			"kk",
			"15200002222",
			Address{
				"陕西省西安市",
				"锦业路",
				"003",
			},
		},
		salary: 25000,
		title:  "安全架构师",
		addr:   "陕西省西安市高新路001",
	}

	fmt.Printf("%#v\n", ekk04)

	// 分别访问别嵌入和想入结构体的addr属性
	fmt.Println(ekk04.addr)
	fmt.Println(ekk04.User.addr)

	// 分别修改被嵌入和嵌入结构体的addr属性
	ekk04.addr = "陕西省西安市高新路002"
	fmt.Println(ekk04.addr)
	fmt.Println(ekk04.User.addr)

	ekk04.User.addr.no = "004"
	fmt.Println(ekk04.addr)
	fmt.Println(ekk04.User.addr)
}


PS E:\go-phase-two\go-course> go run E:\go-phase-two\go-course\1test.go
main.Employee: main.Employee{User:main.User{name:"", tel:"", addr:main.Address{region:"", street:"", no:""}}, salary:0, title:""}
{{kk 152000022222 {陕西省西安市 锦业路 002}} 22000 安全开发工程师}
{{kk 152000022222 {陕西省西安市 锦业路 003}} 23000 架构工程师}
kk 152000022222 {陕西省西安市 锦业路 003} 23000 架构工程师
{{kk 15200003333 {陕西省西安市 锦业路 003}} 23000 架构工程师}
kk 15200003333 {陕西省西安市 锦业路 003} 23000 架构工程师
{{kk 15200004444 {陕西省西安市 锦业路 003}} 23000 架构工程师}
main.Employee2{User:main.User{name:"kk", tel:"15200002222", addr:main.Address{region:"陕西省西安市", street:"锦业 
路", no:"003"}}, salary:25000, title:"安全架构师", addr:"陕西省西安市高新路001"}
陕西省西安市高新路001
{陕西省西安市 锦业路 003}
陕西省西安市高新路002
{陕西省西安市 锦业路 003}
陕西省西安市高新路002
{陕西省西安市 锦业路 004}
```


## 指针类型嵌入

结构体嵌入（命名&匿名）类型也可以为结构体指针

### 定义 

### 声明&初始化&操作 

```go

package main

import (
	"fmt"
)

// 定义地址结构体（区域,街道,门牌号）
type Address struct {
	region string
	street string
	no     string
}

// 定义收货人结构体（名字,电话,地址）
type User struct {
	name string
	tel  string
	addr Address
}

// 定义员工,匿名嵌入user结构体
type Employee struct {
	User
	salary float64
	title  string
}

// 定义员工,匿名嵌入User结构体,并于User结构体具有相同的属性名

type Employee2 struct {
	User
	salary float64
	title  string
	addr   string
}

// 命名嵌入结构体指针
type PUser struct {
	name string
	tel  string
	addr *Address
}

// 匿名嵌入结构体指针
type PEmployee struct {
	*PUser
	salary float64
	title  string
}

func main() {

	var ekk Employee
	fmt.Printf("%T: %#v\n", ekk, ekk)

	// 使用已有变量初始化

	ekk02 := Employee{
		User{
			"kk",
			"152000022222",
			Address{
				"陕西省西安市",
				"锦业路",
				"002",
			},
		},
		22000,
		"安全开发工程师",
	}

	ekk03 := Employee{
		User{
			"kk",
			"152000022222",
			Address{
				"陕西省西安市",
				"锦业路",
				"003",
			},
		},
		23000,
		"架构工程师",
	}

	fmt.Println(ekk02)
	fmt.Println(ekk03)

	// 针对匿名嵌入结构体属性,可以通过对象,结构体名称,属性名范围和修改属性值

	fmt.Println(ekk03.User.name, ekk03.User.tel, ekk03.User.addr, ekk03.salary, ekk03.title)
	ekk03.User.tel = "15200003333"
	fmt.Println(ekk03)

	// 针对匿名嵌入结构体属性值访问和修改时间可省略结构体名称
	fmt.Println(ekk03.name, ekk03.tel, ekk03.addr, ekk03.salary, ekk03.title)

	ekk03.tel = "15200004444"
	fmt.Println(ekk03)

	// 使用执行名称初始化
	ekk04 := Employee2{
		User: User{
			"kk",
			"15200002222",
			Address{
				"陕西省西安市",
				"锦业路",
				"003",
			},
		},
		salary: 25000,
		title:  "安全架构师",
		addr:   "陕西省西安市高新路001",
	}

	fmt.Printf("%#v\n", ekk04)

	// 分别访问别嵌入和想入结构体的addr属性
	fmt.Println(ekk04.addr)
	fmt.Println(ekk04.User.addr)

	// 分别修改被嵌入和嵌入结构体的addr属性
	ekk04.addr = "陕西省西安市高新路002"
	fmt.Println(ekk04.addr)
	fmt.Println(ekk04.User.addr)

	ekk04.User.addr.no = "004"
	fmt.Println(ekk04.addr)
	fmt.Println(ekk04.User.addr)

	// 声明和初始化嵌入指针结构体对象
	paddr := Address{"陕西省西安市", "锦业路", "001"}
	puser := PUser{"kk", "15200000000", &paddr}

	pemployee := PEmployee{
		PUser:  &puser,
		salary: 25000,
		title:  "安全架构师",
	}

	fmt.Printf("%#v\n", pemployee)

	//属性访问和修改
	fmt.Println(paddr)      // no:001
	fmt.Println(puser.addr) // no:001

	puser.addr.no = "002"
	fmt.Println(paddr)      // no:002
	fmt.Println(puser.addr) // no:002

	fmt.Println(puser.name)           // name: kk
	fmt.Println(pemployee.PUser.name) // name: kk
	fmt.Println(pemployee.name)       // name: kk

	pemployee.name = "slience"

	fmt.Println(puser.name)           // name: slience
	fmt.Println(pemployee.PUser.name) // name: silence
	fmt.Println(pemployee.name)       // name: slience
}

PS E:\go-phase-two\go-course> go run E:\go-phase-two\go-course\1test.go
main.Employee: main.Employee{User:main.User{name:"", tel:"", addr:main.Address{region:"", street:"", no:""}}, salary:0, title:""}
{{kk 152000022222 {陕西省西安市 锦业路 002}} 22000 安全开发工程师}
{{kk 152000022222 {陕西省西安市 锦业路 003}} 23000 架构工程师}
kk 152000022222 {陕西省西安市 锦业路 003} 23000 架构工程师
{{kk 15200003333 {陕西省西安市 锦业路 003}} 23000 架构工程师}
kk 15200003333 {陕西省西安市 锦业路 003} 23000 架构工程师
{{kk 15200004444 {陕西省西安市 锦业路 003}} 23000 架构工程师}
main.Employee2{User:main.User{name:"kk", tel:"15200002222", addr:main.Address{region:"陕西省西安市", street:"锦业 
路", no:"003"}}, salary:25000, title:"安全架构师", addr:"陕西省西安市高新路001"}
陕西省西安市高新路001
{陕西省西安市 锦业路 003}
陕西省西安市高新路002
{陕西省西安市 锦业路 003}
陕西省西安市高新路002
{陕西省西安市 锦业路 004}
main.PEmployee{PUser:(*main.PUser)(0xc000078450), salary:25000, title:"安全架构师"}
{陕西省西安市 锦业路 001}
&{陕西省西安市 锦业路 001}
{陕西省西安市 锦业路 002}
&{陕西省西安市 锦业路 002}
kk
kk
kk
slience
slience
slience
```

使用属性为指针类型底层共享数据结构，当底层数据发生变化，所有引用都会发生影响


```go

package main

import (
	"fmt"
)

// 定义地址结构体（区域,街道,门牌号）
type Address struct {
	region string
	street string
	no     string
}

// 定义收货人结构体（名字,电话,地址）
type User struct {
	name string
	tel  string
	addr Address
}

// 定义员工,匿名嵌入user结构体
type Employee struct {
	User
	salary float64
	title  string
}

// 定义员工,匿名嵌入User结构体,并于User结构体具有相同的属性名

type Employee2 struct {
	User
	salary float64
	title  string
	addr   string
}

// 命名嵌入结构体指针
type PUser struct {
	name string
	tel  string
	addr *Address
}

// 匿名嵌入结构体指针
type PEmployee struct {
	*PUser
	salary float64
	title  string
}

func main() {

	var ekk Employee
	fmt.Printf("%T: %#v\n", ekk, ekk)

	// 使用已有变量初始化

	ekk02 := Employee{
		User{
			"kk",
			"152000022222",
			Address{
				"陕西省西安市",
				"锦业路",
				"002",
			},
		},
		22000,
		"安全开发工程师",
	}

	ekk03 := Employee{
		User{
			"kk",
			"152000022222",
			Address{
				"陕西省西安市",
				"锦业路",
				"003",
			},
		},
		23000,
		"架构工程师",
	}

	fmt.Println(ekk02)
	fmt.Println(ekk03)

	// 针对匿名嵌入结构体属性,可以通过对象,结构体名称,属性名范围和修改属性值

	fmt.Println(ekk03.User.name, ekk03.User.tel, ekk03.User.addr, ekk03.salary, ekk03.title)
	ekk03.User.tel = "15200003333"
	fmt.Println(ekk03)

	// 针对匿名嵌入结构体属性值访问和修改时间可省略结构体名称
	fmt.Println(ekk03.name, ekk03.tel, ekk03.addr, ekk03.salary, ekk03.title)

	ekk03.tel = "15200004444"
	fmt.Println(ekk03)

	// 使用执行名称初始化
	ekk04 := Employee2{
		User: User{
			"kk",
			"15200002222",
			Address{
				"陕西省西安市",
				"锦业路",
				"003",
			},
		},
		salary: 25000,
		title:  "安全架构师",
		addr:   "陕西省西安市高新路001",
	}

	fmt.Printf("%#v\n", ekk04)

	// 分别访问别嵌入和想入结构体的addr属性
	fmt.Println(ekk04.addr)
	fmt.Println(ekk04.User.addr)

	// 分别修改被嵌入和嵌入结构体的addr属性
	ekk04.addr = "陕西省西安市高新路002"
	fmt.Println(ekk04.addr)
	fmt.Println(ekk04.User.addr)

	ekk04.User.addr.no = "004"
	fmt.Println(ekk04.addr)
	fmt.Println(ekk04.User.addr)

	// 声明和初始化嵌入指针结构体对象
	paddr := Address{"陕西省西安市", "锦业路", "001"}
	puser := PUser{"kk", "15200000000", &paddr}

	pemployee := PEmployee{
		PUser:  &puser,
		salary: 25000,
		title:  "安全架构师",
	}

	fmt.Printf("%#v\n", pemployee)

	//属性访问和修改
	fmt.Println(paddr)      // no:001
	fmt.Println(puser.addr) // no:001

	puser.addr.no = "002"
	fmt.Println(paddr)      // no:002
	fmt.Println(puser.addr) // no:002

	fmt.Println(puser.name)           // name: kk
	fmt.Println(pemployee.PUser.name) // name: kk
	fmt.Println(pemployee.name)       // name: kk

	pemployee.name = "slience"

	fmt.Println(puser.name)           // name: slience
	fmt.Println(pemployee.PUser.name) // name: silence
	fmt.Println(pemployee.name)       // name: slience

	// 声明和初始化嵌入值结构体对象
	vaddr := Address{"陕西省西安市", "锦业路", "001"}
	vuser := User{"kk", "15200000000", vaddr}

	vemployee := Employee{
		User:   vuser,
		salary: 25000,
		title:  "安全架构师",
	}

	fmt.Printf("%#v\n", vemployee)
	// 属性访问和修改
	fmt.Println(vaddr)      // no:001
	fmt.Println(vuser.addr) // no:001

	vuser.addr.no = "002"
	fmt.Println(vaddr)      // no:001
	fmt.Println(vuser.addr) // no:002

	fmt.Println(vuser.name)          //name: kk
	fmt.Println(vemployee.User.name) // name: kk
	fmt.Println(vemployee.name)      // name.kk

	vemployee.name = "silence"

	fmt.Println(vuser.name)          // name: kk
	fmt.Println(vemployee.User.name) // name: slience
	fmt.Println(vemployee.name)      // name:slience

}


PS E:\go-phase-two\go-course> go run E:\go-phase-two\go-course\1test.go
main.Employee: main.Employee{User:main.User{name:"", tel:"", addr:main.Address{region:"", street:"", no:""}}, salary:0, title:""}
{{kk 152000022222 {陕西省西安市 锦业路 002}} 22000 安全开发工程师}
{{kk 152000022222 {陕西省西安市 锦业路 003}} 23000 架构工程师}
kk 152000022222 {陕西省西安市 锦业路 003} 23000 架构工程师
{{kk 15200003333 {陕西省西安市 锦业路 003}} 23000 架构工程师}
kk 15200003333 {陕西省西安市 锦业路 003} 23000 架构工程师
{{kk 15200004444 {陕西省西安市 锦业路 003}} 23000 架构工程师}
main.Employee2{User:main.User{name:"kk", tel:"15200002222", addr:main.Address{region:"陕西省西安市", street:"锦业 
路", no:"003"}}, salary:25000, title:"安全架构师", addr:"陕西省西安市高新路001"}
陕西省西安市高新路001
{陕西省西安市 锦业路 003}
陕西省西安市高新路002
{陕西省西安市 锦业路 003}
陕西省西安市高新路002
{陕西省西安市 锦业路 004}
main.PEmployee{PUser:(*main.PUser)(0xc00006e450), salary:25000, title:"安全架构师"}
{陕西省西安市 锦业路 001}
&{陕西省西安市 锦业路 001}
{陕西省西安市 锦业路 002}
&{陕西省西安市 锦业路 002}
kk
kk
kk
slience
slience
slience
main.Employee{User:main.User{name:"kk", tel:"15200000000", addr:main.Address{region:"陕西省西安市", street:"锦业路
", no:"001"}}, salary:25000, title:"安全架构师"}
{陕西省西安市 锦业路 001}
{陕西省西安市 锦业路 001}
{陕西省西安市 锦业路 001}
{陕西省西安市 锦业路 002}
kk
kk
kk
kk
silence
silence
```

使用属性为值类型，则在复制时发生拷贝，两者互不影响。

## 可用性 

结构体首字母大写则包外可见（公开的）， 否则进包内可以访问（内部的）
结构体属性名首字母大写包外可见（公开的）， 否则仅包内可以访问（内部的）

**组合：**
- 结构体名首字母大写，属性名大写： 结构体可在包外使用，且访问其大写的属性名。
- 结构体名首字母大写，属性名小写：结构体可在包外使用，且不能访问其小写的属性名。
- 结构体名首字母小写，属性名小写：结构体只能在包内使用，属性访问在结构体嵌入式由被嵌入结构体（外层）决定，被嵌入结构体名首字母大写是属性名包外可见，否则只能在包内使用。
- 结构体名首字母小写，属性名小写：结构体只能在包内使用

# 方法 

## 定义 

方法是添加了接受者的函数，接受者必须是自定义的类型。

```go
package main

import (
	"fmt"
)

// func (t type) method(parameters)  returns {

// }

// 定义结构体Dog
type Dog struct {
	name string
}

// 为结构体Dog 定义方法Call
func (dog Dog) Call() {
	fmt.Printf("%s: 汪汪\n", dog.name)
}

// 为结构体Dog定义方法SetName
func (dog Dog) SetName(name string) {
	dog.name = name
}

// 为结构体Dog定义方法PSetName（接受者为Dog的指针对象）

func (dog *Dog) PSetName(name string) {
	dog.name = name
}

func main() {

	// 初始化结构体对象
	dog := Dog{"豆豆"}

	// 调用结构体对象Call方法
	dog.Call()

	// 调用结构体对象SetName方法
	dog.SetName("大黄")

	dog.Call() // 这里结果是什么？？？ 为什么？？？ 怎么改？ 结果也是豆豆：汪汪

	// 调用结构体指针对象的PSetName方法
	(&dog).PSetName("大黄")
	dog.Call()

	dog2 := &Dog{"大毛"}
	// 调用结构体对象的Call方法
	(*&dog2).Call()

	dog2.PSetName("二毛")

}

PS E:\go-phase-two\go-course> go run E:\go-phase-two\go-course\1test.go
豆豆: 汪汪
豆豆: 汪汪
大黄: 汪汪
大毛: 汪汪
```


## 调用 
调用方法通过自定义类型的对象。方法名进行调用，在调用过程中对象传递（赋值）给方法的接受者（值类型，拷贝）

```go
package main

import (
	"fmt"
)

// func (t type) method(parameters)  returns {

// }

// 定义结构体Dog
type Dog struct {
	name string
}

// 为结构体Dog 定义方法Call
func (dog Dog) Call() {
	fmt.Printf("%s: 汪汪\n", dog.name)
}

// 为结构体Dog定义方法SetName
func (dog Dog) SetName(name string) {
	dog.name = name
}

// 为结构体Dog定义方法PSetName（接受者为Dog的指针对象）

func (dog *Dog) PSetName(name string) {
	dog.name = name
}

func main() {

	// 初始化结构体对象
	dog := Dog{"豆豆"}

	// 调用结构体对象Call方法
	dog.Call()

	// 调用结构体对象SetName方法
	dog.SetName("大黄")

	dog.Call() // 这里结果是什么？？？ 为什么？？？ 怎么改？ 结果也是豆豆：汪汪

	// 调用结构体指针对象的PSetName方法
	(&dog).PSetName("大黄")
	dog.Call()

	dog2 := &Dog{"大毛"}
	// 调用结构体对象的Call方法
	(*&dog2).Call()

	dog2.PSetName("二毛")

}

PS E:\go-phase-two\go-course> go run E:\go-phase-two\go-course\1test.go
豆豆: 汪汪
豆豆: 汪汪
大黄: 汪汪
大毛: 汪汪
```


## 指针接收者

### 声明 
```go

package main

import (
	"fmt"
)

// func (t type) method(parameters)  returns {

// }

// 定义结构体Dog
type Dog struct {
	name string
}

// 为结构体Dog 定义方法Call
func (dog Dog) Call() {
	fmt.Printf("%s: 汪汪\n", dog.name)
}

// 为结构体Dog定义方法SetName
func (dog Dog) SetName(name string) {
	dog.name = name
}

// 为结构体Dog定义方法PSetName（接受者为Dog的指针对象）

func (dog *Dog) PSetName(name string) {
	dog.name = name
}

func main() {

	// 初始化结构体对象
	dog := Dog{"豆豆"}

	// 调用结构体对象Call方法
	dog.Call()

	// 调用结构体对象SetName方法
	dog.SetName("大黄")

	dog.Call() // 这里结果是什么？？？ 为什么？？？ 怎么改？ 结果也是豆豆：汪汪

	// 调用结构体指针对象的PSetName方法
	(&dog).PSetName("大黄")
	dog.Call()

	dog2 := &Dog{"大毛"}
	// 调用结构体对象的Call方法
	(*&dog2).Call()

	dog2.PSetName("二毛")

}

```



当使用接受体指针对象调用值接收者的方法时，go 编译器会自动将指针对象 “解引用” 为值调用方法

当使用结构体对象调用指针接收者的方法时，go 编译器会自动将值对象"取引用" 为指针调用方法。

```go

package main

import (
	"fmt"
)

// func (t type) method(parameters)  returns {

// }

// 定义结构体Dog
type Dog struct {
	name string
}

// 为结构体Dog 定义方法Call
func (dog Dog) Call() {
	fmt.Printf("%s: 汪汪\n", dog.name)
}

// 为结构体Dog定义方法SetName
func (dog Dog) SetName(name string) {
	dog.name = name
}

// 为结构体Dog定义方法PSetName（接受者为Dog的指针对象）

func (dog *Dog) PSetName(name string) {
	dog.name = name
}

func main() {

	// 初始化结构体对象
	dog := Dog{"豆豆"}

	// 调用结构体对象Call方法
	dog.Call()

	// 调用结构体对象SetName方法
	dog.SetName("大黄")

	dog.Call() // 这里结果是什么？？？ 为什么？？？ 怎么改？ 结果也是豆豆：汪汪

	// 调用结构体指针对象的PSetName方法
	(&dog).PSetName("大黄")
	dog.Call()

	dog2 := &Dog{"大毛"}
	// 调用结构体对象的Call方法
	(*&dog2).Call()

	dog2.PSetName("二毛")

	// 调用结构体对象的Call方法
	(*dog2).Call()

	// 使用结构体对象调用指针接收者的PSetName 方法
	dog.PSetName("小黄")

	dog.Call()

	dog2.PSetName("二毛")

	// 使用结构体指针对象调用值接收者的Call方法

	dog2.Call()

	dog2.SetName("三毛")

	dog2.Call() //？？？

}


PS E:\go-phase-two\go-course> go run E:\go-phase-two\go-course\1test.go
豆豆: 汪汪
豆豆: 汪汪
大黄: 汪汪
大毛: 汪汪
二毛: 汪汪
小黄: 汪汪
二毛: 汪汪
二毛: 汪汪

```

注：取引用和解引用发生在接收者中，对于函数/方法的参数必须保持变量类型一一对应

该值接收者还是指针接收者，取决于是否现需要修改原始结构体
- 若不需要修改则使用值，若需要修改则使用指针
- 若存在指针接受者，则所有方法使用指针接收者。

对于接收者为指针类型的方法，需要注意在运行时若接收者为nil 则会发生报错。

```go
package main

import (
	"fmt"
)

// func (t type) method(parameters)  returns {

// }

// 定义结构体Dog
type Dog struct {
	name string
}

// 为结构体Dog 定义方法Call
func (dog Dog) Call() {
	fmt.Printf("%s: 汪汪\n", dog.name)
}

// 为结构体Dog定义方法SetName
func (dog Dog) SetName(name string) {
	dog.name = name
}

// 为结构体Dog定义方法PSetName（接受者为Dog的指针对象）

func (dog *Dog) PSetName(name string) {
	dog.name = name
}

func main() {

	// 初始化结构体对象
	dog := Dog{"豆豆"}

	// 调用结构体对象Call方法
	dog.Call()

	// 调用结构体对象SetName方法
	dog.SetName("大黄")

	dog.Call() // 这里结果是什么？？？ 为什么？？？ 怎么改？ 结果也是豆豆：汪汪

	// 调用结构体指针对象的PSetName方法
	(&dog).PSetName("大黄")
	dog.Call()

	dog2 := &Dog{"大毛"}
	// 调用结构体对象的Call方法
	(*&dog2).Call()

	dog2.PSetName("二毛")

	// 调用结构体对象的Call方法
	(*dog2).Call()

	// 使用结构体对象调用指针接收者的PSetName 方法
	dog.PSetName("小黄")

	dog.Call()

	dog2.PSetName("二毛")

	// 使用结构体指针对象调用值接收者的Call方法

	dog2.Call()

	dog2.SetName("三毛")

	dog2.Call() //？？？

	var dog3 *Dog

	dog3.Call()

}

PS E:\go-phase-two\go-course> go run E:\go-phase-two\go-course\1test.go
豆豆: 汪汪
豆豆: 汪汪
大黄: 汪汪
大毛: 汪汪
二毛: 汪汪
小黄: 汪汪
二毛: 汪汪
二毛: 汪汪
panic: runtime error: invalid memory address or nil pointer dereference
[signal 0xc0000005 code=0x0 addr=0x8 pc=0x49c642]

goroutine 1 [running]:
main.main()
        E:/go-phase-two/go-course/1test.go:75 +0x562
exit status 2
```


## 匿名嵌入 

若结构体匿名嵌入带有方法的结构体时，则在外部结构体可以调用嵌入结构体的方法，并且在调用时只有嵌入的字段会传递给嵌入结构体方法的接收者。

当被嵌入结构体与嵌入结构体具有相同名称的方法时，则使用对象.方法名调用被嵌入结构体方法。若想要调用嵌入结构体方法，则使用对象.嵌入结构体名.方法.

```go
package main

import "fmt"

// 定义User结构体方法
type User struct {
	name string
	addr string
}

func NewUser(name, tel string) *User {
	return &User{name, tel}
}

func (user *User) SetName(name string) {
	user.name = name
}

func (user *User) GetName() string {
	return user.name
}

func (user *User) SetAddr(addr string) {
	user.addr = addr
}

func (user *User) GetAddr() string {
	return user.addr
}

// 定义Employee 结构体和方法
type Employee struct {
	*User
	title  string
	salary float64
}

func NewEmployee(name, addr, title string, salary float64) *Employee {
	return &Employee{
		NewUser(name, addr),
		title,
		salary,
	}

}

func (employee *Employee) GetName() string {
	return fmt.Sprintf("[%s]%s", employee.title, employee.name)

}

func (employee *Employee) SetTitle(title string) {
	employee.title = title
}

func (employee *Employee) GetTitle() string {
	return employee.title
}

func (employee *Employee) SetSalary(salary float64) {
	employee.salary = salary
}

func (employee *Employee) GetSalary() float64 {
	return employee.salary
}

func main() {
	me := NewEmployee("kk", "北京朝阳区", "开发工程师", 15000)
	fmt.Printf("%#v\n", me)

	fmt.Println(me.GetName())      // 调用Employee.GetName
	fmt.Println(me.User.GetName()) //调用User.GetName

	me.SetAddr("西安市高新区")      // 调用User.SetAddr
	fmt.Println(me.GetAddr()) // 调用User.GetAddr

	me.SetSalary(13000)         // 调用Employee.Setsalary
	fmt.Println(me.GetSalary()) // 调用Employee.GetSalary

	fmt.Printf("%#v\n", me)

}

PS E:\go-phase-two\go-course> go run E:\go-phase-two\go-course\1test.go
&main.Employee{User:(*main.User)(0xc0000044a0), title:"开发工程师", salary:15000}
[开发工程师]kk
kk
西安市高新区
13000
&main.Employee{User:(*main.User)(0xc0000044a0), title:"开发工程师", salary:13000}

```

## 方法值/方法表达式

方法也可以赋值给变量，存储在数组、切片、映射中，也可作为参数传递给函数或作为函数返回值进行返回方法有两种，一种时使用对象/对象指针调用的(方法值)，另一种时有类型/类型指针调用的(方法表达式)

### 1) 方法值

```go
package main

import "fmt"

// 定义结构体Dog
type Dog struct {
	name string
}

// 为结构体Dog 定义方法Call
func (dog Dog) Call() {
	fmt.Printf("%s: 汪汪\n", dog.name)
}

// 为结构体Dog定义方法SetName
func (dog Dog) SetName(name string) {
	dog.name = name
}

// 为结构体Dog定义方法PSetName（接受者为Dog的指针对象）

func (dog *Dog) PSetName(name string) {
	dog.name = name
}

// 定义User结构体方法
type User struct {
	name string
	addr string
}

func NewUser(name, tel string) *User {
	return &User{name, tel}
}

func (user *User) SetName(name string) {
	user.name = name
}

func (user *User) GetName() string {
	return user.name
}

func (user *User) SetAddr(addr string) {
	user.addr = addr
}

func (user *User) GetAddr() string {
	return user.addr
}

// 定义Employee 结构体和方法
type Employee struct {
	*User
	title  string
	salary float64
}

func NewEmployee(name, addr, title string, salary float64) *Employee {
	return &Employee{
		NewUser(name, addr),
		title,
		salary,
	}

}

func (employee *Employee) GetName() string {
	return fmt.Sprintf("[%s]%s", employee.title, employee.name)

}

func (employee *Employee) SetTitle(title string) {
	employee.title = title
}

func (employee *Employee) GetTitle() string {
	return employee.title
}

func (employee *Employee) SetSalary(salary float64) {
	employee.salary = salary
}

func (employee *Employee) GetSalary() float64 {
	return employee.salary
}

func main() {

	dog := Dog{"豆豆"}

	dog2 := &Dog{"大毛"}

	// 方法值
	methodv01 := dog.PSetName
	methodv02 := dog.Call
	methodv03 := dog2.PSetName
	methodv04 := dog2.Call

	fmt.Printf("%T, %T, %T, %T\n", methodv01, methodv02, methodv03, methodv04)

	// 通过方法值调用方法
	fmt.Printf("%p, %p\n", &dog, dog2)

	methodv01("乐乐1")
	methodv02() // ???

	methodv03("乐乐2")
	methodv04() // ???

	fmt.Println(dog.name, dog2.name)
}

PS E:\go-phase-two\go-course> go run E:\go-phase-two\go-course\1test.go
func(string), func(), func(string), func()
0xc0000461f0, 0xc000046200
豆豆: 汪汪
大毛: 汪汪
乐乐1 乐乐2
```
在方法表达式赋值时若方法接收者为值类型，则在赋值时会将值类型拷贝（若调用为指针则自动解引用拷贝）

```go
package main

import "fmt"

// 定义结构体Dog
type Dog struct {
	name string
}

// 为结构体Dog 定义方法Call
func (dog Dog) Call() {
	fmt.Printf("%s: 汪汪\n", dog.name)
}

// 为结构体Dog定义方法SetName
func (dog Dog) SetName(name string) {
	dog.name = name
}

// 为结构体Dog定义方法PSetName（接收者为Dog的指针对象）

func (dog *Dog) PSetName(name string) {
	dog.name = name
}

func (dog *Dog) PCall() {
	fmt.Printf("%s： 汪汪\n", dog.name)

}

// 定义User结构体方法
type User struct {
	name string
	addr string
}

func NewUser(name, tel string) *User {
	return &User{name, tel}
}

func (user *User) SetName(name string) {
	user.name = name
}

func (user *User) GetName() string {
	return user.name
}

func (user *User) SetAddr(addr string) {
	user.addr = addr
}

func (user *User) GetAddr() string {
	return user.addr
}

// 定义Employee 结构体和方法
type Employee struct {
	*User
	title  string
	salary float64
}

func NewEmployee(name, addr, title string, salary float64) *Employee {
	return &Employee{
		NewUser(name, addr),
		title,
		salary,
	}

}

func (employee *Employee) GetName() string {
	return fmt.Sprintf("[%s]%s", employee.title, employee.name)

}

func (employee *Employee) SetTitle(title string) {
	employee.title = title
}

func (employee *Employee) GetTitle() string {
	return employee.title
}

func (employee *Employee) SetSalary(salary float64) {
	employee.salary = salary
}

func (employee *Employee) GetSalary() float64 {
	return employee.salary
}

func main() {

	dog := Dog{"豆豆"}

	dog2 := &Dog{"大毛"}

	// 方法值
	methodv01 := dog.PSetName
	methodv02 := dog.Call
	methodv03 := dog2.PSetName
	methodv04 := dog2.Call

	fmt.Printf("%T, %T, %T, %T\n", methodv01, methodv02, methodv03, methodv04)

	// 通过方法值调用方法
	fmt.Printf("%p, %p\n", &dog, dog2)

	methodv01("乐乐1")
	methodv02() // ???

	methodv03("乐乐2")
	methodv04() // ???

	fmt.Println(dog.name, dog2.name)

	// 为结构体Dog定义方法PCall（接收者为Dog的指针对象）

	methodv05 := dog.PCall(小明)
	methodv06 := dog2.PCall(小红)

	methodv05()
	methodv06()
}

测试失败
```

### 方法表达式 

```go
package main

import "fmt"

// 定义结构体Dog
type Dog struct {
	name string
}

// 为结构体Dog 定义方法Call
func (dog Dog) Call() {
	fmt.Printf("%s: 汪汪\n", dog.name)
}

// 为结构体Dog定义方法SetName
func (dog Dog) SetName(name string) {
	dog.name = name
}

// 为结构体Dog定义方法PSetName（接收者为Dog的指针对象）

func (dog *Dog) PSetName(name string) {
	dog.name = name
}

func (dog *Dog) PCall() {
	fmt.Printf("%s： 汪汪\n", dog.name)

}

// 定义User结构体方法
type User struct {
	name string
	addr string
}

func NewUser(name, tel string) *User {
	return &User{name, tel}
}

func (user *User) SetName(name string) {
	user.name = name
}

func (user *User) GetName() string {
	return user.name
}

func (user *User) SetAddr(addr string) {
	user.addr = addr
}

func (user *User) GetAddr() string {
	return user.addr
}

// 定义Employee 结构体和方法
type Employee struct {
	*User
	title  string
	salary float64
}

func NewEmployee(name, addr, title string, salary float64) *Employee {
	return &Employee{
		NewUser(name, addr),
		title,
		salary,
	}

}

func (employee *Employee) GetName() string {
	return fmt.Sprintf("[%s]%s", employee.title, employee.name)

}

func (employee *Employee) SetTitle(title string) {
	employee.title = title
}

func (employee *Employee) GetTitle() string {
	return employee.title
}

func (employee *Employee) SetSalary(salary float64) {
	employee.salary = salary
}

func (employee *Employee) GetSalary() float64 {
	return employee.salary
}

func main() {

	dog := Dog{"豆豆"}

	dog2 := &Dog{"大毛"}

	// 方法值
	methodv01 := dog.PSetName
	methodv02 := dog.Call
	methodv03 := dog2.PSetName
	methodv04 := dog2.Call

	fmt.Printf("%T, %T, %T, %T\n", methodv01, methodv02, methodv03, methodv04)

	// 通过方法值调用方法
	fmt.Printf("%p, %p\n", &dog, dog2)

	methodv01("乐乐1")
	methodv02() // ???

	methodv03("乐乐2")
	methodv04() // ???

	fmt.Println(dog.name, dog2.name)

	// // 为结构体Dog定义方法PCall（接收者为Dog的指针对象）

	// methodv05 := dog.PCall(小明)
	// methodv06 := dog2.PCall(小红)

	// methodv05()
	// methodv06()

	// 方式表达式
	methode01 := (*Dog).PSetName
	methode02 := Dog.Call

	fmt.Printf("%T, %T\n", methode01, methode02)

	// 通过方法表达式指定对象调用方法
	methode01(&dog, "贝贝1") // (*Dog).PSetName(&dog, "贝贝1")
	methode01(dog2, "贝贝2") // (*Dog).PSetName(dog2, "贝贝2")

	methode02(dog)   //Dog.Call(dog)
	methode02(*dog2) //Dog.Call(*dog2)

	// 取引用
	// methode03 := Dog.PSetName
	methode04 := (*Dog).Call

	fmt.Printf("%T\n", methode04)

	methode04(&dog)
	methode04(dog2)
}


PS E:\go-phase-two\go-course> go run E:\go-phase-two\go-course\1test.go
func(string), func(), func(string), func()
0xc0000461f0, 0xc000046200
豆豆: 汪汪
大毛: 汪汪
乐乐1 乐乐2
func(*main.Dog, string), func(main.Dog)
贝贝1: 汪汪
贝贝2: 汪汪
func(*main.Dog)
贝贝1: 汪汪
贝贝2: 汪汪
```


方法表达式在赋值时，针对接收者为值类型的方法使用类型名或类型指针访问（go 自动为指针变量生成隐式的指针类型接收者方法），针对接收者为指针类型则使用类型指针访问。

同时在调用时需要传递对应的值对象或指针对象

### 自动生成指针接收者方法 

为什么根据接收者为值类型生成对应指针类型接收者方法，而不根据接收者为指针类型生成对应值接收者方法

```go
package main

import (
	"fmt"	"go/types"
)



type User struct {
	id int 
	name string
}

func (user User) SetId(id int) {
	user.id
}

func (user User) GetId() int {
	return user.id
}

/** 
// 隐式
func (user *User) SetId(id int) {
	// 获取user地址的值，并拷贝调用SetId, 只影响拷贝(*user)的值，并不影响调用者的值
	// 与(user User)SetId(id int) 行为一致
	(*user).SetId(id)
}

func (user *User) GetId() int {
	return (*user).GetId()
}

// 使用值和指针都不改变调用者，行为一致
*/

func (user *User) SetName(name string)  {
	user.name = name 
}

func (user *User) GetName() string {
	return user.name	
}

/**

func (user User) SetName(name string) {
	// user为值接收者，拷贝，在调用（&user），SetName只会影响接收者（user）的值，并不会影响调用者的值
	// 与(user *User)SetName(name string) 行为不一样
	(&user).SetName(name)
}

func (user User) GetName(name string) {
	return (&user).GetName()
}

// 使用指针改变接收者，使用值不改变调用者，行为不一致

*/




func main() {

	type User struct {
		Fields(2):
			id  int,
			name string,
		Methods(2):
			func GetId(objs.User) (int) {},
			func SetId(objs.User, int) {},
	}

	*{
		type User struct {
			Fields(2):
				id int,
				name string,
			Methods(2):
				func GetId(objs.User) (int) {},
				func SetId(objs.User, int) {},
		}
		Methods(4):
			func GetId(*objs.User) (int) {},
			func GetName(*objs.User) (string) {},
			func SetName(*objs.User, int) {},
			func SetName(*objs.user, string) {} 
	}
	
}

测试失败
```
使用反射获取User对象和*User对象结构


# 接口 

接口是自定义类型，是对其他类型行为的抽象

## 定义 

接口定义使用interface 表示，声明了一系列的函数签名（函数名、函数参数、函数返回值），在定义接口可以指定接口名称，在后续声明接口变量时使用。

## 声明 
声明接口变量只需要定义变量类型为接口名，此时变量被初始化为nil 

```go
package main

import "fmt"

// 定义Sender接口
type Sender interface {
	Send(to, msg string) error
	SendAll(tos []string, msg string) error
}

func main() {

	var sender Sender
	fmt.Printf("%T: %v", sender, sender)
}
PS E:\go-phase-two\go-course> go run E:\go-phase-two\go-course\1test.go
<nil>: <nil>
```

## 赋值 
### 类型对象 

当自定义类型实现了接口类型中声明的所有函数时，则该类型的对象可以赋值给接口变量，并使用接口变量调用实现的接口。

1. 方法接收者为值类型的方法 

```go
package main

import "fmt"

// 定义Sender接口
type Sender interface {
	Send(to, msg string) error
	SendAll(tos []string, msg string) error
}

// 定义EmailSender结构体
type EmailSender struct {
	addr           string
	port           int
	user, password string
}

func NewEmailSender(addr string, port int, user, password string) EmailSender {
	return EmailSender{addr, port, user, password}
}

// 接收者为值对象
func (sender EmailSender) Send(to, msg string) error {
	fmt.Printf("发送邮件给：%s， 内容： %s\n", to, msg)
	return nil
}

// 接收者为值对象
func (sender EmailSender) SendAll(tos []string, msg string) error {
	for _, to := range tos {
		fmt.Printf("发送邮件给%s, 内容： %s\n", to, msg)
	}
	return nil
}

func main() {

	var sender Sender
	fmt.Printf("%T: %v", sender, sender)

	// 赋值sender对象为NewEmailSender值对象
	var sender01 Sender = EmailSender{"smtp.qq.com", 465, "kk", "123456"}
	fmt.Printf("%T：%v\n", sender01, sender01)
	sender01.Send("imsilence@outlook.com", "Hi Silence")
	sender01.SendAll([]string{"imsilence@outlook.com", "iamkk@outlook.com"}, "早上好")
}


PS E:\go-phase-two\go-course> go run E:\go-phase-two\go-course\1test.go
<nil>: <nil>main.EmailSender：{smtp.qq.com 465 kk 123456}
发送邮件给：imsilence@outlook.com， 内容： Hi Silence
发送邮件给imsilence@outlook.com, 内容： 早上好
发送邮件给iamkk@outlook.com, 内容： 早上好
```


2. 方法接收者全为指针类型的

```go
package main

import "fmt"

// 定义Sender接口
type Sender interface {
	Send(to, msg string) error
	SendAll(tos []string, msg string) error
}

// 定义EmailSender结构体
type EmailSender struct {
	addr           string
	port           int
	user, password string
}

func NewEmailSender(addr string, port int, user, password string) EmailSender {
	return EmailSender{addr, port, user, password}
}

// 接收者为值对象
func (sender EmailSender) Send(to, msg string) error {
	fmt.Printf("发送邮件给：%s， 内容： %s\n", to, msg)
	return nil
}

// 接收者为值对象
func (sender EmailSender) SendAll(tos []string, msg string) error {
	for _, to := range tos {
		fmt.Printf("发送邮件给%s, 内容： %s\n", to, msg)
	}
	return nil
}

type SMSSender struct {
	url, key string
}

func NewSMSSender(url, key string) *SMSSender {
	return &SMSSender{url, key}

}

// 接收者为指针对象
func (sender *SMSSender) Send(to, msg string) error {

	fmt.Printf("发送短信给： %s, 内容：%s\n", to, msg)
	return nil
}

// 接收者为指针对象
func (sender *SMSSender) SendAll(tos []string, msg string) error {
	for _, to := range tos {
		fmt.Printf("发送短信给： %s, 内容：%s\n", to, msg)
	}
	return nil
}

func main() {

	var sender Sender
	fmt.Printf("%T: %v", sender, sender)

	// 赋值sender对象为NewEmailSender值对象
	var sender01 Sender = EmailSender{"smtp.qq.com", 465, "kk", "123456"}
	fmt.Printf("%T：%v\n", sender01, sender01)
	sender01.Send("imsilence@outlook.com", "Hi Silence")
	sender01.SendAll([]string{"imsilence@outlook.com", "iamkk@outlook.com"}, "早上好")

	// 赋值sender 对象为NewSMSSender指针对象
	var sender02 Sender = &SMSSender{"http://v.juhe.cn/sms/send", "123"}
	fmt.Printf("%T: %v\n", sender02, sender02)
	sender02.Send("15200000000", "Hi KK")
	sender02.SendAll([]string{"15200000000", "15800000000"}, "早上好")
}

PS E:\go-phase-two\go-course> go run E:\go-phase-two\go-course\1test.go
<nil>: <nil>main.EmailSender：{smtp.qq.com 465 kk 123456}
发送邮件给：imsilence@outlook.com， 内容： Hi Silence
发送邮件给imsilence@outlook.com, 内容： 早上好
发送邮件给iamkk@outlook.com, 内容： 早上好
*main.SMSSender: &{http://v.juhe.cn/sms/send 123}
发送短信给： 15200000000, 内容：Hi KK
发送短信给： 15200000000, 内容：早上好
发送短信给： 15800000000, 内容：早上好
```


3. 方法接收者既有值类型又有指针类型的

```go
package main

import "fmt"

// 定义Sender接口
type Sender interface {
	Send(to, msg string) error
	SendAll(tos []string, msg string) error
}

// 定义EmailSender结构体
type EmailSender struct {
	addr           string
	port           int
	user, password string
}

func NewEmailSender(addr string, port int, user, password string) EmailSender {
	return EmailSender{addr, port, user, password}
}

// 接收者为值对象
func (sender EmailSender) Send(to, msg string) error {
	fmt.Printf("发送邮件给：%s， 内容： %s\n", to, msg)
	return nil
}

// 接收者为值对象
func (sender EmailSender) SendAll(tos []string, msg string) error {
	for _, to := range tos {
		fmt.Printf("发送邮件给%s, 内容： %s\n", to, msg)
	}
	return nil
}

type SMSSender struct {
	url, key string
}

func NewSMSSender(url, key string) *SMSSender {
	return &SMSSender{url, key}

}

// 接收者为指针对象
func (sender *SMSSender) Send(to, msg string) error {

	fmt.Printf("发送短信给： %s, 内容：%s\n", to, msg)
	return nil
}

// 接收者为指针对象
func (sender *SMSSender) SendAll(tos []string, msg string) error {
	for _, to := range tos {
		fmt.Printf("发送短信给： %s, 内容：%s\n", to, msg)
	}
	return nil
}

// 定义WeChatSender 结构体
type WeChatSender struct {
	sessionFile string
}

func NewWeChatSender(sessionFile string) *WeChatSender {
	return &WeChatSender{sessionFile}
}

// 接收者为值对象
func (sender WeChatSender) Send(to, msg string) error {
	fmt.Printf("发送短信给：%s, 内容：%s\n", to, msg)
	return nil
}

// 接收者为指针对象
func (sender *WeChatSender) SendAll(tos []string, msg string) error {
	for _, to := range tos {
		fmt.Printf("发送微信给：%s, 内容：%s\n", to, msg)
	}
	return nil
}

func main() {

	var sender Sender
	fmt.Printf("%T: %v", sender, sender)

	// 赋值sender对象为NewEmailSender值对象
	var sender01 Sender = EmailSender{"smtp.qq.com", 465, "kk", "123456"}
	fmt.Printf("%T：%v\n", sender01, sender01)
	sender01.Send("imsilence@outlook.com", "Hi Silence")
	sender01.SendAll([]string{"imsilence@outlook.com", "iamkk@outlook.com"}, "早上好")

	// 赋值sender 对象为NewSMSSender指针对象
	var sender02 Sender = &SMSSender{"http://v.juhe.cn/sms/send", "123"}
	fmt.Printf("%T: %v\n", sender02, sender02)
	sender02.Send("15200000000", "Hi KK")
	sender02.SendAll([]string{"15200000000", "15800000000"}, "早上好")

	// 赋值应该使用哪一个？原因？
	// var sender03 Sender = WeChatSender{"/tmp/wx.session"} // ？？？
	// var sender03 Sender = &WeChatSender{"/tmp/wx.session"} // ？？？
	var sender03 Sender = &WeChatSender{"/tmp.session"}
	fmt.Printf("%T: %v\n", sender03, sender03)
	sender03.Send("7867258956", "Hi kk")
	sender03.SendAll([]string{"785736895", "785632342"}, "早上好")

	// 问题？
	// var sender04 Sender = &EmailSender{"smtp.qq.com", 465, "kk", "123456"} // ???
	// var sender05 Sender = SMSSender{"http://v.juhe.cn/sms/send", "1234"} // ???
}


PS E:\go-phase-two\go-course> go run E:\go-phase-two\go-course\1test.go
<nil>: <nil>main.EmailSender：{smtp.qq.com 465 kk 123456}
发送邮件给：imsilence@outlook.com， 内容： Hi Silence
发送邮件给imsilence@outlook.com, 内容： 早上好
发送邮件给iamkk@outlook.com, 内容： 早上好
*main.SMSSender: &{http://v.juhe.cn/sms/send 123}
发送短信给： 15200000000, 内容：Hi KK
发送短信给： 15200000000, 内容：早上好
发送短信给： 15800000000, 内容：早上好
*main.WeChatSender: &{/tmp.session}
发送短信给：7867258956, 内容：Hi kk
发送微信给：785736895, 内容：早上好
发送微信给：785632342, 内容：早上好

```


### 接口对象 

当接口(A)包含另外一个接口(B)中声明的所有函数时(A 接口函数时B 接口函数的父集， B时A的子集)，则接口(A)的对象也可以赋值给子集的接口(B)变量 

```go
package main

import "fmt"

// 定义Sender接口
type Sender interface {
	Send(to, msg string) error
	SendAll(tos []string, msg string) error
}

// 定义EmailSender结构体
type EmailSender struct {
	addr           string
	port           int
	user, password string
}

func NewEmailSender(addr string, port int, user, password string) EmailSender {
	return EmailSender{addr, port, user, password}
}

// 接收者为值对象
func (sender EmailSender) Send(to, msg string) error {
	fmt.Printf("发送邮件给：%s， 内容： %s\n", to, msg)
	return nil
}

// 接收者为值对象
func (sender EmailSender) SendAll(tos []string, msg string) error {
	for _, to := range tos {
		fmt.Printf("发送邮件给%s, 内容： %s\n", to, msg)
	}
	return nil
}

type SMSSender struct {
	url, key string
}

func NewSMSSender(url, key string) *SMSSender {
	return &SMSSender{url, key}

}

// 接收者为指针对象
func (sender *SMSSender) Send(to, msg string) error {

	fmt.Printf("发送短信给： %s, 内容：%s\n", to, msg)
	return nil
}

// 接收者为指针对象
func (sender *SMSSender) SendAll(tos []string, msg string) error {
	for _, to := range tos {
		fmt.Printf("发送短信给： %s, 内容：%s\n", to, msg)
	}
	return nil
}

// 定义WeChatSender 结构体
type WeChatSender struct {
	sessionFile string
}

func NewWeChatSender(sessionFile string) *WeChatSender {
	return &WeChatSender{sessionFile}
}

// 接收者为值对象
func (sender WeChatSender) Send(to, msg string) error {
	fmt.Printf("发送短信给：%s, 内容：%s\n", to, msg)
	return nil
}

// 接收者为指针对象
func (sender *WeChatSender) SendAll(tos []string, msg string) error {
	for _, to := range tos {
		fmt.Printf("发送微信给：%s, 内容：%s\n", to, msg)
	}
	return nil
}

type SingleSender interface {
	Send(to, mgs string) error
}

func main() {

	var sender Sender
	fmt.Printf("%T: %v", sender, sender)

	// 赋值sender对象为NewEmailSender值对象
	var sender01 Sender = EmailSender{"smtp.qq.com", 465, "kk", "123456"}
	fmt.Printf("%T：%v\n", sender01, sender01)
	sender01.Send("imsilence@outlook.com", "Hi Silence")
	sender01.SendAll([]string{"imsilence@outlook.com", "iamkk@outlook.com"}, "早上好")

	// 赋值sender 对象为NewSMSSender指针对象
	var sender02 Sender = &SMSSender{"http://v.juhe.cn/sms/send", "123"}
	fmt.Printf("%T: %v\n", sender02, sender02)
	sender02.Send("15200000000", "Hi KK")
	sender02.SendAll([]string{"15200000000", "15800000000"}, "早上好")

	// 赋值应该使用哪一个？原因？
	// var sender03 Sender = WeChatSender{"/tmp/wx.session"} // ？？？
	// var sender03 Sender = &WeChatSender{"/tmp/wx.session"} // ？？？
	var sender03 Sender = &WeChatSender{"/tmp.session"}
	fmt.Printf("%T: %v\n", sender03, sender03)
	sender03.Send("7867258956", "Hi kk")
	sender03.SendAll([]string{"785736895", "785632342"}, "早上好")

	// 问题？
	// var sender04 Sender = &EmailSender{"smtp.qq.com", 465, "kk", "123456"} // ???
	// var sender05 Sender = SMSSender{"http://v.juhe.cn/sms/send", "1234"} // ???

	var ssender SingleSender = sender03
	ssender.Send("78523423423", "Hi KK")

}

PS E:\go-phase-two\go-course> go run E:\go-phase-two\go-course\1test.go
<nil>: <nil>main.EmailSender：{smtp.qq.com 465 kk 123456}
发送邮件给：imsilence@outlook.com， 内容： Hi Silence
发送邮件给imsilence@outlook.com, 内容： 早上好
发送邮件给iamkk@outlook.com, 内容： 早上好
*main.SMSSender: &{http://v.juhe.cn/sms/send 123}
发送短信给： 15200000000, 内容：Hi KK
发送短信给： 15200000000, 内容：早上好
发送短信给： 15800000000, 内容：早上好
*main.WeChatSender: &{/tmp.session}
发送短信给：7867258956, 内容：Hi kk
发送微信给：785736895, 内容：早上好
发送微信给：785632342, 内容：早上好
发送短信给：78523423423, 内容：Hi KK
```

若两个接口声明同样的函数签名，则两个接口完全等价

当类型与父集接口赋值给接口变量口，只能调用变量定义接口中声明的函数(方法)


## 类型断言&查询

当父集接口或者类型对象赋值给接口变量后，需要将接口变量重新转为原来的类型，需要使用类型断言/查询

### 断言

语法：接口变量(Type)

```go
package main

import "fmt"

// 定义Sender接口
type Sender interface {
	Send(to, msg string) error
	SendAll(tos []string, msg string) error
}

// 定义EmailSender结构体
type EmailSender struct {
	addr           string
	port           int
	user, password string
}

func NewEmailSender(addr string, port int, user, password string) EmailSender {
	return EmailSender{addr, port, user, password}
}

// 接收者为值对象
func (sender EmailSender) Send(to, msg string) error {
	fmt.Printf("发送邮件给：%s， 内容： %s\n", to, msg)
	return nil
}

// 接收者为值对象
func (sender EmailSender) SendAll(tos []string, msg string) error {
	for _, to := range tos {
		fmt.Printf("发送邮件给%s, 内容： %s\n", to, msg)
	}
	return nil
}

type SMSSender struct {
	url, key string
}

func NewSMSSender(url, key string) *SMSSender {
	return &SMSSender{url, key}

}

// 接收者为指针对象
func (sender *SMSSender) Send(to, msg string) error {

	fmt.Printf("发送短信给： %s, 内容：%s\n", to, msg)
	return nil
}

// 接收者为指针对象
func (sender *SMSSender) SendAll(tos []string, msg string) error {
	for _, to := range tos {
		fmt.Printf("发送短信给： %s, 内容：%s\n", to, msg)
	}
	return nil
}

// 定义WeChatSender 结构体
type WeChatSender struct {
	sessionFile string
}

func NewWeChatSender(sessionFile string) *WeChatSender {
	return &WeChatSender{sessionFile}
}

// 接收者为值对象
func (sender WeChatSender) Send(to, msg string) error {
	fmt.Printf("发送短信给：%s, 内容：%s\n", to, msg)
	return nil
}

// 接收者为指针对象
func (sender *WeChatSender) SendAll(tos []string, msg string) error {
	for _, to := range tos {
		fmt.Printf("发送微信给：%s, 内容：%s\n", to, msg)
	}
	return nil
}

type SingleSender interface {
	Send(to, mgs string) error
}

func main() {

	var sender Sender
	fmt.Printf("%T: %v", sender, sender)

	// 赋值sender对象为NewEmailSender值对象
	var sender01 Sender = EmailSender{"smtp.qq.com", 465, "kk", "123456"}
	fmt.Printf("%T：%v\n", sender01, sender01)
	sender01.Send("imsilence@outlook.com", "Hi Silence")
	sender01.SendAll([]string{"imsilence@outlook.com", "iamkk@outlook.com"}, "早上好")

	// 赋值sender 对象为NewSMSSender指针对象
	var sender02 Sender = &SMSSender{"http://v.juhe.cn/sms/send", "123"}
	fmt.Printf("%T: %v\n", sender02, sender02)
	sender02.Send("15200000000", "Hi KK")
	sender02.SendAll([]string{"15200000000", "15800000000"}, "早上好")

	// 赋值应该使用哪一个？原因？
	// var sender03 Sender = WeChatSender{"/tmp/wx.session"} // ？？？
	// var sender03 Sender = &WeChatSender{"/tmp/wx.session"} // ？？？
	var sender03 Sender = &WeChatSender{"/tmp.session"}
	fmt.Printf("%T: %v\n", sender03, sender03)
	sender03.Send("7867258956", "Hi kk")
	sender03.SendAll([]string{"785736895", "785632342"}, "早上好")

	// 问题？
	// var sender04 Sender = &EmailSender{"smtp.qq.com", 465, "kk", "123456"} // ???
	// var sender05 Sender = SMSSender{"http://v.juhe.cn/sms/send", "1234"} // ???

	var ssender SingleSender = sender03
	ssender.Send("78523423423", "Hi KK")

	// 使用类型断言进行转换
	sender04, ok := ssender.(Sender)
	fmt.Printf("%T, %#v, %v\n", sender04, sender04, ok)
	sender03.SendAll([]string{"78234235234", "234234235"}, "早上好")

	emailSender, ok := sender01.(EmailSender)
	fmt.Printf("%T, %#v, $v\n", emailSender, emailSender, ok)
	if ok {
		fmt.Println(emailSender.addr)
	}

	if smsSender, ok := sender02.(*SMSSender); ok {
		fmt.Printf("%T, %#v, %v\n", smsSender, smsSender, ok)
		fmt.Println(smsSender.url)
	}

	if smsSender02, ok := sender03.(*SMSSender); !ok {
		fmt.Printf("%T, %v, %v\n", smsSender02, smsSender02, ok)
	}
}

PS E:\go-phase-two\go-course> go run E:\go-phase-two\go-course\1test.go
<nil>: <nil>main.EmailSender：{smtp.qq.com 465 kk 123456}
发送邮件给：imsilence@outlook.com， 内容： Hi Silence
发送邮件给imsilence@outlook.com, 内容： 早上好
发送邮件给iamkk@outlook.com, 内容： 早上好
*main.SMSSender: &{http://v.juhe.cn/sms/send 123}
发送短信给： 15200000000, 内容：Hi KK
发送短信给： 15200000000, 内容：早上好
发送短信给： 15800000000, 内容：早上好
*main.WeChatSender: &{/tmp.session}
发送短信给：7867258956, 内容：Hi kk
发送微信给：785736895, 内容：早上好
发送微信给：785632342, 内容：早上好
发送短信给：78523423423, 内容：Hi KK
*main.WeChatSender, &main.WeChatSender{sessionFile:"/tmp.session"}, true
发送微信给：78234235234, 内容：早上好
发送微信给：234234235, 内容：早上好
main.EmailSender, main.EmailSender{addr:"smtp.qq.com", port:465, user:"kk", password:"123456"}, $v
%!(EXTRA bool=true)smtp.qq.com
*main.SMSSender, &main.SMSSender{url:"http://v.juhe.cn/sms/send", key:"123"}, true
http://v.juhe.cn/sms/send
*main.SMSSender, <nil>, false

```

### 查询

可以通过switch-case+接口变量.(type查询变量类型,并选择对应的分支块)

```go
package main

import "fmt"

// 定义Sender接口
type Sender interface {
	Send(to, msg string) error
	SendAll(tos []string, msg string) error
}

// 定义EmailSender结构体
type EmailSender struct {
	addr           string
	port           int
	user, password string
}

func NewEmailSender(addr string, port int, user, password string) EmailSender {
	return EmailSender{addr, port, user, password}
}

// 接收者为值对象
func (sender EmailSender) Send(to, msg string) error {
	fmt.Printf("发送邮件给：%s， 内容： %s\n", to, msg)
	return nil
}

// 接收者为值对象
func (sender EmailSender) SendAll(tos []string, msg string) error {
	for _, to := range tos {
		fmt.Printf("发送邮件给%s, 内容： %s\n", to, msg)
	}
	return nil
}

type SMSSender struct {
	url, key string
}

func NewSMSSender(url, key string) *SMSSender {
	return &SMSSender{url, key}

}

// 接收者为指针对象
func (sender *SMSSender) Send(to, msg string) error {

	fmt.Printf("发送短信给： %s, 内容：%s\n", to, msg)
	return nil
}

// 接收者为指针对象
func (sender *SMSSender) SendAll(tos []string, msg string) error {
	for _, to := range tos {
		fmt.Printf("发送短信给： %s, 内容：%s\n", to, msg)
	}
	return nil
}

// 定义WeChatSender 结构体
type WeChatSender struct {
	sessionFile string
}

func NewWeChatSender(sessionFile string) *WeChatSender {
	return &WeChatSender{sessionFile}
}

// 接收者为值对象
func (sender WeChatSender) Send(to, msg string) error {
	fmt.Printf("发送短信给：%s, 内容：%s\n", to, msg)
	return nil
}

// 接收者为指针对象
func (sender *WeChatSender) SendAll(tos []string, msg string) error {
	for _, to := range tos {
		fmt.Printf("发送微信给：%s, 内容：%s\n", to, msg)
	}
	return nil
}

type SingleSender interface {
	Send(to, mgs string) error
}

func main() {

	var sender Sender
	fmt.Printf("%T: %v", sender, sender)

	// 赋值sender对象为NewEmailSender值对象
	var sender01 Sender = EmailSender{"smtp.qq.com", 465, "kk", "123456"}
	fmt.Printf("%T：%v\n", sender01, sender01)
	sender01.Send("imsilence@outlook.com", "Hi Silence")
	sender01.SendAll([]string{"imsilence@outlook.com", "iamkk@outlook.com"}, "早上好")

	// 赋值sender 对象为NewSMSSender指针对象
	var sender02 Sender = &SMSSender{"http://v.juhe.cn/sms/send", "123"}
	fmt.Printf("%T: %v\n", sender02, sender02)
	sender02.Send("15200000000", "Hi KK")
	sender02.SendAll([]string{"15200000000", "15800000000"}, "早上好")

	// 赋值应该使用哪一个？原因？
	// var sender03 Sender = WeChatSender{"/tmp/wx.session"} // ？？？
	// var sender03 Sender = &WeChatSender{"/tmp/wx.session"} // ？？？
	var sender03 Sender = &WeChatSender{"/tmp.session"}
	fmt.Printf("%T: %v\n", sender03, sender03)
	sender03.Send("7867258956", "Hi kk")
	sender03.SendAll([]string{"785736895", "785632342"}, "早上好")

	// 问题？
	// var sender04 Sender = &EmailSender{"smtp.qq.com", 465, "kk", "123456"} // ???
	// var sender05 Sender = SMSSender{"http://v.juhe.cn/sms/send", "1234"} // ???

	var ssender SingleSender = sender03
	ssender.Send("78523423423", "Hi KK")

	// 使用类型断言进行转换
	sender04, ok := ssender.(Sender)
	fmt.Printf("%T, %#v, %v\n", sender04, sender04, ok)
	sender03.SendAll([]string{"78234235234", "234234235"}, "早上好")

	emailSender, ok := sender01.(EmailSender)
	fmt.Printf("%T, %#v, $v\n", emailSender, emailSender, ok)
	if ok {
		fmt.Println(emailSender.addr)
	}

	if smsSender, ok := sender02.(*SMSSender); ok {
		fmt.Printf("%T, %#v, %v\n", smsSender, smsSender, ok)
		fmt.Println(smsSender.url)
	}

	if smsSender02, ok := sender03.(*SMSSender); !ok {
		fmt.Printf("%T, %v, %v\n", smsSender02, smsSender02, ok)
	}

	// 使用类型查询
	senders := []SingleSender{sender, ssender, sender01, sender02, sender03, sender04}
	for _, s := range senders {
		switch v := s.(type) {
		case EmailSender:
			fmt.Printf("EmailSender: %#v, %v\n", v, v.addr)
		case *SMSSender:
			fmt.Printf("SMSSender: %#v, %v\n", v, v.url)
		// case *WeChatSender
		// fmt.Printf("WeChatSender: %#v, %v\n", v, v.sessionFile)
		case Sender:
			fmt.Printf("Sender: %#v, %v\n", v, v)
			v.SendAll([]string{"one", "two"}, "sender")
		default:
			fmt.Printf("error, %#v\n", v)
		}
	}
}

PS E:\go-phase-two\go-course> go run E:\go-phase-two\go-course\1test.go
<nil>: <nil>main.EmailSender：{smtp.qq.com 465 kk 123456}
发送邮件给：imsilence@outlook.com， 内容： Hi Silence
发送邮件给imsilence@outlook.com, 内容： 早上好
发送邮件给iamkk@outlook.com, 内容： 早上好
*main.SMSSender: &{http://v.juhe.cn/sms/send 123}
发送短信给： 15200000000, 内容：Hi KK
发送短信给： 15200000000, 内容：早上好
发送短信给： 15800000000, 内容：早上好
*main.WeChatSender: &{/tmp.session}
发送短信给：7867258956, 内容：Hi kk
发送微信给：785736895, 内容：早上好
发送微信给：785632342, 内容：早上好
发送短信给：78523423423, 内容：Hi KK
*main.WeChatSender, &main.WeChatSender{sessionFile:"/tmp.session"}, true
发送微信给：78234235234, 内容：早上好
发送微信给：234234235, 内容：早上好
main.EmailSender, main.EmailSender{addr:"smtp.qq.com", port:465, user:"kk", password:"123456"}, $v
%!(EXTRA bool=true)smtp.qq.com
*main.SMSSender, &main.SMSSender{url:"http://v.juhe.cn/sms/send", key:"123"}, true
http://v.juhe.cn/sms/send
*main.SMSSender, <nil>, false
error, <nil>
Sender: &main.WeChatSender{sessionFile:"/tmp.session"}, &{/tmp.session}
发送微信给：one, 内容：sender
发送微信给：two, 内容：sender
EmailSender: main.EmailSender{addr:"smtp.qq.com", port:465, user:"kk", password:"123456"}, smtp.qq.com
SMSSender: &main.SMSSender{url:"http://v.juhe.cn/sms/send", key:"123"}, http://v.juhe.cn/sms/send
Sender: &main.WeChatSender{sessionFile:"/tmp.session"}, &{/tmp.session}
发送微信给：one, 内容：sender
发送微信给：two, 内容：sender
Sender: &main.WeChatSender{sessionFile:"/tmp.session"}, &{/tmp.session}
发送微信给：one, 内容：sender
发送微信给：two, 内容：sender
```



## 接口匿名嵌入

接口之中也可以嵌入已存在的接口，从而实现接口的扩展。

### 定义 

```go
package main

import "fmt"

// 定义Sender接口
type Sender interface {
	Send(msg string) error
}

// 定义Sender接口
type Receiver interface {
	Receiver() (string, error)
}

// 定义Connection接口， 由Sender和Receiver组合
type Connection interface {
	Sender
	Receiver
	Open() error
	Close() error
}

// 定义TCPConnection 结构体，并实现Connection接口中Open、Send、Receive、Close方法
type TCPConnection struct {
	addr string
	port int
}

func (conn TCPConnection) Open() error {
	fmt.Println("打开")
	return nil
}

func (conn TCPConnection) Send(msg string) error {
	fmt.Println("发送消息")
	return nil
}

func (conn TCPConnection) Receiver() (string, error) {
	fmt.Println("接收消息")
	return "", nil
}

func (conn TCPConnection) Close() error {
	fmt.Printf("关闭")
	return nil
}

func main() {

	var conn Connection = TCPConnection{"127.0.0.1", 8888}
	conn.Open()
	conn.Send("HI")
	msg, _ := conn.Receiver()
	fmt.Printf("%q\n", msg)
	conn.Close()
}

```

### 实现 

```go
package main

import "fmt"

// 定义Sender接口
type Sender interface {
	Send(msg string) error
}

// 定义Sender接口
type Receiver interface {
	Receiver() (string, error)
}

// 定义Connection接口， 由Sender和Receiver组合
type Connection interface {
	Sender
	Receiver
	Open() error
	Close() error
}

// 定义TCPConnection 结构体，并实现Connection接口中Open、Send、Receive、Close方法
type TCPConnection struct {
	addr string
	port int
}

func (conn TCPConnection) Open() error {
	fmt.Println("打开")
	return nil
}

func (conn TCPConnection) Send(msg string) error {
	fmt.Println("发送消息")
	return nil
}

func (conn TCPConnection) Receiver() (string, error) {
	fmt.Println("接收消息")
	return "", nil
}

func (conn TCPConnection) Close() error {
	fmt.Printf("关闭")
	return nil
}

func main() {

	var conn Connection = TCPConnection{"127.0.0.1", 8888}
	conn.Open()
	conn.Send("HI")
	msg, _ := conn.Receiver()
	fmt.Printf("%q\n", msg)
	conn.Close()
}

```

### 使用 

```go
package main

import "fmt"

// 定义Sender接口
type Sender interface {
	Send(msg string) error
}

// 定义Sender接口
type Receiver interface {
	Receiver() (string, error)
}

// 定义Connection接口， 由Sender和Receiver组合
type Connection interface {
	Sender
	Receiver
	Open() error
	Close() error
}

// 定义TCPConnection 结构体，并实现Connection接口中Open、Send、Receive、Close方法
type TCPConnection struct {
	addr string
	port int
}

func (conn TCPConnection) Open() error {
	fmt.Println("打开")
	return nil
}

func (conn TCPConnection) Send(msg string) error {
	fmt.Println("发送消息")
	return nil
}

func (conn TCPConnection) Receiver() (string, error) {
	fmt.Println("接收消息")
	return "", nil
}

func (conn TCPConnection) Close() error {
	fmt.Printf("关闭")
	return nil
}

func main() {

	var conn Connection = TCPConnection{"127.0.0.1", 8888}
	conn.Open()
	conn.Send("HI")
	msg, _ := conn.Receiver()
	fmt.Printf("%q\n", msg)
	conn.Close()
}


PS E:\go-phase-two\go-course> go run E:\go-phase-two\go-course\1test.go
打开
发送消息
接收消息
""
关闭
```


## 匿名接口 

在定义变量时将类型指定为接口的函数签名的接口，此时叫匿名接口。匿名接口常用于初始化一次接口变量的场景。


```go

package main

import "fmt"

// 定义Sender接口
type Sender interface {
	Send(msg string) error
}

// 定义Sender接口
type Receiver interface {
	Receiver() (string, error)
}

// 定义Connection接口， 由Sender和Receiver组合
type Connection interface {
	Sender
	Receiver
	Open() error
	Close() error
}

// 定义TCPConnection 结构体，并实现Connection接口中Open、Send、Receive、Close方法
type TCPConnection struct {
	addr string
	port int
}

func (conn TCPConnection) Open() error {
	fmt.Println("打开")
	return nil
}

func (conn TCPConnection) Send(msg string) error {
	fmt.Println("发送消息")
	return nil
}

func (conn TCPConnection) Receiver() (string, error) {
	fmt.Println("接收消息")
	return "", nil
}

func (conn TCPConnection) Close() error {
	fmt.Printf("关闭")
	return nil
}

func main() {

	var conn Connection = TCPConnection{"127.0.0.1", 8888}
	conn.Open()
	conn.Send("HI")
	msg, _ := conn.Receiver()
	fmt.Printf("%q\n", msg)
	conn.Close()

	var closer interface {
		Close() error
	}
	fmt.Printf("%T, %#v\n", closer, closer)
	closer = conn
	closer.Close()
}


PS E:\go-phase-two\go-course> go run E:\go-phase-two\go-course\1test.go
打开
发送消息
接收消息
""
关闭<nil>, <nil>
关闭
```


## 空接口 

不包含任何函数签名的接口叫做空接口，空接口声明的变量可以赋值为任何类型的变量（任意接口）

### 定义 

语法： interface{}

```go

package main

import "fmt"

func main() {
	var i01 interface{} = 1
	var i02 interface{} = true
	var i03 interface{} = "我是kk"
	var i04 interface{} = &i03
	var i05 interface{} = [...]int{1, 2, 3, 4, 5}
	var i06 interface{} = []int{1, 2, 3, 4, 5}
	var i07 interface{} = map[string]string{"name": "kk"}
	var i08 interface{} = struct{ X, Y int }{1, 2}

	fmt.Printf("%#v\n", i01)
	fmt.Printf("%#v\n", i02)
	fmt.Printf("%#v\n", i03)
	fmt.Printf("%#v\n", i04)
	fmt.Printf("%#v\n", i05)
	fmt.Printf("%#v\n", i06)
	fmt.Printf("%#v\n", i07)
	fmt.Printf("%#v\n", i08)
}

PS E:\go-phase-two\go-course> go run E:\go-phase-two\go-course\1test.go
1
true
"我是kk"
(*interface {})(0xc0000481f0)
[5]int{1, 2, 3, 4, 5}
[]int{1, 2, 3, 4, 5}
map[string]string{"name":"kk"}
struct { X int; Y int }{X:1, Y:2}
```


### 使用场景 

常声明函数参数类型为interface{}, 用于接收任意类型的变量

```go
// package main

// import "fmt"

// // 定义Sender接口
// type Sender interface {
// 	Send(msg string) error
// }

// // 定义Sender接口
// type Receiver interface {
// 	Receiver() (string, error)
// }

// // 定义Connection接口， 由Sender和Receiver组合
// type Connection interface {
// 	Sender
// 	Receiver
// 	Open() error
// 	Close() error
// }

// // 定义TCPConnection 结构体，并实现Connection接口中Open、Send、Receive、Close方法
// type TCPConnection struct {
// 	addr string
// 	port int
// }

// func (conn TCPConnection) Open() error {
// 	fmt.Println("打开")
// 	return nil
// }

// func (conn TCPConnection) Send(msg string) error {
// 	fmt.Println("发送消息")
// 	return nil
// }

// func (conn TCPConnection) Receiver() (string, error) {
// 	fmt.Println("接收消息")
// 	return "", nil
// }

// func (conn TCPConnection) Close() error {
// 	fmt.Printf("关闭")
// 	return nil
// }

// func main() {

// 	var conn Connection = TCPConnection{"127.0.0.1", 8888}
// 	conn.Open()
// 	conn.Send("HI")
// 	msg, _ := conn.Receiver()
// 	fmt.Printf("%q\n", msg)
// 	conn.Close()

// 	var closer interface {
// 		Close() error
// 	}
// 	fmt.Printf("%T, %#v\n", closer, closer)
// 	closer = conn
// 	closer.Close()
// }

package main

import "fmt"

func printType(vs ...interface{}) {
	for _, v := range vs {
		switch v.(type) {
		case nil:
			fmt.Println("nil")
		case int:
			fmt.Println("int")
		case bool:
			fmt.Println("bool")
		case string:
			fmt.Println("string")
		case [5]int:
			fmt.Println("[5]int")
		case []int:
			fmt.Println("[]int")
		case map[string]string:
			fmt.Println("map[string]string")
		default:
			fmt.Println("unknow")
		}
	}
}

func main() {
	var i01 interface{} = 1
	var i02 interface{} = true
	var i03 interface{} = "我是kk"
	var i04 interface{} = &i03
	var i05 interface{} = [...]int{1, 2, 3, 4, 5}
	var i06 interface{} = []int{1, 2, 3, 4, 5}
	var i07 interface{} = map[string]string{"name": "kk"}
	var i08 interface{} = struct{ X, Y int }{1, 2}

	fmt.Printf("%#v\n", i01)
	fmt.Printf("%#v\n", i02)
	fmt.Printf("%#v\n", i03)
	fmt.Printf("%#v\n", i04)
	fmt.Printf("%#v\n", i05)
	fmt.Printf("%#v\n", i06)
	fmt.Printf("%#v\n", i07)
	fmt.Printf("%#v\n", i08)

	printType(i01)
	printType(i02)
	printType(i03)
	printType(i04)
	printType(i05)
	printType(i05, i06, i07, i08, nil)
}


PS E:\go-phase-two\go-course> go run E:\go-phase-two\go-course\1test.go
1
true
"我是kk"
(*interface {})(0xc0000481f0)
[5]int{1, 2, 3, 4, 5}
[]int{1, 2, 3, 4, 5}
map[string]string{"name":"kk"}
struct { X int; Y int }{X:1, Y:2}
int
bool
string
unknow
[5]int
[5]int
[]int
map[string]string
unknow
nil
```



# 反射

反射是指在运行时动态的访问和修改任意类型对象的结构和成员，在go 语言中提供reflect包提供反射的功能，每一个变量都有两个属性：类型(Type)和值(Value)

## Type 

reflect.Type 是一个接口类型，用于获取变量类型的信息，可通过reflect.Type 函数获取某个变量的类型信息

### 通用方法

- Name(): 类型名 
- PkgPath()：包路径
- Kind(): 类型枚举值
- String(): Type字符串
- Comparable(): 是否可以进行比较
- Implements(Type): 是否实现某类型
- AssignableTo(Type): 是否可赋值给某类型
- ConvertibleTo(Type): 是否可转换为某类型
- NumMethod(): 方法个数
- Method(int): 通过索引获取方法类型
  Method结构体常用属性：
  - Name: 方法名 
  - Type：函数类型 
  - Func：方法值(Value)
- MethodByName(string): 通过方法名字获取方法类型 

### 特定类型方法 

1. reflect.Int*, reflect.UInt*, reflect.Float*, reflect. Complex*
   1. Bits(): 获取占用字节位数
2. reflact.Array
   1. Len(): 获取数组长度 
   2. Elem(): 获取数据元素类型
3. reflect.Slice
   1. Elem(): 获取切片元素类型
4. reflect.Map
   1. Key(): 获取映射键类型
   2. Elem(): 获取映射值类型
5. reflect.Ptr
   1. Elem(): 获取指向值类型
6. reflect.Func
   1. IsVariadic(): 是否具有可变参数
   2. NumIn(): 参数个数
   3. In(int): 通过索引获取参数类型
   4. NumOut: 返回值个数
   5. Out(int): 通过索引获取返回值类型
7. reflect.Struct
   1. NumField: 属性个数
   2. Field(int)：通过索引获取属性
   3. StructField 结构体常用属性: 
      1. Name:属性名
      2. Anonymous:是否为匿名
      3. Tag:标签
         1. StructTag 常用方法: 
            1. Get(string)
            2. Lookup(string)
   4. FieldByName(string): 通过属性名获取属性


## Value 

reflect.Value 是一个结构体类型，用于获取变量值的信息，可通过reflect.ValueOf函数获取某个变量的值信息。

### 创建方法

### 通用方法 

- Type(): 获取值类型
- CanAddr()：是否可获取地址
- Addr(): 获取地址
- CanInterface(): 是否可以获取接口的
- InterfaceData():
- Interface(): 将变量转换为 interface{}
- CanSet(): 是否可更新
- isValid(): 是否初始化为零值
- Kind()：获取值类型枚举值
- NumMethod(): 方法个数
- Method(int): 通过索引获取方法值
- MethodByName(string): 通过方法名字获取方法值
- Convert(Type)：转换为对应类型的值

### 修改方法 

- Set/Set*: 设置变量值 
  
### 调用方法 

- Call 
- CallSlice 

### 特定类型方法 

1. reflect.Int*, reflect.Unit*, 
   1. Int(): 获取对应类型值
   2. Unit(): 获取对应类型值
2. reflect.Float* 
   1. Float():获取对应类型值
3. reflect. Complex* 
   1. Complex():获取对应类型值
4. reflact.Array
   1. Len(): 获取数组长度
   2. Index(int): 根据索引获取元素
   3. Slice(int, int): 获取切片
   4. Slice3(int, int, int): 获取切片
5. reflect.Slice
   1. IsNil(): 判断是否为
   2. Len(): 获取元素数量
   3. Cap(): 获取容量
   4. Index(int): 根据索引获取元素
   5. Slice(int, int): 获取切片
   6. Slice3(int, int, int): 获取切片
6. reflect.Map
   1. IsNil(): 判断是否为
   2. Len(): 获取元素数量
   3. MapKeys(): 获取所有键
   4. MapIndex(Value): 根据键获取值
   5. MapRange():获取键值组成的可迭代对象
7. reflect.Ptr
   1. Elem(): 获取指向值类型（解引用）
8. reflect.Func
   1. IsVariadic(): 是否具有可变参数
   2. NumIn(): 参数个数 ⚫ In(int): 通过索引获取参数类型
   3. NumOut: 返回值个数
   4. Out(int): 通过索引获取返回值类型
9. reflect.Struct
   1.  NumField: 属性个数
   2.  Field(int)：通过索引获取属性
   3.  StructField 结构体常用属性: 
       1.  Name:属性名
       2.  Anonymous:是否为匿名
       3.  Tag:标签
       4.  StructTag 常用方法:
           1.  Get(string)
           2.  Lookup(string)
10. FieldByName(string): 通过属性名获取属性





## 开发 

objs包中创建用于测试的对象 

### 打印变量信息

```go
package main

import (
	"fmt"
	"reflect"
)

// 定义User结构体，并为每个属性定义标签

type User struct {
	id     int     `json:"id"`
	name   string  `json:"name"`
	Tel    string  `json:"addr"`
	Height float32 `json:"height"`
	Desc   *string `json:"desc"`
	Weight *int    `json:"weight"`
}

func NewUser(id int, name string, tel string, height float32, desc string, weight int) *User {
	return &User{id, name, tel, height, &desc, &weight}
}

// 定义String方法
func (user User) String() string {
	return fmt.Sprintf("User{id: %d, name: %s, tel: %s, height: %e, desc: %s, weight: %d}", user.id, user.name, user.Tel, user.Height, *user.Desc, *user.Weight)
}

func (user User) GetId() int {
	return user.id
}

func (user User) SetId(id int) {
	user.id = id
}

func (user *User) GetName() string {
	return user.name
}

func (user *User) SetName(name string) {
	user.name = name
}

// 定义接口Closer
type Closer interface {
	Close() error
}

type Address struct {
	ip   string "json:ip"
	port int    "json:port"
}

func (address Address) GetIp() string {
	return address.ip

}

func (address Address) GetPort() int {
	return address.port
}

type Connection struct {
	Address
	status int
}

func NewConnection(ip string, port int) *Connection {
	return &Connection{Address: Address{ip, port}}
}

func (conn *Connection) Send(msg string) error {
	fmt.Printf("发送消息给[%s:%d]:%s", conn.ip, conn.port, msg)
	return nil
}

func (conn *Connection) Close() error {
	fmt.Printf("关闭连接[%s:%d]", conn.ip, conn.port)
	return nil
}

// 打印变量类型信息
func displayType(t reflect.Type, tab string) {
	// 处理类型为nil
	if t == nil {
		fmt.Printf("<nil>")
		return
	}
	// 获取类型对应的枚举值使用选择语句分别处理每种类型
	switch t.Kind() {
	case reflect.Int, reflect.Float32, reflect.Bool, reflect.String:
		// 针对基本数据类型显示类型名
		fmt.Printf("%s%s", tab, t.Name())
	case reflect.Array, reflect.Slice:
		// 针对数组和切片，直接打印Type对象
		fmt.Printf("%s%s", tab, t)
	case reflect.Map:
		// 对于映射类型打印键和值的Type对象
		fmt.Printf("%smap{\n", tab)
		fmt.Printf("%s\tKey:", tab)
		fmt.Printf("%s%s", tab, t.Key()) // 获取键的Type对象
		fmt.Println()
		fmt.Printf("%s\tValue: ", tab)
		fmt.Printf("%s%s", tab, t.Elem()) // 获取值的Type对象
		fmt.Println()
		fmt.Printf("%s}", tab)
	case reflect.Func:
		// 针对函数打印参数和返回值，对应可变参数在最后一个参数添加之后添加。。。
		fmt.Printf("%sfunc (", tab)

		// 打印参数信息
		// 获取参数数量并遍历
		for i := 0; i < t.NumIn(); i++ {
			fmt.Printf("%s", t.In(i)) // 根据索引获取第i个参数的Type 对象
			if i != t.NumIn()-1 {
				fmt.Printf(", ")
			}
		}
		if t.IsVariadic() {
			fmt.Printf("...")
		}
		fmt.Printf(") ")

		// 打印返回值信息
		if t.NumOut() > 0 {
			fmt.Printf("( ")
			for i := 0; i < t.NumOut(); i++ {
				fmt.Printf("%s", t.Out(i)) // 根据索引获取第i个返回值的Type对象
				if i != t.NumOut()-1 {
					fmt.Printf(", ")
				}
			}
			fmt.Printf(" )")
		}
		fmt.Printf("{}")
	case reflect.Struct:
		// 针对结构体显示结构体属性和方法
		fmt.Printf("%stype %s struct {\n", tab, t.Name())

		// 获取属性数量并遍历
		fmt.Printf("%s\tFields(%d):\n", tab, t.NumField())
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)                                                      // 根据索引获取第i个属性的StructField对象
			fmt.Printf("%s\t\t%s\t%s\t`%s`", tab, field.Name, field.Type, field.Tag) // 显示属性名，属性的Type对象和标签
			fmt.Printf(",\n")
		}

		// 获取方法数量并遍历
		fmt.Printf("\n%s\tMethods(%d):\n", tab, t.NumMethod())
		for i := 0; i < t.NumMethod(); i++ {
			// 打印方法信息
			displayMethod(t.Method(i), tab+"\t\t")
			fmt.Printf(",\n")
		}
		fmt.Printf("%s}", tab)
	case reflect.Ptr:
		// 对于指针类型，递归分析其引用值
		fmt.Printf("%s*{\n", tab)
		displayType(t.Elem(), tab+"\t") // 获取指针的引用值，并递归调用displayType进行分析显示
		// 打印指针变量的方法
		if t.NumMethod() > 0 { // 获取方法数量
			fmt.Printf("\n%s\tMethods(%d):\n", tab, t.NumMethod())
			for i := 0; i < t.NumMethod(); i++ {
				// 打印方法信息
				displayMethod(t.Method(i), tab+"\t\t")
				if i != t.NumMethod()-1 {
					fmt.Printf(",\n")
				}
			}
		}
		fmt.Printf("\n%s", tab)
	default:
		fmt.Printf("%sUnkonw[%s]", tab, t)
	}
}

// 打印reflect.Method类型变量method的信息
func displayMethod(method reflect.Method, tab string) {
	// 获取方法接收者
	t := method.Type
	fmt.Printf("%sfunc %s(", tab, method.Name) // 显示方法名

	// 打印参数信息
	// 获取参数数量并遍历
	for i := 0; i < t.NumIn(); i++ {
		fmt.Printf("%s", t.In(i)) // 根据索引获取第i个参数的Type对象
		if i != t.NumIn()-1 {
			fmt.Printf(", ")
		}
	}

	// 打印可变参数信息
	if t.IsVariadic() {
		fmt.Printf("...")
	}

	// 打印返回值信息
	fmt.Printf(") ")

	if t.NumOut() > 0 {
		fmt.Printf("(")
		// 获取返回值数量并遍历
		for i := 0; i < t.NumOut(); i++ {
			fmt.Printf("%s", t.Out(i)) // 根据索引获取第i个返回值的Type对象
			if i != t.NumOut()-1 {
				fmt.Printf(", ")
			}
		}
		fmt.Printf(") ")
	}
	fmt.Printf("{}")
}

func main() {
	vars := make([]interface{}, 0, 20)

	var intV int = 1
	var floatV float32 = 3.14
	var boolV bool = true
	var stringV string = "吾日三省吾身，为人谋而不忠乎？与朋友交而不信呼？传不习呼？"
	var arrayV [5]int = [...]int{1, 2, 3, 4, 5}
	var sliceV []int = make([]int, 3, 5)
	var mapV map[string]string = map[string]string{"name": "kk"}
	var funcV1 func(...interface{}) error = func(x ...interface{}) error { fmt.Println(x...); return nil }
	var funcV2 func(string, int) *Connection = NewConnection
	var userV *User = NewUser(1, "kk", "12000000000", 1.68, "少年经不得顺境，中间经不得闲境，晚年经不得逆境", 72)
	var CloserV Closer
	vars = append(vars, intV, &intV, floatV, boolV, stringV, arrayV, sliceV, mapV, funcV1, funcV2, *userV, userV, CloserV)
	for _, v := range vars {
		displayType(reflect.TypeOf(v), "")
		fmt.Println()
	}
}



PS E:\go-phase-two\go-course> go run E:\go-phase-two\go-course\1test.go
int
*{
        int

float32
bool
string
[5]int
[]int
map{
        Key:string
        Value: string
}
func ([]interface {}...) ( error ){}
func (string, int) ( *main.Connection ){}
type User struct {
        Fields(6):
                id      int     `json:"id"`,
                name    string  `json:"name"`,
                Tel     string  `json:"addr"`,
                Height  float32 `json:"height"`,
                Desc    *string `json:"desc"`,
                Weight  *int    `json:"weight"`,

        Methods(3):
                func GetId(main.User) (int) {},
                func SetId(main.User, int) {},
                func String(main.User) (string) {},
}
*{
        type User struct {
                Fields(6):
                        id      int     `json:"id"`,
                        name    string  `json:"name"`,
                        Tel     string  `json:"addr"`,
                        Height  float32 `json:"height"`,
                        Desc    *string `json:"desc"`,
                        Weight  *int    `json:"weight"`,

                Methods(3):
                        func GetId(main.User) (int) {},
                        func SetId(main.User, int) {},
                        func String(main.User) (string) {},
        }
        Methods(5):
                func GetId(*main.User) (int) {},
                func GetName(*main.User) (string) {},
                func SetId(*main.User, int) {},
                func SetName(*main.User, string) {},
                func String(*main.User) (string) {}

<nil>
```

### 打印变量值信息

```go
package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// 定义User结构体，并为每个属性定义标签

type User struct {
	id     int     `json:"id"`
	name   string  `json:"name"`
	Tel    string  `json:"addr"`
	Height float32 `json:"height"`
	Desc   *string `json:"desc"`
	Weight *int    `json:"weight"`
}

func NewUser(id int, name string, tel string, height float32, desc string, weight int) *User {
	return &User{id, name, tel, height, &desc, &weight}
}

// 定义String方法
func (user User) String() string {
	return fmt.Sprintf("User{id: %d, name: %s, tel: %s, height: %e, desc: %s, weight: %d}", user.id, user.name, user.Tel, user.Height, *user.Desc, *user.Weight)
}

func (user User) GetId() int {
	return user.id
}

func (user User) SetId(id int) {
	user.id = id
}

func (user *User) GetName() string {
	return user.name
}

func (user *User) SetName(name string) {
	user.name = name
}

// 定义接口Closer
type Closer interface {
	Close() error
}

type Address struct {
	ip   string "json:ip"
	port int    "json:port"
}

func (address Address) GetIp() string {
	return address.ip

}

func (address Address) GetPort() int {
	return address.port
}

type Connection struct {
	Address
	status int
}

func NewConnection(ip string, port int) *Connection {
	return &Connection{Address: Address{ip, port}}
}

func (conn *Connection) Send(msg string) error {
	fmt.Printf("发送消息给[%s:%d]:%s", conn.ip, conn.port, msg)
	return nil
}

func (conn *Connection) Close() error {
	fmt.Printf("关闭连接[%s:%d]", conn.ip, conn.port)
	return nil
}

// 打印reflect.Value类型变量v的信息
func displayValue(value reflect.Value, tab string) {

	// 获取值对应的枚举类型使用选择语句分别处理每种类型
	switch value.Kind() {
	case reflect.Int:
		// 针对整数类型，获取值对应的Type对象并使用Int函数转换为对应的基本类型，并使用strconv将转换为string类型
		fmt.Printf("%s[%s] %s", tab, value.Type(), strconv.FormatInt(value.Int(), 10))
	case reflect.Float32:
		// 针对浮点类型，获取值对应的Type对象并使用Float函数转换为对应的基本类型，并使用strconv将转化为string类型
		fmt.Printf("%s[%s] %s", tab, value.Type(), strconv.FormatFloat(value.Float(), 'E', -1, 64))
	case reflect.Bool:
		// 针对布尔类型，获取值对应的Type对象并使用Bool函数转换为对应的基本类型，并使用strconv将转化为string类型
		fmt.Printf("%s[%s] %s", tab, value.Type(), strconv.FormatBool(value.Bool()))
	case reflect.String:
		// 针对字符集类型，获取值对应的Type对象并打印值
		fmt.Printf("%s[%s] %s", tab, value.Type(), value)
	case reflect.Array:
		// 针对数组类型，获取值对应的Type对象
		fmt.Printf("%s[%s] {\n", tab, value.Type())

		// 获取数组的长度
		for i := 0; i < value.Len(); i++ {
			displayValue(value.Index(i), tab+"\t") // 根据索引获取数组的每个元素，并调用displayValue递归显示
			fmt.Printf(",\n")
		}
		fmt.Printf("%s}", tab)
	case reflect.Slice:
		// 针对切片类型，获取值对应的Type对象，长度和容量
		fmt.Printf("%s[%s](%d:%d) {\n", tab, value.Type(), value.Len(), value.Cap())

		// 获取切片的长度
		for i := 0; i < value.Len(); i++ {
			displayValue(value.Index(i), tab+"\t") // 根据索引获取数组的每个元素，并调用displayValue递归显示
			fmt.Printf(",\n")
		}
		fmt.Printf("%s}", tab)
	case reflect.Map:
		// 针对映射类型，获取值对应的Type对象
		fmt.Printf("%s[%s] {\n", tab, value.Type())
		// 获取映射迭代对象遍历键值对
		iter := value.MapRange()
		for iter.Next() { // 判断迭代对象是否到末尾
			displayValue(iter.Key(), tab+"\t") // 根据从迭代对象获取当前键，并调用displayValue递归显示
			fmt.Printf(" : ")
			displayValue(iter.Value(), "") // 根据从迭代对象获取当前值，并调用displayValue递归显示
			fmt.Printf(",\n")
		}
		fmt.Printf("%s}", tab)
	case reflect.Struct:
		// 针对结构体类型，获取值对应的Type对象
		structType := value.Type()
		fmt.Printf("%s[%s] {\n", tab, structType)

		// 获取属性数量并遍历
		for i := 0; i < value.NumField(); i++ {
			structField := structType.Field(i)           // 根据索引获取属性的类型
			field := value.Field(i)                      // 根据索引获取属性的值
			fmt.Printf("%s\t%s:", tab, structField.Name) // 打印类型名
			displayValue(field, tab+"\t")                // 调用displayValue递归显示
			fmt.Printf(", \n")

		}
		fmt.Printf("%s}", tab)

	case reflect.Ptr:
		// 针对指针类型，获取值对应的Type对象
		fmt.Printf("%s[%s] (\n", tab, value.Type())

		// 获取指针的引用值，并递归调用displayValue函数递归分析显示
		displayValue(value.Elem(), tab+"\t")
		fmt.Printf("\n%s", tab)
	default:
		fmt.Printf("%sUnkonw[%#v]", tab, value)
	}
}

func main() {
	vars := make([]interface{}, 0, 20)

	var intV int = 1
	var floatV float32 = 3.14
	var boolV bool = true
	var stringV string = "吾日三省吾身，为人谋而不忠乎？与朋友交而不信呼？传不习呼？"
	var arrayV [5]int = [...]int{1, 2, 3, 4, 5}
	var sliceV []int = make([]int, 3, 5)
	var mapV map[string]string = map[string]string{"name": "kk"}
	var userV *User = NewUser(1, "kk", "12000000000", 1.68, "少年经不得顺境，中间经不得闲境，晚年经不得逆境", 72)
	var CloserV Closer
	vars = append(vars, intV, &intV, floatV, boolV, stringV, arrayV, sliceV, mapV, *userV, userV, CloserV)
	for _, v := range vars {
		displayValue(reflect.ValueOf(v), "")
		fmt.Println()
	}
}




PS E:\go-phase-two\go-course> go run E:\go-phase-two\go-course\1test.go
[int] 1
[*int] (
        [int] 1

[float32] 3.140000104904175E+00
[bool] true
[string] 吾日三省吾身，为人谋而不忠乎？与朋友交而不信呼？传不习呼？
[[5]int] {
        [int] 1,
        [int] 2,
        [int] 3,
        [int] 4,
        [int] 5,
}
[[]int](3:5) {
        [int] 0,
        [int] 0,
        [int] 0,
}
[map[string]string] {
        [string] name : [string] kk,
}
[main.User] {
        id:     [int] 1,
        name:   [string] kk,
        Tel:    [string] 12000000000,
        Height: [float32] 1.6799999475479126E+00,
        Desc:   [*string] (
                [string] 少年经不得顺境，中间经不得闲境，晚年经不得逆境
        ,
        Weight: [*int] (
                [int] 72
        ,
}
[*main.User] (
        [main.User] {
                id:             [int] 1,
                name:           [string] kk,
                Tel:            [string] 12000000000,
                Height:         [float32] 1.6799999475479126E+00,
                Desc:           [*string] (
                        [string] 少年经不得顺境，中间经不得闲境，晚年经不得逆境
                ,
                Weight:         [*int] (
                        [int] 72
                ,
        }

Unkonw[<invalid reflect.Value>]
```

### 更新变量值

```go
package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

// 定义User结构体，并为每个属性定义标签

type User struct {
	id     int     `json:"id"`
	name   string  `json:"name"`
	Tel    string  `json:"addr"`
	Height float32 `json:"height"`
	Desc   *string `json:"desc"`
	Weight *int    `json:"weight"`
}

func NewUser(id int, name string, tel string, height float32, desc string, weight int) *User {
	return &User{id, name, tel, height, &desc, &weight}
}

// 定义String方法
func (user User) String() string {
	return fmt.Sprintf("User{id: %d, name: %s, tel: %s, height: %e, desc: %s, weight: %d}", user.id, user.name, user.Tel, user.Height, *user.Desc, *user.Weight)
}

func (user User) GetId() int {
	return user.id
}

func (user User) SetId(id int) {
	user.id = id
}

func (user *User) GetName() string {
	return user.name
}

func (user *User) SetName(name string) {
	user.name = name
}

// 定义接口Closer
type Closer interface {
	Close() error
}

type Address struct {
	ip   string "json:ip"
	port int    "json:port"
}

func (address Address) GetIp() string {
	return address.ip

}

func (address Address) GetPort() int {
	return address.port
}

type Connection struct {
	Address
	status int
}

func NewConnection(ip string, port int) *Connection {
	return &Connection{Address: Address{ip, port}}
}

func (conn *Connection) Send(msg string) error {
	fmt.Printf("发送消息给[%s:%d]:%s", conn.ip, conn.port, msg)
	return nil
}

func (conn *Connection) Close() error {
	fmt.Printf("关闭连接[%s:%d]", conn.ip, conn.port)
	return nil
}

// 修改reflect.Value类型变量value的信息
func changeValue(value reflect.Value, path string) {

	// 获取值对应的枚举类型使用选择语句分别处理每种类型
	switch value.Kind() {
	case reflect.Int:
		// 对于可设置的Int类型使用SetInt更新内存数据
		if value.CanSet() {
			fmt.Printf("Int CanSet: %s.%s\n", path, value.Type())
			value.SetInt(value.Int() + rand.Int63n(100))
		}
	case reflect.Float32:
		// 对于可设置的Float类型使用SetFloat更新内存数据
		if value.CanSet() {
			fmt.Printf("Float CanSet: %s.%s\n", path, value.Type())
			value.SetFloat(value.Float() + rand.Float64())
		}
	case reflect.Bool:
		// 对于可设置的Float类型使用SetInt更新内存数据
		if value.CanSet() {
			fmt.Printf("Bool CanSet: %s.%s\n", path, value.Type())
			value.SetBool(!value.Bool())
		}
	case reflect.String:
		// 对于可设置的String类型使用SetString更新内存数据
		if value.CanSet() {
			fmt.Printf("String CanSet: %s.%s\n", path, value.Type())
			value.SetString("change: " + value.String())
		}
	case reflect.Array:
		// 对于可设置的数据类型提柜调用ChangeValue对每个元素更新内存数据
		if value.CanSet() {
			fmt.Printf("Array CanSet: %s.%s\n", path, value.Type())
		}
		// 获取数组的长度
		for i := 0; i < value.Len(); i++ {
			changeValue(value.Index(i), path+".array") // 根据索引获取数组的每个元素，并调用changeValue递归显示
		}
	case reflect.Slice:
		// 针对切片类型递归调用ChangeValue对每个元素更新内存数据
		for i := 0; i < value.Len(); i++ {
			changeValue(value.Index(i), path+".slice") // 根据索引获取数组的每个元素，并调用changeValue递归显示
		}
	case reflect.Map:
		// 针对映射类型，获取值对应的Type对象
		fmt.Printf("%s[%s] {\n", path, value.Type())
		// 获取映射迭代对象遍历键值对
		keys := value.MapKeys() // 获取映射所有Key组成的Value切片
		for _, key := range keys {
			value.SetMapIndex(key, reflect.ValueOf("change: "+value.MapIndex(key).String()))
		}
	case reflect.Struct:
		if value.CanSet() {
			fmt.Printf("Struct CanSet: %s.%s\n", path, value.Type())
		}
		// 对于结构体类型递归调用ChangeValue对每个元素更新内存数据
		for i := 0; i < value.NumField(); i++ {
			changeValue(value.Field(i), path+".struct") // 根据索引获取结构体的每个属性，并调用changeValue递归修改
		}
	case reflect.Ptr:
		if value.CanSet() {
			fmt.Printf("CanSet: %s.%s\n", path, value.Type())
		}
		// 对于指针类型，递归调用对其引用值进行修改
		changeValue(value.Elem(), path+".pointer")
	default:
		fmt.Printf("%sUnkonw[%#v]", path, value)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	vars := make([]interface{}, 0, 20)

	var intV int = 1
	var floatV float32 = 3.14
	var boolV bool = true
	var stringV string = "吾日三省吾身，为人谋而不忠乎？与朋友交而不信呼？传不习呼？"
	var arrayV [5]int = [...]int{1, 2, 3, 4, 5}
	var sliceV []int = make([]int, 3, 5)
	var mapV map[string]string = map[string]string{"name": "kk"}
	var userV *User = NewUser(1, "kk", "12000000000", 1.68, "少年经不得顺境，中间经不得闲境，晚年经不得逆境", 72)
	var closerV Closer
	vars = append(vars, intV, &intV, floatV, boolV, stringV, arrayV, sliceV, mapV, *userV, userV, closerV)
	// intV, floatV, boolV, stringV, stringV, arrayV 不可修改
	// &intV, &floatV, &boolV, &stringV, &stringV, &arrayV 可修改， 指针类型通过Elem获取的引用值可获取地址

	vars = append(vars, sliceV, mapV, *userV, userV, closerV)
	// sliceV, mapV 可修改
	// userV 通过Elem获取的引用值可获取地址且公开的属性修改
	// *userV 公开属性可修改（结构体通过Elem获取的引用值可获取地址）
	for _, v := range vars {
		fmt.Println("------------------------------")
		fmt.Printf("%v\n", v)
		vv := reflect.ValueOf(v)
		changeValue(vv, "")
		fmt.Printf("%v\n", reflect.Indirect(vv))
	}
}





PS E:\go-phase-two\go-course> go run E:\go-phase-two\go-course\1test.go
------------------------------
1
1
------------------------------
0xc0000100a0
Int CanSet: .pointer.int
7
------------------------------
3.14
3.14
------------------------------
true
true
------------------------------
吾日三省吾身，为人谋而不忠乎？与朋友交而不信呼？传不习呼？
吾日三省吾身，为人谋而不忠乎？与朋友交而不信呼？传不习呼？
------------------------------
[1 2 3 4 5]
[1 2 3 4 5]
------------------------------
[0 0 0]
Int CanSet: .slice.int
Int CanSet: .slice.int
Int CanSet: .slice.int
[87 71 84]
------------------------------
map[name:kk]
[map[string]string] {
map[name:change: kk]
------------------------------
User{id: 1, name: kk, tel: 12000000000, height: 1.680000e+00, desc: 少年经不得顺境，中间经不得闲境，晚年经不得逆境, weight: 72}
String CanSet: .struct.pointer.string
Int CanSet: .struct.pointer.int
User{id: 1, name: kk, tel: 12000000000, height: 1.680000e+00, desc: change: 少年经不得顺境，中间经不得闲境，晚年经不得逆
境, weight: 76}
------------------------------
User{id: 1, name: kk, tel: 12000000000, height: 1.680000e+00, desc: change: 少年经不得顺境，中间经不得闲境，晚年经不得逆
境, weight: 76}
Struct CanSet: .pointer.main.User
String CanSet: .pointer.struct.string
Float CanSet: .pointer.struct.float32
CanSet: .pointer.struct.*string
String CanSet: .pointer.struct.pointer.string
CanSet: .pointer.struct.*int
Int CanSet: .pointer.struct.pointer.int
User{id: 1, name: kk, tel: change: 12000000000, height: 2.397326e+00, desc: change: change: 少年经不得顺境，中间经不得闲
境，晚年经不得逆境, weight: 122}
------------------------------
<nil>
Unkonw[<invalid reflect.Value>]<invalid reflect.Value>
------------------------------
[87 71 84]
Int CanSet: .slice.int
Int CanSet: .slice.int
Int CanSet: .slice.int
[125 150 130]
------------------------------
map[name:change: kk]
[map[string]string] {
map[name:change: change: kk]
------------------------------
User{id: 1, name: kk, tel: 12000000000, height: 1.680000e+00, desc: change: change: 少年经不得顺境，中间经不得闲境，晚年
经不得逆境, weight: 122}
String CanSet: .struct.pointer.string
Int CanSet: .struct.pointer.int
User{id: 1, name: kk, tel: 12000000000, height: 1.680000e+00, desc: change: change: change: 少年经不得顺境，中间经不得闲
境，晚年经不得逆境, weight: 217}
------------------------------
User{id: 1, name: kk, tel: change: 12000000000, height: 2.397326e+00, desc: change: change: change: 少年经不得顺境，中间
经不得闲境，晚年经不得逆境, weight: 217}
Struct CanSet: .pointer.main.User
String CanSet: .pointer.struct.string
Float CanSet: .pointer.struct.float32
CanSet: .pointer.struct.*string
String CanSet: .pointer.struct.pointer.string
CanSet: .pointer.struct.*int
Int CanSet: .pointer.struct.pointer.int
User{id: 1, name: kk, tel: change: change: 12000000000, height: 2.994064e+00, desc: change: change: change: change: 少年
经不得顺境，中间经不得闲境，晚年经不得逆境, weight: 236}
------------------------------
<nil>
Unkonw[<invalid reflect.Value>]<invalid reflect.Value>
```

变量可修改的条件：
1. 对于基本数据类型变量可获取地址
2. 对于结构体属性必须是可获取地址且为公开的属性




### 调用函数 


```go
package main

import (
	"fmt"
	"reflect"
)

// 定义User结构体，并为每个属性定义标签

type User struct {
	id     int     `json:"id"`
	name   string  `json:"name"`
	Tel    string  `json:"addr"`
	Height float32 `json:"height"`
	Desc   *string `json:"desc"`
	Weight *int    `json:"weight"`
}

func NewUser(id int, name string, tel string, height float32, desc string, weight int) *User {
	return &User{id, name, tel, height, &desc, &weight}
}

// 定义String方法
func (user User) String() string {
	return fmt.Sprintf("User{id: %d, name: %s, tel: %s, height: %e, desc: %s, weight: %d}", user.id, user.name, user.Tel, user.Height, *user.Desc, *user.Weight)
}

func (user User) GetId() int {
	return user.id
}

func (user User) SetId(id int) {
	user.id = id
}

func (user *User) GetName() string {
	return user.name
}

func (user *User) SetName(name string) {
	user.name = name
}

// 定义接口Closer
type Closer interface {
	Close() error
}

type Address struct {
	ip   string "json:ip"
	port int    "json:port"
}

func (address Address) GetIp() string {
	return address.ip

}

func (address Address) GetPort() int {
	return address.port
}

type Connection struct {
	Address
	status int
}

func NewConnection(ip string, port int) *Connection {
	return &Connection{Address: Address{ip, port}}
}

func (conn *Connection) Send(msg string) error {
	fmt.Printf("发送消息给[%s:%d]:%s", conn.ip, conn.port, msg)
	return nil
}

func (conn *Connection) Close() error {
	fmt.Printf("关闭连接[%s:%d]", conn.ip, conn.port)
	return nil
}

// 修改reflect.Value类型变量value的信息
func callValue(value reflect.Value) {

	fmt.Println("-----------------------------", value.Type())
	// 获取值对应的枚举类型使用选择语句分别处理每种类型
	switch value.Kind() {
	case reflect.Array, reflect.Slice:
		// 针对数组/切片类型，对数组/切片中元素进行递归处理
		for i := 0; i < value.Len(); i++ {
			callValue(value.Index(i)) // 根据索引获取数组的每个元素，并调用callValue递归调用

		}
	case reflect.Map:
		// 针对映射类型，对映射中值进行递归处理
		iter := value.MapRange()
		for iter.Next() {
			callValue(iter.Value()) // 根据切片获取数组的每个元素，并调用callValue递归调用
		}

	case reflect.Struct:
		// 针对结构体类型，执行所有方法
		for i := 0; i < value.NumMethod(); i++ {
			callMethod(value.Method(i), value.Type(), value.Type().Method(i))
		}
	case reflect.Ptr:
		// 针对指针类型，执行所有方法
		for i := 0; i < value.NumMethod(); i++ {
			callMethod(value.Method(i), value.Elem().Type(), value.Type().Method(i))
		}
	case reflect.Func:
		// 针对方法类型，执行方法
		callFunc(value)
	default:
		fmt.Printf("Unkonw[%#v]\n", value)
	}
}

// 调用函数
func callFunc(f reflect.Value) {
	// 获取函数类型并打印
	ftype := f.Type()
	fmt.Printf("%s\n", ftype.String())

	// 组装函数参数
	parameters := make([]reflect.Value, 0)
	for i := 0; i < ftype.NumIn(); i++ {
		parameters = append(parameters, reflect.Zero(ftype.In(i))) // 通过reflect.Zero创建参数同类型零值

	}

	// 调用函数并打印执行结果
	if ftype.IsVariadic() {
		fmt.Printf("%#v", f.CallSlice(parameters))

	} else {
		fmt.Println(f.Call(parameters))
	}
}

// 调用方法
func callMethod(f reflect.Value, t reflect.Type, m reflect.Method) {
	// 获取函数类型并打印
	ftype := f.Type()
	fmt.Printf("method: %s.%s => %s\n", t.Name(), m.Name, ftype.String())

	// 组装函数参数
	parameters := make([]reflect.Value, 0)
	for i := 0; i < ftype.NumIn(); i++ {
		parameters = append(parameters, reflect.New(ftype.In(i)).Elem()) // 通过reflect.New创建参数同类型零值的指针
	}
	// 调用函数并打印执行结果
	if ftype.IsVariadic() {
		fmt.Println(f.CallSlice(parameters))
	} else {
		fmt.Println(f.Call(parameters))
	}
}

func main() {
	vars := make([]interface{}, 0, 20)

	var funcs []func() = make([]func(), 0)
	funcs = append(funcs, func() { fmt.Println(1) }, func() { fmt.Println(2) }, func() { fmt.Println(3) })

	var funcV1 func(...interface{}) error = func(x ...interface{}) error { fmt.Println(x...); return nil }
	var funcV2 func(string, int) *Connection = NewConnection
	var userV *User = NewUser(1, "kk", "15200000000", 1.68, "少年经不得顺境，中年经不得闲境，晚年经不得逆境", 72)

	vars = append(vars, funcs, funcV1, funcV2, *userV, userV)
	for _, v := range vars {
		callValue(reflect.ValueOf(v))
	}
}





PS E:\go-phase-two\go-course> go run E:\go-phase-two\go-course\1test.go
----------------------------- []func()
----------------------------- func()
func()
1
[]
----------------------------- func()
func()
2
[]
----------------------------- func()
func()
3
[]
----------------------------- func(...interface {}) error
func(...interface {}) error

[]reflect.Value{reflect.Value{typ:(*reflect.rtype)(0x4f2f80), ptr:(unsafe.Pointer)(0xc000070528), flag:0x94}}----------------------------- func(string, int) *main.Connection
func(string, int) *main.Connection
[<*main.Connection Value>]
----------------------------- main.User
method: User.GetId => func() int
[<int Value>]
method: User.SetId => func(int)
[]
method: User.String => func() string
[User{id: 1, name: kk, tel: 15200000000, height: 1.680000e+00, desc: 少年经不得顺境，中年经不得闲境，晚年经不得逆境, weight: 72}]
----------------------------- *main.User
method: User.GetId => func() int
[<int Value>]
method: User.GetName => func() string
[kk]
method: User.SetId => func(int)
[]
method: User.SetName => func(string)
[]
method: User.String => func() string
[User{id: 1, name: , tel: 15200000000, height: 1.680000e+00, desc: 少年经不得顺境，中年经不得闲境，晚年经不得逆境, weight: 72}]

```

### 动态创建结构体

```go
package main

import (
	"fmt"
	"reflect"
)

// 定义User结构体，并为每个属性定义标签

type User struct {
	id     int     `json:"id"`
	name   string  `json:"name"`
	Tel    string  `json:"addr"`
	Height float32 `json:"height"`
	Desc   *string `json:"desc"`
	Weight *int    `json:"weight"`
}

func NewUser(id int, name string, tel string, height float32, desc string, weight int) *User {
	return &User{id, name, tel, height, &desc, &weight}
}

// 定义String方法
func (user User) String() string {
	return fmt.Sprintf("User{id: %d, name: %s, tel: %s, height: %e, desc: %s, weight: %d}", user.id, user.name, user.Tel, user.Height, *user.Desc, *user.Weight)
}

func (user User) GetId() int {
	return user.id
}

func (user User) SetId(id int) {
	user.id = id
}

func (user *User) GetName() string {
	return user.name
}

func (user *User) SetName(name string) {
	user.name = name
}

// 定义接口Closer
type Closer interface {
	Close() error
}

type Address struct {
	ip   string "json:ip"
	port int    "json:port"
}

func (address Address) GetIp() string {
	return address.ip

}

func (address Address) GetPort() int {
	return address.port
}

type Connection struct {
	Address
	status int
}

func NewConnection(ip string, port int) *Connection {
	return &Connection{Address: Address{ip, port}}
}

func (conn *Connection) Send(msg string) error {
	fmt.Printf("发送消息给[%s:%d]:%s", conn.ip, conn.port, msg)
	return nil
}

func (conn *Connection) Close() error {
	fmt.Printf("关闭连接[%s:%d]", conn.ip, conn.port)
	return nil
}

// 调用函数
func callFunc(f reflect.Value) {
	// 获取函数类型并打印
	ftype := f.Type()
	fmt.Printf("%s\n", ftype.String())

	// 组装函数参数
	parameters := make([]reflect.Value, 0)
	for i := 0; i < ftype.NumIn(); i++ {
		parameters = append(parameters, reflect.Zero(ftype.In(i))) // 通过reflect.Zero创建参数同类型零值

	}

	// 调用函数并打印执行结果
	if ftype.IsVariadic() {
		fmt.Printf("%#v", f.CallSlice(parameters))

	} else {
		fmt.Println(f.Call(parameters))
	}
}

// 调用方法
func callMethod(f reflect.Value, t reflect.Type, m reflect.Method) {
	// 获取函数类型并打印
	ftype := f.Type()
	fmt.Printf("method: %s.%s => %s\n", t.Name(), m.Name, ftype.String())

	// 组装函数参数
	parameters := make([]reflect.Value, 0)
	for i := 0; i < ftype.NumIn(); i++ {
		parameters = append(parameters, reflect.New(ftype.In(i)).Elem()) // 通过reflect.New创建参数同类型零值的指针
	}
	// 调用函数并打印执行结果
	if ftype.IsVariadic() {
		fmt.Println(f.CallSlice(parameters))
	} else {
		fmt.Println(f.Call(parameters))
	}
}

func main() {
	// 定义结构体属性
	fields := []reflect.StructField{
		{
			Name: "Name",
			Type: reflect.TypeOf(""),
		},
		{
			Name: "Score",
			Type: reflect.TypeOf(int64(0)),
		},
	}
	// 定义结构体类型
	UserType := reflect.StructOf(fields)

	// 创建结构体对象
	user := reflect.New(UserType).Elem()

	// 设置属性值
	user.Field(0).Set(reflect.ValueOf("Silence"))
	user.Field(1).Set(reflect.ValueOf(int64(10000)))

	// 打印对象和对象指针
	fmt.Printf("%#v\n", user.Interface())
	fmt.Printf("%#v\n", user.Addr().Interface())
}




PS E:\go-phase-two\go-course> go run E:\go-phase-two\go-course\1test.go
struct { Name string; Score int64 }{Name:"Silence", Score:10000}
&struct { Name string; Score int64 }{Name:"Silence", Score:10000}
```








动态创建函数并调用

```go
package main

import (
	"fmt"
	"reflect"
)

// 定义User结构体，并为每个属性定义标签

type User struct {
	id     int     `json:"id"`
	name   string  `json:"name"`
	Tel    string  `json:"addr"`
	Height float32 `json:"height"`
	Desc   *string `json:"desc"`
	Weight *int    `json:"weight"`
}

func NewUser(id int, name string, tel string, height float32, desc string, weight int) *User {
	return &User{id, name, tel, height, &desc, &weight}
}

// 定义String方法
func (user User) String() string {
	return fmt.Sprintf("User{id: %d, name: %s, tel: %s, height: %e, desc: %s, weight: %d}", user.id, user.name, user.Tel, user.Height, *user.Desc, *user.Weight)
}

func (user User) GetId() int {
	return user.id
}

func (user User) SetId(id int) {
	user.id = id
}

func (user *User) GetName() string {
	return user.name
}

func (user *User) SetName(name string) {
	user.name = name
}

// 定义接口Closer
type Closer interface {
	Close() error
}

type Address struct {
	ip   string "json:ip"
	port int    "json:port"
}

func (address Address) GetIp() string {
	return address.ip

}

func (address Address) GetPort() int {
	return address.port
}

type Connection struct {
	Address
	status int
}

func NewConnection(ip string, port int) *Connection {
	return &Connection{Address: Address{ip, port}}
}

func (conn *Connection) Send(msg string) error {
	fmt.Printf("发送消息给[%s:%d]:%s", conn.ip, conn.port, msg)
	return nil
}

func (conn *Connection) Close() error {
	fmt.Printf("关闭连接[%s:%d]", conn.ip, conn.port)
	return nil
}


// 定义func用于求和
sum := func(parameters []reflect.Value) []reflect.Value {
	// 若函数为不可变参数函数，则函数所有参数都包含在parameters中
	// 若函数为可变参数函数，则parameters最后一个元素为切片类型，包含调用函数传递的为对应的所有参数
	var total int64
	// 将前N-1个元素进行相加
	for _, parameter := range parameters[:len(parameters)-1] {
		total += parameter.Int()
	}

	// 对最后一个参数进行特殊处理
	last := parameter[len(parameters)-1]
	switch last.Kind() {
	case reflect.Int:
		total += last.Int()
	case reflect.Slice:
		// 若为切片，则为可变参数函数调用，遍历切片中元素进行求和
		for i := 0; i < last.Len(); i++ {
			total += last.Index(i).Int()
		}
	}

	// 返回包含总和初始化的value切片，切片元素数量对应函数返回值数量
	return []reflect.Value{reflect.ValueOf(total)}

}
// 定义函数变量add2
var add2 func(int, int) int64

ref := reflect.ValueOf(&add2).Elem() // 获取add2的引用
fv  := refalect.MakeFunc(ref.Type(), sum) // 创建add2的函数值
ref.Set(fv) // 通过add2的引导设置函数值

// 调用add2 函数并打印结果
fmt.Println(add2(1, 2))

// 定义匿名函数用于将函数类型变量进行初始化
// 参数为函数类型变量的指针
makeFunc := func(fn interface{}) {
	ref := reflect.ValueOf(fn).Elem()   // fn的引用
	fv  := reflect.MakeFunc(ref.Type(), sum) // 根据fn的类型定义函数值
	ref.Set(fv)
}

// 定义函数变量
var add3 func(int, int, int) int64
var add4 func(int, int, int, int) int64

// 通过引用设置函数变量值
makeFunc(&add3)
makeFunc(&add4)

// 调用函数
fmt.Println(add3(1, 2, 3))
fmt.Println(add4(1, 2, 3, 4))

// 定义可变参数类型函数
var add func(int, int, ...int) int64

// 通过引用设置函数变量值
makeFunc(&add)

// 调用函数
fmt.Println(add(3, 4))
fmt.Println(add(3, 4, 5))


func main() {
	// 定义结构体属性
	fields := []reflect.StructField{
		{
			Name: "Name",
			Type: reflect.TypeOf(""),
		},
		{
			Name: "Score",
			Type: reflect.TypeOf(int64(0)),
		},
	}
	// 定义结构体类型
	UserType := reflect.StructOf(fields)

	// 创建结构体对象
	user := reflect.New(UserType).Elem()

	// 设置属性值
	user.Field(0).Set(reflect.ValueOf("Silence"))
	user.Field(1).Set(reflect.ValueOf(int64(10000)))

	// 打印对象和对象指针
	fmt.Printf("%#v\n", user.Interface())
	fmt.Printf("%#v\n", user.Addr().Interface())
}





-------------------
测试失败
```




## 应用 

反射常用在动态处理所有数据类型(json/xml 序列化和反序列化, orm 等)或调用所有函数(路由函数调用)的场景。

### encoding/json

1. 介绍

在内存数据进行持久化存储或网络交换时常可采用 json 格式字符串，go 语言提供 json包进行 json 序列化和反序列化。

对于 Go 提供的基本类型和符合类型可以直接使用 json 包进行序列化和反序列化操作，针对结构体可通过标签声明属性和 json 字符串的转换关系，标签名为 json，常用格式为：

- json: 
  - 默认形式，可省略，json键名使用属性名，类型通过属性对应的类型
- json:"key"
  - json 键名使用指定名称 key
- json:"-"
  - 忽略该属性
- json:"key,type"
  - json 键名使用指定名称 key，类型使用指定类型 type
- json:"key,omitempty"
  - json 键名使用指定名称 key，当值为零值时省略该属性
- json:"key,type,omitempty"
  - json 键名使用指定名称 key，类型使用指定类型 type，当值为零值时省略该属性
- json:",type"
  - 类型使用指定类型 type
- json:",omitempty"
  - 值为零值时省略该属性

```go
type User struct {
	ID           int       `json: `                             // json键使用属性名ID
	Name         string    `json: "name"`                       // json键重命名为name
	Password     string    `json: "-"`                          // json 序列化忽略该字段
	gender       bool      `json: "gender"`                     // 非公开字段不进行序列化
	Online       bool      `json: "is_online,string,omitempty"` // json键重命名为is_online,并设置值类型为string
	RegisterTime time.time `json: "register_time"`              // json键重命名为register_time
	Status       int       `json: ",omitempty"`                 // json键使用属性名Status, 当为0值时不进行序列化

}

```

2. 数据初始化 

```go
// 定义用户类型容器(切片和映射)
users_slice := make([]User, 10, 20)
users_map  := make(map[int]User)

// 初始化切片和映射

for i := 0; i < 10; i++ {
	u := User{
		ID: 1,
		Name: fmt.Sprintf("user.%d", i), 
		Password: fmt.Sprintf("password.%d", i)
		gender: i%2 == 0,
		Online: i%3 == 0,
		RegisterTime: time.Now(),
		Status: i % 5,
	}
	users_slice[i] = u
	users_map[i] = u
}


// 定义json格式字符串
jsonStr := `
[
	{
		"ID": 0, 
		"Name": "slience", 
		"Password": "test@123", 
		"is_online": "true", 
		"register_time": "2019-04-19T14:34:41.5234523+08:00"
	},
	{
		"ID": 1, 
		"Name": "kk", 
		"is_online": "false", 
		"register_time": "2019-04-19T14:34:41.5200023+08:00"
		"status": 1
	}
]`
```

3. 常用函数
1) Marshal: 用于将 go 语言的数据序列化为 json 字符串
2) Indent: 将 json 字符串进行格式化

```go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// 定义User结构体，并使用标签声明json属性

type User struct {
	ID           int       `json: `                             // json键使用属性名ID
	Name         string    `json: "name"`                       // json键重命名为name
	Password     string    `json: "-"`                          // json 序列化忽略该字段
	gender       bool      `json: "gender"`                     // 非公开字段不进行序列化
	Online       bool      `json: "is_online,string,omitempty"` // json键重命名为is_online,并设置值类型为string
	RegisterTime time.Time `json: "register_time"`              // json键重命名为register_time
	Status       int       `json: ",omitempty"`                 // json键使用属性名Status, 当为0值时不进行序列化

}

func main() {

	// 定义用户类型容器(切片和映射)
	users_slice := make([]User, 10, 20)
	users_map := make(map[int]User)

	// 初始化切片和映射

	for i := 0; i < 10; i++ {
		u := User{
			ID:           1,
			Name:         fmt.Sprintf("user.%d", i),
			Password:     fmt.Sprintf("password.%d", i),
			gender:       i%2 == 0,
			Online:       i%3 == 0,
			RegisterTime: time.Now(),
			Status:       i % 5,
		}
		users_slice[i] = u
		users_map[i] = u
	}

	// 对切片进行json序列化
	b, err := json.Marshal(users_slice)
	if err == nil {
		fmt.Println(string(b))
		// 将json 字符串格式化并输出到bytes.Buffer变量中
		var buffer bytes.Buffer
		json.Indent(&buffer, b, "", "\t")
		buffer.WriteTo(os.Stdout)
	}
}




PS E:\go-phase-two\go-course> go run E:\go-phase-two\go-course\1test.go
[{"ID":1,"Name":"user.0","Password":"password.0","Online":true,"RegisterTime":"2020-10-15T19:16:42.6813396+08:00","Status":0},{"ID":1,"Name":"user.1","Password":"password.1","Online":false,"RegisterTime":"2020-10-15T19:16:42.6813396+08:00","Status":1},{"ID":1,"Name":"user.2","Password":"password.2","Online":false,"RegisterTime":"2020-10-15T19:16:42.6813396+08:00","Status":2},{"ID":1,"Name":"user.3","Password":"password.3","Online":true,"RegisterTime":"2020-10-15T19:16:42.6813396+08:00","Status":3},{"ID":1,"Name":"user.4","Password":"password.4","Online":false,"RegisterTime":"2020-10-15T19:16:42.6813396+08:00","Status":4},{"ID":1,"Name":"user.5","Password":"password.5","Online":false,"RegisterTime":"2020-10-15T19:16:42.6813396+08:00","Status":0},{"ID":1,"Name":"user.6","Password":"password.6","Online":true,"RegisterTime":"2020-10-15T19:16:42.6813396+08:00","Status":1},{"ID":1,"Name":"user.7","Password":"password.7","Online":false,"RegisterTime":"2020-10-15T19:16:42.6813396+08:00","Status":2},{"ID":1,"Name":"user.8","Password":"password.8","Online":false,"RegisterTime":"2020-10-15T19:16:42.6813396+08:00","Status":3},{"ID":1,"Name":"user.9","Password":"password.9","Online":true,"RegisterTime":"2020-10-15T19:16:42.6813396+08:00","Status":4}]
[
        {
                "ID": 1,
                "Name": "user.0",
                "Password": "password.0",
                "Online": true,
                "RegisterTime": "2020-10-15T19:16:42.6813396+08:00",
                "Status": 0
        },
        {
                "ID": 1,
                "Name": "user.1",
                "Password": "password.1",
                "Online": false,
                "RegisterTime": "2020-10-15T19:16:42.6813396+08:00",
                "Status": 1
        },
        {
                "ID": 1,
                "Name": "user.2",
                "Password": "password.2",
                "Online": false,
                "RegisterTime": "2020-10-15T19:16:42.6813396+08:00",
                "Status": 2
        },
        {
                "ID": 1,
                "Name": "user.3",
                "Password": "password.3",
                "Online": true,
                "RegisterTime": "2020-10-15T19:16:42.6813396+08:00",
                "Status": 3
        },
        {
                "ID": 1,
                "Name": "user.4",
                "Password": "password.4",
                "Online": false,
                "RegisterTime": "2020-10-15T19:16:42.6813396+08:00",
                "Status": 4
        },
        {
                "ID": 1,
                "Name": "user.5",
                "Password": "password.5",
                "Online": false,
                "RegisterTime": "2020-10-15T19:16:42.6813396+08:00",
                "Status": 0
        },
        {
                "ID": 1,
                "Name": "user.6",
                "Password": "password.6",
                "Online": true,
                "RegisterTime": "2020-10-15T19:16:42.6813396+08:00",
                "Status": 1
        },
        {
                "ID": 1,
                "Name": "user.7",
                "Password": "password.7",
                "Online": false,
                "RegisterTime": "2020-10-15T19:16:42.6813396+08:00",
                "Status": 2
        },
        {
                "ID": 1,
                "Name": "user.8",
                "Password": "password.8",
                "Online": false,
                "RegisterTime": "2020-10-15T19:16:42.6813396+08:00",
                "Status": 3
        },
        {
                "ID": 1,
                "Name": "user.9",
                "Password": "password.9",
                "Online": true,
                "RegisterTime": "2020-10-15T19:16:42.6813396+08:00",
                "Status": 4
        }
]
```


3) UnMarshal: 用于json字符串反序列化为go语言的数据

```go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// 定义User结构体，并使用标签声明json属性

type User struct {
	ID           int       `json: `                             // json键使用属性名ID
	Name         string    `json: "name"`                       // json键重命名为name
	Password     string    `json: "-"`                          // json 序列化忽略该字段
	gender       bool      `json: "gender"`                     // 非公开字段不进行序列化
	Online       bool      `json: "is_online,string,omitempty"` // json键重命名为is_online,并设置值类型为string
	RegisterTime time.Time `json: "register_time"`              // json键重命名为register_time
	Status       int       `json: ",omitempty"`                 // json键使用属性名Status, 当为0值时不进行序列化

}

func main() {

	// 定义用户类型容器(切片和映射)
	users_slice := make([]User, 10, 20)
	users_map := make(map[int]User)

	// 初始化切片和映射

	for i := 0; i < 10; i++ {
		u := User{
			ID:           1,
			Name:         fmt.Sprintf("user.%d", i),
			Password:     fmt.Sprintf("password.%d", i),
			gender:       i%2 == 0,
			Online:       i%3 == 0,
			RegisterTime: time.Now(),
			Status:       i % 5,
		}
		users_slice[i] = u
		users_map[i] = u
	}

	// 定义json格式字符串
	jsonStr := `
	[
	{
		"ID": 0,
		"Name": "slience",
		"Password": "test@123",
		"is_online": "true",
		"register_time": "2019-04-19T14:34:41.5234523+08:00"
	},
	{
		"ID": 1,
		"Name": "kk",
		"is_online": "false",
		"register_time": "2019-04-19T14:34:41.5200023+08:00"
		"status": 1
	}
	]
	`

	// 对切片进行json序列化
	b, err := json.Marshal(users_slice)
	if err == nil {
		fmt.Println(string(b))
		// 将json 字符串格式化并输出到bytes.Buffer变量中
		var buffer bytes.Buffer
		json.Indent(&buffer, b, "", "\t")
		buffer.WriteTo(os.Stdout)
	}

	// 将json 字符串进行反序列化到切片对象中
	var users []User
	err = json.Unmarshal([]byte(jsonStr), &users)
	if err == nil {
		fmt.Println(users)
	}
}





PS E:\go-phase-two\go-course> go run E:\go-phase-two\go-course\1test.go
[{"ID":1,"Name":"user.0","Password":"password.0","Online":true,"RegisterTime":"2020-10-15T19:20:28.6491651+08:00","Status":0},{"ID":1,"Name":"user.1","Password":"password.1","Online":false,"RegisterTime":"2020-10-15T19:20:28.6491651+08:00","Status":1},{"ID":1,"Name":"user.2","Password":"password.2","Online":false,"RegisterTime":"2020-10-15T19:20:28.6491651+08:00","Status":2},{"ID":1,"Name":"user.3","Password":"password.3","Online":true,"RegisterTime":"2020-10-15T19:20:28.6491651+08:00","Status":3},{"ID":1,"Name":"user.4","Password":"password.4","Online":false,"RegisterTime":"2020-10-15T19:20:28.6491651+08:00","Status":4},{"ID":1,"Name":"user.5","Password":"password.5","Online":false,"RegisterTime":"2020-10-15T19:20:28.6491651+08:00","Status":0},{"ID":1,"Name":"user.6","Password":"password.6","Online":true,"RegisterTime":"2020-10-15T19:20:28.6491651+08:00","Status":1},{"ID":1,"Name":"user.7","Password":"password.7","Online":false,"RegisterTime":"2020-10-15T19:20:28.6491651+08:00","Status":2},{"ID":1,"Name":"user.8","Password":"password.8","Online":false,"RegisterTime":"2020-10-15T19:20:28.6491651+08:00","Status":3},{"ID":1,"Name":"user.9","Password":"password.9","Online":true,"RegisterTime":"2020-10-15T19:20:28.6491651+08:00","Status":4}]
[
        {
                "ID": 1,
                "Name": "user.0",
                "Password": "password.0",
                "Online": true,
                "RegisterTime": "2020-10-15T19:20:28.6491651+08:00",
                "Status": 0
        },
        {
                "ID": 1,
                "Name": "user.1",
                "Password": "password.1",
                "Online": false,
                "RegisterTime": "2020-10-15T19:20:28.6491651+08:00",
                "Status": 1
        },
        {
                "ID": 1,
                "Name": "user.2",
                "Password": "password.2",
                "Online": false,
                "RegisterTime": "2020-10-15T19:20:28.6491651+08:00",
                "Status": 2
        },
        {
                "ID": 1,
                "Name": "user.3",
                "Password": "password.3",
                "Online": true,
                "RegisterTime": "2020-10-15T19:20:28.6491651+08:00",
                "Status": 3
        },
        {
                "ID": 1,
                "Name": "user.4",
                "Password": "password.4",
                "Online": false,
                "RegisterTime": "2020-10-15T19:20:28.6491651+08:00",
                "Status": 4
        },
        {
                "ID": 1,
                "Name": "user.5",
                "Password": "password.5",
                "Online": false,
                "RegisterTime": "2020-10-15T19:20:28.6491651+08:00",
                "Status": 0
        },
        {
                "ID": 1,
                "Name": "user.6",
                "Password": "password.6",
                "Online": true,
                "RegisterTime": "2020-10-15T19:20:28.6491651+08:00",
                "Status": 1
        },
        {
                "ID": 1,
                "Name": "user.7",
                "Password": "password.7",
                "Online": false,
                "RegisterTime": "2020-10-15T19:20:28.6491651+08:00",
                "Status": 2
        },
        {
                "ID": 1,
                "Name": "user.8",
                "Password": "password.8",
                "Online": false,
                "RegisterTime": "2020-10-15T19:20:28.6491651+08:00",
                "Status": 3
        },
        {
                "ID": 1,
                "Name": "user.9",
                "Password": "password.9",
                "Online": true,
                "RegisterTime": "2020-10-15T19:20:28.6491651+08:00",
                "Status": 4
        }
]
```


4) MarshalIndent: 用于将go 语言的数据序列化为为格式化的json 字符串 

```go
	// 对映射进行json序列化(并进行格式化)
	b, err = json.MarshalIndent(users_map, "", "\t")
	if err == nil {
		fmt.Println(string(b))
	}
```

5) Valid: 验证是否为正确json 字符串 

```go
	// 验证是否为正确的json字符串
	fmt.Println(json.Valid([]byte("1")))
	fmt.Println(json.Valid([]byte("true")))
	fmt.Println(json.Valid([]byte("1.1")))
	fmt.Println(json.Valid([]byte(`{}`)))
	fmt.Println(json.Valid([]byte(`[]`)))
	fmt.Println(json.Valid([]byte(`[{}, {}]`)))
	fmt.Println(json.Valid([]byte(`{"name" : "kk"}`)))
	fmt.Println(json.Valid([]byte(`{"name": kk}`)))     // json格式字符串需要使用"双引号包含
	fmt.Println(json.Valid([]byte(`{"name" : "kk",}`))) // json格式最后一个字段后不许有,
```


汇总测试

```go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// 定义User结构体，并使用标签声明json属性

type User struct {
	ID           int       `json: `                             // json键使用属性名ID
	Name         string    `json: "name"`                       // json键重命名为name
	Password     string    `json: "-"`                          // json 序列化忽略该字段
	gender       bool      `json: "gender"`                     // 非公开字段不进行序列化
	Online       bool      `json: "is_online,string,omitempty"` // json键重命名为is_online,并设置值类型为string
	RegisterTime time.Time `json: "register_time"`              // json键重命名为register_time
	Status       int       `json: ",omitempty"`                 // json键使用属性名Status, 当为0值时不进行序列化

}

func main() {

	// 定义用户类型容器(切片和映射)
	users_slice := make([]User, 10, 20)
	users_map := make(map[int]User)

	// 初始化切片和映射

	for i := 0; i < 10; i++ {
		u := User{
			ID:           1,
			Name:         fmt.Sprintf("user.%d", i),
			Password:     fmt.Sprintf("password.%d", i),
			gender:       i%2 == 0,
			Online:       i%3 == 0,
			RegisterTime: time.Now(),
			Status:       i % 5,
		}
		users_slice[i] = u
		users_map[i] = u
	}

	// 定义json格式字符串
	jsonStr := `
	[
	{
		"ID": 0,
		"Name": "slience",
		"Password": "test@123",
		"is_online": "true",
		"register_time": "2019-04-19T14:34:41.5234523+08:00"
	},
	{
		"ID": 1,
		"Name": "kk",
		"is_online": "false",
		"register_time": "2019-04-19T14:34:41.5200023+08:00"
		"status": 1
	}
	]
	`

	// 对切片进行json序列化
	b, err := json.Marshal(users_slice)
	if err == nil {
		fmt.Println(string(b))
		// 将json 字符串格式化并输出到bytes.Buffer变量中
		var buffer bytes.Buffer
		json.Indent(&buffer, b, "", "\t")
		buffer.WriteTo(os.Stdout)
	}

	// 将json 字符串进行反序列化到切片对象中
	var users []User
	err = json.Unmarshal([]byte(jsonStr), &users)
	if err == nil {
		fmt.Println(users)
	}

	// 对映射进行json序列化(并进行格式化)
	b, err = json.MarshalIndent(users_map, "", "\t")
	if err == nil {
		fmt.Println(string(b))
	}

	// 验证是否为正确的json字符串
	fmt.Println(json.Valid([]byte("1")))
	fmt.Println(json.Valid([]byte("true")))
	fmt.Println(json.Valid([]byte("1.1")))
	fmt.Println(json.Valid([]byte(`{}`)))
	fmt.Println(json.Valid([]byte(`[]`)))
	fmt.Println(json.Valid([]byte(`[{}, {}]`)))
	fmt.Println(json.Valid([]byte(`{"name" : "kk"}`)))
	fmt.Println(json.Valid([]byte(`{"name": kk}`)))     // json格式字符串需要使用"双引号包含
	fmt.Println(json.Valid([]byte(`{"name" : "kk",}`))) // json格式最后一个字段后不许有,
}



PS E:\go-phase-two\go-course> go run E:\go-phase-two\go-course\1test.go
[{"ID":1,"Name":"user.0","Password":"password.0","Online":true,"RegisterTime":"2020-10-15T19:34:58.9365118+08:00","Status":0},{"ID":1,"Name":"user.1","Password":"password.1","Online":false,"RegisterTime":"2020-10-15T19:34:58.9365118+08:00","Status":1},{"ID":1,"Name":"user.2","Password":"password.2","Online":false,"RegisterTime":"2020-10-15T19:34:58.9365118+08:00","Status":2},{"ID":1,"Name":"user.3","Password":"password.3","Online":true,"RegisterTime":"2020-10-15T19:34:58.9365118+08:00","Status":3},{"ID":1,"Name":"user.4","Password":"password.4","Online":false,"RegisterTime":"2020-10-15T19:34:58.9365118+08:00","Status":4},{"ID":1,"Name":"user.5","Password":"password.5","Online":false,"RegisterTime":"2020-10-15T19:34:58.9365118+08:00","Status":0},{"ID":1,"Name":"user.6","Password":"password.6","Online":true,"RegisterTime":"2020-10-15T19:34:58.9365118+08:00","Status":1},{"ID":1,"Name":"user.7","Password":"password.7","Online":false,"RegisterTime":"2020-10-15T19:34:58.9365118+08:00","Status":2},{"ID":1,"Name":"user.8","Password":"password.8","Online":false,"RegisterTime":"2020-10-15T19:34:58.9365118+08:00","Status":3},{"ID":1,"Name":"user.9","Password":"password.9","Online":true,"RegisterTime":"2020-10-15T19:34:58.9365118+08:00","Status":4}]
[
        {
                "ID": 1,
                "Name": "user.0",
                "Password": "password.0",
                "Online": true,
                "RegisterTime": "2020-10-15T19:34:58.9365118+08:00",
                "Status": 0
        },
        {
                "ID": 1,
                "Name": "user.1",
                "Password": "password.1",
                "Online": false,
                "RegisterTime": "2020-10-15T19:34:58.9365118+08:00",
                "Status": 1
        },
        {
                "ID": 1,
                "Name": "user.2",
                "Password": "password.2",
                "Online": false,
                "RegisterTime": "2020-10-15T19:34:58.9365118+08:00",
                "Status": 2
        },
        {
                "ID": 1,
                "Name": "user.3",
                "Password": "password.3",
                "Online": true,
                "RegisterTime": "2020-10-15T19:34:58.9365118+08:00",
                "Status": 3
        },
        {
                "ID": 1,
                "Name": "user.4",
                "Password": "password.4",
                "Online": false,
                "RegisterTime": "2020-10-15T19:34:58.9365118+08:00",
                "Status": 4
        },
        {
                "ID": 1,
                "Name": "user.5",
                "Password": "password.5",
                "Online": false,
                "RegisterTime": "2020-10-15T19:34:58.9365118+08:00",
                "Status": 0
        },
        {
                "ID": 1,
                "Name": "user.6",
                "Password": "password.6",
                "Online": true,
                "RegisterTime": "2020-10-15T19:34:58.9365118+08:00",
                "Status": 1
        },
        {
                "ID": 1,
                "Name": "user.7",
                "Password": "password.7",
                "Online": false,
                "RegisterTime": "2020-10-15T19:34:58.9365118+08:00",
                "Status": 2
        },
        {
                "ID": 1,
                "Name": "user.8",
                "Password": "password.8",
                "Online": false,
                "RegisterTime": "2020-10-15T19:34:58.9365118+08:00",
                "Status": 3
        },
        {
                "ID": 1,
                "Name": "user.9",
                "Password": "password.9",
                "Online": true,
                "RegisterTime": "2020-10-15T19:34:58.9365118+08:00",
                "Status": 4
        }
]{
        "0": {
                "ID": 1,
                "Name": "user.0",
                "Password": "password.0",
                "Online": true,
                "RegisterTime": "2020-10-15T19:34:58.9365118+08:00",
                "Status": 0
        },
        "1": {
                "ID": 1,
                "Name": "user.1",
                "Password": "password.1",
                "Online": false,
                "RegisterTime": "2020-10-15T19:34:58.9365118+08:00",
                "Status": 1
        },
        "2": {
                "ID": 1,
                "Name": "user.2",
                "Password": "password.2",
                "Online": false,
                "RegisterTime": "2020-10-15T19:34:58.9365118+08:00",
                "Status": 2
        },
        "3": {
                "ID": 1,
                "Name": "user.3",
                "Password": "password.3",
                "Online": true,
                "RegisterTime": "2020-10-15T19:34:58.9365118+08:00",
                "Status": 3
        },
        "4": {
                "ID": 1,
                "Name": "user.4",
                "Password": "password.4",
                "Online": false,
                "RegisterTime": "2020-10-15T19:34:58.9365118+08:00",
                "Status": 4
        },
        "5": {
                "ID": 1,
                "Name": "user.5",
                "Password": "password.5",
                "Online": false,
                "RegisterTime": "2020-10-15T19:34:58.9365118+08:00",
                "Status": 0
        },
        "6": {
                "ID": 1,
                "Name": "user.6",
                "Password": "password.6",
                "Online": true,
                "RegisterTime": "2020-10-15T19:34:58.9365118+08:00",
                "Status": 1
        },
        "7": {
                "ID": 1,
                "Name": "user.7",
                "Password": "password.7",
                "Online": false,
                "RegisterTime": "2020-10-15T19:34:58.9365118+08:00",
                "Status": 2
        },
        "8": {
                "ID": 1,
                "Name": "user.8",
                "Password": "password.8",
                "Online": false,
                "RegisterTime": "2020-10-15T19:34:58.9365118+08:00",
                "Status": 3
        },
        "9": {
                "ID": 1,
                "Name": "user.9",
                "Password": "password.9",
                "Online": true,
                "RegisterTime": "2020-10-15T19:34:58.9365118+08:00",
                "Status": 4
        }
}
true
true
true
true
true
true
true
false
false
```

4. Encoder 与 Decoder 

Encoder 和 Decoder 是构建在流之上进行Json 序列化和反序列化，构造函数分别为NewEncoder和NewDecoder， 参数要求分别实现接口io.Writer 和 io.Reader的对象

1) Encoder.Encode 序列化


```go

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// 定义User结构体，并使用标签声明json属性

type User struct {
	ID           int       `json: `                             // json键使用属性名ID
	Name         string    `json: "name"`                       // json键重命名为name
	Password     string    `json: "-"`                          // json 序列化忽略该字段
	gender       bool      `json: "gender"`                     // 非公开字段不进行序列化
	Online       bool      `json: "is_online,string,omitempty"` // json键重命名为is_online,并设置值类型为string
	RegisterTime time.Time `json: "register_time"`              // json键重命名为register_time
	Status       int       `json: ",omitempty"`                 // json键使用属性名Status, 当为0值时不进行序列化

}

func main() {

	// 定义用户类型容器(切片和映射)
	users_slice := make([]User, 10, 20)
	users_map := make(map[int]User)

	// 初始化切片和映射

	for i := 0; i < 10; i++ {
		u := User{
			ID:           1,
			Name:         fmt.Sprintf("user.%d", i),
			Password:     fmt.Sprintf("password.%d", i),
			gender:       i%2 == 0,
			Online:       i%3 == 0,
			RegisterTime: time.Now(),
			Status:       i % 5,
		}
		users_slice[i] = u
		users_map[i] = u
	}

	// 定义json格式字符串
	jsonStr := `
	[
	{
		"ID": 0,
		"Name": "slience",
		"Password": "test@123",
		"is_online": "true",
		"register_time": "2019-04-19T14:34:41.5234523+08:00"
	},
	{
		"ID": 1,
		"Name": "kk",
		"is_online": "false",
		"register_time": "2019-04-19T14:34:41.5200023+08:00"
		"status": 1
	}
	]
	`

	// 对切片进行json序列化
	b, err := json.Marshal(users_slice)
	if err == nil {
		fmt.Println(string(b))
		// 将json 字符串格式化并输出到bytes.Buffer变量中
		var buffer bytes.Buffer
		json.Indent(&buffer, b, "", "\t")
		buffer.WriteTo(os.Stdout)
	}

	// 将json 字符串进行反序列化到切片对象中
	var users []User
	err = json.Unmarshal([]byte(jsonStr), &users)
	if err == nil {
		fmt.Println(users)
	}

	// 对映射进行json序列化(并进行格式化)
	b, err = json.MarshalIndent(users_map, "", "\t")
	if err == nil {
		fmt.Println(string(b))
	}

	// 验证是否为正确的json字符串
	fmt.Println(json.Valid([]byte("1")))
	fmt.Println(json.Valid([]byte("true")))
	fmt.Println(json.Valid([]byte("1.1")))
	fmt.Println(json.Valid([]byte(`{}`)))
	fmt.Println(json.Valid([]byte(`[]`)))
	fmt.Println(json.Valid([]byte(`[{}, {}]`)))
	fmt.Println(json.Valid([]byte(`{"name" : "kk"}`)))
	fmt.Println(json.Valid([]byte(`{"name": kk}`)))     // json格式字符串需要使用"双引号包含
	fmt.Println(json.Valid([]byte(`{"name" : "kk",}`))) // json格式最后一个字段后不许有,

	// 使用Encoder将json 结果输出到标准输出
	encoder := json.NewDecoder(os.Stdout)
	encoder.SetIndent("", "\t")
	err = encoder.Encode(users_slice)

	if err != nil {
		fmt.Println(err)
	}

	// 使用Encoder将序列化结构输出到bytes.Buffer/strings.Builder
	var out byte.Buffer
	// var out strings.Buffer

	encoder = json.NewDecoder(&out)
	err = encoder.Encode(users_map)
	if err == nil {
		fmt.Println(out.String())
	} else {
		fmt.Println(err)
	}
}


测试失败

```



2) Decoder.Decode 反序列化

```go
	// 使用Decoder 从bytes.Buffer或strings.Reader中读取结果并反序列化

	// 将json 字符串写入到bytes.Buffer中供Decoder读取
	in := bytes.NewBufferString(jsonStr)
	// 将json 字符串写入到strings.Reader中供Decoder读取
	// in := strings.NewReader(jsonStr)

	var users2 []User
	decoder := json.NewDecoder(in)
	err = decoder.Decode(&users2)
	if err == nil {
		fmt.Println(users2)
	} else {
		fmt.Println(err)
	}

```


## encodeing/xml 

1) 介绍
go 语言提供 xml 包进行 xml 序列化和反序列化
在进行 xml 序列化和反序列化时，必须定义根结构体对应 xml 的根节点，并可通过结构体中 xml.Name 类型且名称为 XMLName 属性来定义根节点，可通过各属性的标签声明属性和 xml 节点的转换关系，标签名为 xml，常用格式为：
- xml:
  - 默认形式，可省略，xml 节点使用结构体属性名
- xml:"name"
  - xml 节点名使用指定的名称 name
-  xml:"-"
   - 忽略该属性
-  xml:"name,attr"
   - 表示该属性为节点的属性值，属性名使用指定名称 name
-  xml:",attr"
   - 表示为节点的属性值，属性名为结构体属性名
-  xml:"name,omitempty"
   - xml 节点名使用指定的名称 name，当值为零值时省略该属性
-  xml:",omitempty"
   - 值为零值时省略该属性
-  xml:", chardata"
   - 表示为文本
-  xml:",cdata"
   - 表示使用 CDATA 包含
-  xml:",comment"
   - 表示为注释内容
-  xml:", innerxml"
   - 表示为数据不进行实体编码
-  xml:"A>B>C"
   - 定义 XML 节点结构


```go


package main

import (
	"encoding/xml"
	"time"
)

type CDATA struct {
	Content string `xml: ",cdata"` // 定义数据使用CDATA包含
}

type Comment struct {
	Content string `xml:",comment"` // 定义数据为注释
}

type Address struct {
	City   string `xml:"address>city"`
	Street string `xml:"address>street"`
	No     string `xml:"address>no"`
}

// 定义User结构体，并使用标签声明json属性
type User struct {
	ID           int       `xml:"id,attr"`             // xml节点属性，且属性名为id
	Name         string    `xml:"name"`                // xml节点键重命名为name
	Password     string    `xml:"-"`                   // xml序列化忽略该字段
	gender       bool      `xml:"gender"`              // 非公开字段不进行序列化
	Online       bool      `xml:"is_online,omitempty"` // xml节点重命名为is_online
	RegisterTime time.Time `xml:"register_time"`       // xml节点重命名为register_time
	Status       int       `xml:",omitempty"`          // xml节点使用属性名Status，当为0值是不进行序列化
	Address
	Desc    CDATA   `xml:"desc"`
	Comment Comment `xml:"comment"`
	// Extend string `xml:",innerxml"` // 定义数据不使用xml转义
}

type Root struct {
	XMLName xml.Name `xml:"root"` // 定义根节点名称
	User    []User   `xml:"User"` //
}

func main() {

}

```