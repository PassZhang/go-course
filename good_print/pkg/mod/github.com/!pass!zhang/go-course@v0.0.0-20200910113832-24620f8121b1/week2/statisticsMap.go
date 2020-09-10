package main

import "fmt"

func main()  {
	article := `
	I am happy to join with you today in what will go down in history as the greatest demonstration for freedom in the history of our nation. 
	`

	stats := make(map[rune]int)
	for _, ch := range article {
		if ch >= 'a' && ch <= 'z' {
			stats[ch]++
		}
		for ch, cnt := range stats {
			fmt.Printf("%c, %v\n", ch, cnt)
			
		}
	}
}