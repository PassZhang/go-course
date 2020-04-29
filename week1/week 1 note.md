# Golang第一周笔记
## 认识Go语言
golang github 地址：https://github.com/golang/go

golang 生态环境及开源软件介绍地址：https://github.com/golang/go/wiki/projects

## 环境安装
go 中文官网：https://golang.google.cn/

推荐下载版本：
![](https://raw.githubusercontent.com/PassZhang/passzhang.github.io/images-picgo/20200405140631.png)

现在使用的是Windows的版本，go软件下载好点击next安装就可以了，不需要多余的操作。安装好之后需要设置gomode和goproxy配置才可以，国内设置goproxy才能安装go对应的软件包

当前我安装的是go的1.13.5版本，打开cmd测试一下
![](https://raw.githubusercontent.com/PassZhang/passzhang.github.io/images-picgo/20200405141253.png)


Windows设置gomode和goproxy配置,在cmd中
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct


linux设置gomode和goproxy 
[www@localhost profile.d]$ pwd 
/etc/profile.d
cat >> go.sh <<EOF
if [[ "x" == "x${GOROOT}" ]]; then
    export GOROOT=/usr/local/go
    export GOPATH=${HOME}/go
    export PATH=${GOROOT}/bin:${GOPATH}/bin:${PATH}
fi
EOF



启动配置中必须要设定的变量有：
GOROOT
GOPATH
GO111MODULE
GOPROXY
----



关于如何学习一门编程：
1. 知识点分解
2. 刻意练习(练习不会的，不懂或者难以理解的知识点)
3. 反馈
   1. 主动反馈:看别人的代码
   2. 被动反馈：codereview

总结：多写（禁止复制和粘贴），多看，多问


变量定义：分为包变量(packageVar)，函数变量(funcVar)，块级别(blockVar)，子块变量(innerBlockVar)

从上到下的顺序分为：包变量(packageVar)<--函数变量(funcVar)<--块级别(blockVar)<--子块变量(innerBlockVar)

变量引用规则：下层引用上层变量可以，但是上层不能引用下层变量。例如函数变量可以引用包变量，包变量不能引用函数变量。因为变量的作用域不一样

代码执行顺序：从上到下


定义变量的规则：
标识符：程序中定义的名字、变量名、常量名字、函数名字、自定义类型、接口、包名字
对应规范：
    1. 必须满足： 组成只能由非空的Unicode编码字符串、数字、下划线组成。
    2. 必须满足： 必须以Unicode编码的字符串或下划线开头（不能以数字开头）
    3. 必须满足： 不能与go 的关键字冲突（package, func, var... 25个）

建议：
    1. Ascill编码（a-z，A-Z）、数字、下划线
    2. 变量使用驼峰式
       1. 多个英文字母 my_name myName（驼峰）
    3. 与go内置的标识符不要冲突（string...）

说明： 标识符区分大小写
    my = ""
    My = ""


go 语言提供一些预先定义的36个标识符用来表示内置的常量、类型、函数，在自定义标识符时应避免使用：
- 内置常量：true、false、nil、iota
- 内置类型：bool、type、rune、int、int8、int16、int32、uint、uint8、unit16、unit32、unit64、uintptr、float32、float64、complex64、complex128、string、error
- 内置函数：make、len、cap、new、append、copy、close、delete、complex、real、imag、panic、recover
- 空白标识符：_


go的关键字：用于特定的语法结构，go语言定义25个关键字：
声明：import、package


实体声明和定义：chan、const、func、interface、map、struct、type、var、

流程控制：break、case、continue、default、defer、else、fallthrough、for、go、goto、if、range、return、select、switch


字母量
字面量是值的表示方法，常用与对变量/常量进行初始化，主要分为：
1. 表示基础数据类型值的字面量，例如：0，1.1，true，3+4i，‘a’，“我爱中国”
2. 构造自定义的复合数据类型的类型字面量，例如



常用数据类型：
1. 布尔类型：用来表示真假，只有两个值 真：true 假：false
2. 整数类型
   1. int uint
   2. int8 int16 int32 int64
   3. uint8 uint16 uint32 uint64
   4. byte rune
   5. uintptr
   
3. 浮点类型
4. 复数类型
5. 字符串
6. 数组
7. 切片
8. 映射
9. 自定类型



if 使用技巧：
1. 不需要使用括号将条件包含起来
2. 大括号{} 必须存在，及时只有一行语句
3. 左括号必须在if或else的同一行
4. 在if之后，条件语句之前，可以添加变量初始化语句，使用; 进行分隔
5. 在有返回值的函数中，最终的return不能在条件语句中。
6. 
















## 程序结构
## 基本组成元素
## 基础知识点(1)


