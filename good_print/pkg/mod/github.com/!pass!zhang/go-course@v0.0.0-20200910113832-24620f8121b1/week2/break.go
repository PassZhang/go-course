package main

import "fmt"

func main() {
BREAK:
	for z := 0; z < 3; z++ {
		for j := 0; j < 3; j++ {
			for i := 0; i < 5; i++ {
				fmt.Printf("%#v\n", i)
				if i == 3 {
					break BREAK
				}
			}
		}
	}
}