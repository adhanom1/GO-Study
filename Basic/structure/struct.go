package main

import "fmt"

func main() {
	type vertex struct {
		x int
		y int
		z float64
	}
	fmt.Println(vertex{6, 3, 36.5})

}
