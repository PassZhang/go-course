package main

import "fmt"

/*

demo5(5,0)
demo5(4,5)
demo5(3,9)
demo5(2,12)
demo5(1,14)
demo5(0,15)
*/
func demo5(val int, total int) int {
	if val > 0 {
		return demo5(val-1, total+val)
	} else {
		return total
	}
}
func main() {
	fmt.Println(demo5(5, 0))
}
