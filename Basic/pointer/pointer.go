package main

import "fmt"

var p *int

func main() {

	i := 40
	p = &i
	fmt.Println(*p)
}
