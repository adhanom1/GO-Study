package main

import (
	"fmt"       // fmt is input and out put package in go
	"math/rand" // math/rand imported package
)

func main() {
	fmt.Println("the number that has been generated was ", rand.Float32())
	//fmt.Println(rand.float32())     //is error because with rand package if method or function is has to be exported it should be capital latter the first latter
	// this apply in identifer in any case like funtions , variables , slice
	fmt.Println("the number that been generated was ", rand.Intn(50))

}
