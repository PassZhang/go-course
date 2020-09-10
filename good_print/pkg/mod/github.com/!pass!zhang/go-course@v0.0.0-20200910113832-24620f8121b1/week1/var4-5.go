// Unicode字符编码
/*
1. Unicode（字符集）：为每一个[字符]分配一个唯一的ID（学名为码位/码点/Code Point）

2. utf-8(编码规则):将[码位]转换为字节序列的规则（解码/编码 可以理解为加密/解密的过程）- 以8位为一个编码单位的可变长编码。会将一个码位编码为1到4个字节

3. go 语言的所有源代码，都必须按照Unicode编码规范中的UTF-8编码格式进行编码。

4. utf-8编码： 一个英文字符等于一个自耳机，一个中文（含繁体）等于三个字节。中文标点占三个字节，英文标点占一个字节。

5. Unicode编码：一个英文等于两个字节，一个中文（含繁体）等于两个字节。中文变电占两个字节，英文标点占两个字节。

6. 一个rune类型的值在底层其实就是一个utf-8编码值。

*/

// package main 

// import "fmt"

// func main()  {

// 	// type byte = uint8 不带符号8位
// 	// type rune = int32 带符号32位

// 	str := "Hello! 一二三"
// 	strlist1 := []byte(str)
// 	strlist2 := []rune(str)

// 	// go中的字符串底层默认是[] byte形式存储的
// 	fmt.Println(len(str))

// 	// utf-8编码： 一个中文字符等于三个字节，显然一个uint8 无法存储三个字节
// 	fmt.Println(strlist1)
// 	fmt.Printf("byte:%s---%d---%b\n", string(strlist1), len(strlist1), strlist1)

// 	fmt.Println(strlist2)
// 	fmt.Printf("byte:%s---%d---%b\n", string(strlist2), len(strlist2), strlist2)

// }


// 循环取值

package main 

import "fmt"

func main()  {
	// type byte = uint8 不带符号8位
	// type rune = int32 带符号32位

	str := "Hello! 一二三"
	// 错误
	for i := 0; i < len(str); i++ {
		fmt.Print(string(str[i]))
	}
	fmt.Println()

	strlist2 := []rune(str)
	for i := 0; i < len(strlist2); i++ {
		fmt.Print(string(strlist2[i]))
	}
	fmt.Println()

	for _, v := range str {
		fmt.Print(string(v))
	}
}