package main
import "fmt"
const pi float32 = 3.14159
const billion int64 = 1e9
const name string = "hello there what is you name ?"
const imaginary complex64 = complex(4,3)
const (
	monday,tuesday, wenseday = 1,2,3
	thursday,friday, saturday = 4,5,6
)
// iota is incremented by one when ever the satement end and start with new line
// iota only be use in constant declaration scope 
const (
	unknown =iota
	female 
	male 
	i int = iota + 15
	s 
)
func main(){
	fmt.Println(i)
	fmt.Println("constant value",pi)
	fmt.Println("constant value",billion)
	fmt.Println("constant value",name)
	fmt.Println("constant value",imaginary)
	fmt.Println(monday,tuesday,wenseday,thursday,friday,saturday)
	fmt.Println(unknown,male,female,s)
}