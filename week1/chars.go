package main

import "fmt"

func main() {
	letters := "abcdefg"
	//0 1 2 3 ...len(letters) - 1
	for i := 0; i < len(letters); i++ { //len函数时检查变量的长度 并且显示出来，上面len(letters),就代表数字为7
		fmt.Printf("%c\n", letters[i])

	}

	letters = "我爱中华"
	for _, v := range letters {
		fmt.Printf("%q\n", v)
	}
}
