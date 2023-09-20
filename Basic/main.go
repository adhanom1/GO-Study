package main

import (
	"fmt"

	"news.com/event/Casting"
)

type (
	hotdog int
	race   string
)

func main() {
	var A string = "adhanom"
	n, err := fmt.Println("hello world")
	fmt.Println(n)
	fmt.Println(err)
	fmt.Printf("the name is %v\n", A)
	fmt.Printf("%#v\n", A)
	Casting.Adhanom()
	var ad hotdog = 50
	var animal race = "lion"
	animal := string(ad)
	fmt.Printf("%t\n", animal)
	fmt.Println(ad)
	fmt.Printf("%T\n", ad)
	fmt.Println(animal)
	fmt.Printf("%T\n", animal)

}
