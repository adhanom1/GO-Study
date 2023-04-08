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
	fmt.Println(Float, v, ints)
}
