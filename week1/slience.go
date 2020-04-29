package main

import "fmt"

func main() {
	var name string = "kk"

	fmt.Println(name)

	name = "silence"

	fmt.Println(name)

	{
		name = "aaaaaAAAAAAAAAAAAAA"
	}

	fmt.Println(name)
}
