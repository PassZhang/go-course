package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func main() {
	ascii := "abc我爱中华人民共和国"
	fmt.Println([]byte(ascii))
	fmt.Println([]rune(ascii))

	fmt.Println(len(ascii))
	fmt.Println(len([]rune(ascii)))

	fmt.Println(utf8.RuneCountInString(ascii))

	fmt.Println(string([]byte(ascii)))
	fmt.Println(string([]rune(ascii)))

	// int, float, bool
	fmt.Println(strconv.Itoa('a'))
	ch, err := strconv.Atoi("97")
	fmt.Println(ch, err)

}
