package main

import "fmt"

func main() {
	sum := 1
	for i := 0; i < 10; i++ {
		sum = sum + 1

	}
	fmt.Println(sum)
	// with out puting init and post statement those two are optional
	for sum < 10 {
		sum += sum
	}
	fmt.Println(sum)
	/// for loop with init and post statment will act like while loop in c and c like language
	for sum < 10 {
		sum += sum
	}
	fmt.Println(sum)
}
