package main

import (
	"fmt"
	"unicode" //unicode package is part of standard go package to work with unicode related staff with any build in methods
)

var ch int = '\u0041'
var ch2 int = '\u03B2'
var ch3 int = '\U00101234'

func main() {
	// format string specifier have been  used like %c for cahracter out put
	// format string specifier %u for out put the unicode format of the character that has  been given
	// format string specifier %v and %d will give the integer represention of that given value

	fmt.Println(unicode.IsLetter(rune(ch)))
	fmt.Println(unicode.IsNumber(rune(ch3)))
	fmt.Println(unicode.IsLower(rune(ch2)))
	fmt.Printf("%c - %c - %c\n", ch, ch2, ch3)
	fmt.Printf("%X - %X - %X\n", ch, ch2, ch3)
	fmt.Printf("%v - %v - %v\n", ch, ch2, ch3)
}
