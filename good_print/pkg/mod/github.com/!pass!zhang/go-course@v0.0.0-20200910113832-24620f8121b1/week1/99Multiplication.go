package main

import "fmt"

// func main () {
// 	for n := 1; n < 10; n++ {
// 		for m := 1; m <= n; m++ {
// 			fmt.Printf(" %d, %d, %d", n, m, n*m)
// 		}
// 		fmt.Println("")
// 	}
// }

func main() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%-2d ", i, j, i*j)
		}
		fmt.Println()
	}
}
