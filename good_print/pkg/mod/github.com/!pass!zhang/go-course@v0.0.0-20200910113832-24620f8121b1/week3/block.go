package main

import "fmt"

func main() {
	name, desc := "kk", "i'm kk"

	func(name string) {
		fmt.Println(name, desc) // malukang, i'm kk
		name, desc = "烽火", "yanxin"
		fmt.Println(name, desc) //烽火， yanxin

	}("nalukang")
	fmt.Println(name, desc) // kk, yanxin
}
