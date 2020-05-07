package main

import "fmt"

func main() {
CONTINUE:
	for j := 0; j < 5; j++ {
		fmt.Println(j, "\n-----------------------\n")
		for i := 0; i < 5; i++ {
			if i == 3 {
				continue CONTINUE
			}
			fmt.Println(i)
		}
	}
}
