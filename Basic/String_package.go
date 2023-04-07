package main

import (
	"fmt"
	"strings"
)

func main() {
	st := "hello world how are you ?"
	fmt.Println("is the string %s has prefix of %s", st)
	fmt.Println(strings.HasPrefix(st, "he"))
	fmt.Println(strings.HasPrefix(st, "world"))
	fmt.Println(strings.HasSuffix(st, "?"))
}
