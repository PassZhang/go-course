package pkg01

import "fmt"

// public
var Name string = "good print pkg01"

//private
var version string = "1.0.0"

func init() {
	fmt.Println("Version: ", version)
}
