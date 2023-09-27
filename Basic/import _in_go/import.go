package main

import (
	"fmt"       // fmt is input and out put package in go
	"math/rand" // math/rand imported package
)

func main() {
	fmt.Println("the number that has been generated was ", rand.Float32())
	fmt.Println("the number that been generated was ", rand.Intn(50))

}
