package main

import "fmt"

func main() {
	students := make([]string, 3, 5)
	fmt.Println(len(students), cap(students))

	fmt.Printf("%q\n", students)
}
