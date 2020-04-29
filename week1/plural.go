// go 提供complex64 和 complex128 两种复数类型，针对complex64 复数的实部和虚部均使用float32，针对complex128复数的实部和虚部均使用float64

package main

import "fmt"

func main() {
	var (
		c1 complex64 = 1 + 2i
		c2 complex64 = complex(3, 4)
	)
	fmt.Println(c1 + c2)
	fmt.Println(real(c1), imag(c1))
}
