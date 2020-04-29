// go不会自动对数据类型转换，因此左右操作数类型必须一致或者某个字面量，可通过类型名(数据)的语法将数据转换为对应类型。需要注意值截断和值溢出问题。

/*
使用fmt.Printf进行格式化参数输出，占位符
1. %b：二进制
2. %c：字符
3. %d：十进制
	1. %+d表示对正整数带+符号
	2. %nd表示最小占位n个宽度且右对齐
	3. %-nd表示最小占位n个宽度且左对齐
	4. %0nd表示最小占位n个宽度且右对齐，空字符使用0填充
4. %o： 八进制，%#o带0的前缀
5. %x、%X：十六进制，%#x(%#X)带0x*(0X)的前缀
6. %U: Unicode码点，%#U带字符的Unicode码点
7. %q：带单引号的字符
*/

package main

import "fmt"

func main() {
	var (
		n5 int  = 10
		n6 int  = -10
		c0 byte = 'a'
		u0 rune = '牛'
	)

	fmt.Printf("%d %b %o %x %X %#o %#x %#X\n", n5, n5, n5, n5, n5, n5, n5, n5)
	fmt.Printf("%d|%+d|%10d|%-10d|%010d|%+-10d|%+010d|\n", n5, n5, n5, n5, n5, n5, n5, n5, n5)
	fmt.Printf("%d|%+d|%10d|%-10d|%010d|%+-10d|%+010d|\n", n6, n6, n6, n6, n6, n6, n6, n6)
	fmt.Printf("%c %c %q %q\n", c0, u0, c0, u0)
	fmt.Printf("%U %#U\n", u0, u0)
}
