package main

import (
	"bytes"
	"fmt"
)

// bytes 包实现了用于操作 []byte 的函数，类似于 strings 包中的函数
func main() {
	fmt.Println(bytes.Compare([]byte("abc"), []byte("xyz")))
}
