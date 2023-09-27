package main

import (
	"fmt" // fmt is input and out put package in go
	"math"
	"math/rand" // math/rand imported package
)

func main() {
	fmt.Println("the number that has been generated was ", rand.Float32())
	//fmt.Println(rand.float32())     //is error because with rand package if method or function is has to be exported it should be capital latter the first latter
	// this apply in identifer in any case like funtions , variables , slice
	fmt.Println("the number that been generated was ", rand.Intn(50))
	fmt.Println(math.Cos(45)) // as you have seen Cos is function in math package also imported because it is visible out side the package scope because of the function start the name with capital latter
	fmt.Println(math.Exp(1))  // exponent e power 1 value
}
