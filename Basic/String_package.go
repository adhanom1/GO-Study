package main

import (
	"fmt"
	"strings"
)

/*
	package strings has many function on this example

we will se three function that are basic like

					         Hasprefix function with two input ( the string ,specific given value string we have been given)
							 Hassuffix function with two input (the string , specific given value to compare with the string we have on hand )
							 Contains  a function to search for specific character or substring on given string
				             Index function return the sub string starting point
							 Count function for counting the occurences of the substring or character in string
							 Repeat function repeat the string given with given intger value return string
		                     ToLower to change the string in to small letter
							 ToUpper to chage the string in to all capital letter
	                         Trim to remove substrings from string
*/
func main() {
	st := "Hello World How are you ?"
	fmt.Printf("is the string %s has prefix of %s", st, strings.HasPrefix(st, "He"))
	fmt.Println(strings.HasPrefix(st, "He"))
	fmt.Println(strings.HasPrefix(st, "world"))
	fmt.Println(strings.HasSuffix(st, "?"))
	fmt.Println(strings.Contains(st, "how"))
	fmt.Println(strings.Index(st, "How"))
	fmt.Println(strings.LastIndex(st, "how"))
	fmt.Println(strings.Count(st, "l"))
	fmt.Println(strings.Repeat(st, 2))
	fmt.Println(strings.ToLower(st))
	fmt.Println(strings.ToUpper(st))
	fmt.Println(strings.Trim(st, "Hello World"))
}
