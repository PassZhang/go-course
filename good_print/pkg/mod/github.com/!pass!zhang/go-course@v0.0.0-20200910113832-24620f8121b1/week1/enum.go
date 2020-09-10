package main

import "fmt"

func main() {
	const (
		Mon = iota
		Tuesd
		Wed
		Thur
		Fir
		Sat
		Sun
	)
	fmt.Println(Mon, Tuesd, Wed, Thur, Fir, Sat, Sun)
}
