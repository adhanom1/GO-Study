package main

import "fmt"

func main() {
	defer fmt.Println("hello world ")
	fmt.Println("am there first you are using defer ")
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)

	}

	fmt.Println("done")
}
