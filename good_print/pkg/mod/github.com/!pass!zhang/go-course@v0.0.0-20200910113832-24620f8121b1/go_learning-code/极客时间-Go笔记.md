# 第一节 入门和介绍

## 软件开发的新挑战
1. 多核硬件架构
2. 超大规模分布式集群
3. Web模式导致的前所未有的开发规模和开发速度

go创始人：rob pike / ken thompson / robert griesemer 

## 程序的基本结构
```go
package main //包，表明代码所有的模块（包）
import "fmt" //引入代码依赖

// 功能实现
func main() {
    fmt.Println("Hello World!")
}
```

## 应用程序入口
1. 必须是main包： package mian 
2. 必须是main方法：func main()
3. 文件名不一定是main.go

## 退出返回值
与其他主要编程的语言的差异
1. go中main函数不支持任何返回值
2. 通过os.Exit来返回状态

```go
package main //包，表明代码所有的模块（包）
import (
    "fmt"
    "os"
)

// 功能实现
func main() {
    fmt.Println("Hello World!")
    os.Exit(-1)
}
```

## 获取命令行参数
与其他主要编程语言的差异
1. main函数不支持传入参数
    func main(arg []string)
2. 在程序中直接通过os.Args 获取命令行参数
```go
package main 

import (
    "fmt"
    "go"
)

func main(
    if len(os.Args) > 1 {
        fmt.Println("Hello World!", os.Args[1])
    }
)
```

# go 语言变量和常量与其他语言的差异
## 编写测试程序
1. 源码文件以_test结尾：xxx_test.go
2. 测试方法以Test： func TestXXX(t *testing.T) {...}

## 变量赋值
与其他主要编程语言的差异
1. 赋值可以进行自动类型推断
2. 在一个赋值语句可以对多个变量进行同时赋值。





 






