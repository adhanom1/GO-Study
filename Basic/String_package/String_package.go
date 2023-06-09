package main

import (
	"fmt"
	"strings"
	"time"
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
	                             Field function
								 Split  function
*/
func main() {
	T := time.Now()
	st := "Hello World | How are you ?"
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
	fmt.Println(strings.Trim(st, "\r\n"))
	str := strings.Fields(st)
	for _, val := range str {
		fmt.Printf("%s -", val)
	}
	fmt.Println()
	str2 := strings.Split(st, "|")
	fmt.Println(str2)
	for _, val2 := range str2 {
		fmt.Printf("%s -", val2)
	}
	fmt.Println()
	str3 := strings.Join(str2, "james bond")
	fmt.Printf("sl2 joined by ;: %s\n", str3)
	i := "99.5"
	j := string(i)
	fmt.Println(j)
	T2 := time.Minute
	fmt.Println(T)
	fmt.Println(T2)

}
