package main

import "fmt"

func main() {
	defer fmt.Println("hello world ")
	fmt.Println("am there first you are using defer ")
}
