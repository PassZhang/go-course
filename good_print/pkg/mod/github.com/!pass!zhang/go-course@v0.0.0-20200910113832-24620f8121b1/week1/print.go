package main

import "fmt"

func main() {
	var name = "kk"

	fmt.Println("*")
	fmt.Println(name)
	fmt.Println("*")
	fmt.Print(name)
	fmt.Println("*")

	fmt.Printf("%T, %v, %#v", name, name, name)
}
