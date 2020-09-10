package main

import "fmt"

func main() {
	var a int = 21
	var c int

	c = a
	// 普通赋值运算符
	fmt.Printf("第一行c值 = %d\n", c)

	c += a
	// 相加后再赋值 C += a 等于 C = C + a
	fmt.Printf("第二行c值 = %d\n", c)

	c -= a
	// 相减后再赋值 C -= a 等于 c = c - a
	fmt.Printf("第三行c值 = %d\n", c)

	c *= a
	// 相乘后再赋值	C *= A 等于 C = C * A
	fmt.Printf("第四行c值 = %d\n", c)

	c /= a
	// 相除后再赋值	C /= A 等于 C = C / A
	fmt.Printf("第五行c值 = %d\n", c)

	c %= a
	// 求余后再赋值	C %= A  等于 C = C % A
	fmt.Printf("第五行c值 = %d\n", c)

	c = 200

	c <<= 2
	// 	左移后赋值	C <<= 2 等于 C = C << 2
	fmt.Printf("第6行c值 = %d\n", c)

	c >>= 2
	//  右移后赋值	C >>= 2 等于 C = C >> 2
	fmt.Printf("第7行c值 = %d\n", c)

	c &= 2
	// 按二进制位与后赋值	C &= 2 等于 C = C & 2
	fmt.Printf("第8行c值 = %d\n", c)

	c ^= 2
	// 按二进制位异或后赋值	C ^= 2 等于 C = C ^ 2
	fmt.Printf("第9行c值 ^= %d\n", c)

	c |= 2
	// 按二进制位或后赋值	C |= 2 等于 C = C | 2
	fmt.Printf("第10行c值 |= %d\n", c)

}
