package main

import (
	"fmt"
	"math"
)

var pi float64

func init() {
	pi = 4 * math.Atan(1) //compute the value of

}
func main() {
	two_pi := 2 * pi
	fmt.Printf("the value is %g\n", two_pi)
	intr := 10.5
	ints := "hello"
	var Float *string
	Float = &ints
	v := *Float
	v = "adhanom"
	fmt.Println(&intr, &ints)
	fmt.Println(Float, *Float, v, ints)

	var i1 = 5
	fmt.Printf("An integer: %d, its location in memory: %p\n", i1, &i1)
	var intP *int
	intP = &i1
	fmt.Printf("The value at memory location %p is %d\n", intP, *intP)

}
