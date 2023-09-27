package main

import (
	"fmt"
	"math/rand"

	"news.com/event/Casting"
)

var z int = 15
var b string = "adhanom"

type (
	hotdog int
	race   string
)

func main() {
	fmt.Println("my favorite number is ", rand.Intn(100))

	var A string = "adhanom"
	n, err := fmt.Println("hello world")
	fmt.Println(n)
	fmt.Println(err)
	fmt.Printf("the name is %v\n", A)
	fmt.Printf("%#v\n", A)
	Casting.Adhanom()
	var ad hotdog = 50
	var animal race = "lion"
	fmt.Printf("%t\n", animal)
	fmt.Println(ad)
	fmt.Printf("%T\n", ad)
	fmt.Println(animal)
	fmt.Printf("%T\n", animal)
	fmt.Println(z)
	fmt.Println(b)
	b = string(z)
	fmt.Println(b)

}
