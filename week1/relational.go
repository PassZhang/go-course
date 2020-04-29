package main 

import "fmt"

func main() {
	var a int = 21
	var b int = 10 

	if ( a == b ) {
		fmt.Printf(" a = b ")
	} else {
		fmt.Printf(" a != b ")
	} 
	
	if ( a < b ) {
		fmt.Printf(" a < b ")
	} else {
		fmt.Printf(" a 不小于 b ")
	}
	
	if ( a > b ) {
		fmt.Printf(" a > b ")
	} else {
		fmt.Printf(" a 不大于 b ")
	}

	b = 20 
	a = 5
	if ( a <= b ) {
		fmt.Printf(" a <= b ")
	} 
	
	if ( b >= a ) {
		fmt.Printf(" b >= a ")
	}
}