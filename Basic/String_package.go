package main

import (
	"fmt"
	"strings"
)

/*
	package strings has many function on this example

we will se three function that are basic like

		         hasprefix function with two input ( the string ,specific given value string we have been given)
				 hassuffix function with two input (the string , specific given value to compare with the string we have on hand )
				 contains  a function to search for specific character or substring on given string
	             index function return the sub string starting point
*/
func main() {
	st := "hello world how are you ?"
	fmt.Println("is the string %s has prefix of %s", st)
	fmt.Println(strings.HasPrefix(st, "he"))
	fmt.Println(strings.HasPrefix(st, "world"))
	fmt.Println(strings.HasSuffix(st, "?"))
	fmt.Println(strings.Contains(st, "how"))
	fmt.Println(strings.Index(st, "how"))
	fmt.Println(strings.LastIndex(st, "how"))
}
