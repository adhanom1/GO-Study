package main

import "fmt"

var c, python, java string // package scope variable
var name = 42

// z:=12 this kind of declaration will not be working on package scope variable declaration
func main() { // function level variable must be used with in the function other wise they will be error
	var gender = "male"
	y := 5
	fmt.Println(gender)
	fmt.Println(y)

}
