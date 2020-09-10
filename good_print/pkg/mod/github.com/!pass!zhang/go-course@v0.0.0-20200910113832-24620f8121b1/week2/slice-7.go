package main

import "fmt"

func main() {

	x := []int{1}

	fmt.Println("-----TODO1-----")
	fmt.Printf("%v, %d, %d\n", x, len(x), cap(x))

	fmt.Println("-----TODO2-----")
	x = append(x, 3)
	fmt.Printf("%v, %d, %d\n", x, len(x), cap(x))

	fmt.Println("-----TODO3-----")
	x = append(x, 5)
	fmt.Printf("%v, %d, %d\n", x, len(x), cap(x))

	fmt.Println("-----TODO4-----")
	y := append(x, 7)
	fmt.Printf("%v, %d, %d\n", x, len(x), cap(x))
	fmt.Printf("%v, %d, %d\n", y, len(y), cap(y))

	fmt.Println("-----TODO5-----")
	z := append(x, 9)
	fmt.Printf("%v, %d, %d\n", x, len(x), cap(x))
	fmt.Printf("%v, %d, %d\n", z, len(z), cap(z))

	fmt.Println("-----TODO6-----")
	fmt.Println(x, y, z)
}
