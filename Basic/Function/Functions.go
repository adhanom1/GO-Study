package main

import "fmt"

const a string = "hello am testing init fuction and compare it"

func init() { // init function is a function that run before the main function in the package
	fmt.Println(a)

}
func main() { // the main function is there the program written start executed
	fmt.Println("hello am from main function ")
	fmt.Println(add(5, 10))
	name, old := swap("James", 26)
	fmt.Println(name, old)
	fmt.Println(adds(15, 16))
}

func add(a, b int) int {
	return a + b

}
func swap(name string, age int) (string, int) {
	return name, age
}

func adds(a, b int) int {
	sum := a + b
	return sum
}
