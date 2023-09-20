package main

import "fmt"

func main() {
	/* the numerical type are integer , floating and complex value .
	on the contect of integer there are two different catagories
	# unsigned integer
	# signed integer
	the difference between then this the use of values contaning negeative values or haveing used all
	the values for positive numbers like uint8 is from -128 to 128 but if we used int8 is all value
	from 0 to 256
	*/
	x := 42
	y := 45.66
	var w int8 = 127
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(w)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
}
