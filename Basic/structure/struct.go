package main

import "fmt"

type vertex struct {
	X int
	y int
	z float64
}

func main() {

	v := vertex{6, 15, 36.5}
	v.X = 50
	fmt.Println(v)

}
