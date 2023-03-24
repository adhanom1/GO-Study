package main
import "fmt"
//global variable can be accessed with in the package every where 
var a,b int
var (
	str string 
	float float32
	imaginary complex64
	d bool
	i int
	
)

func main(){
//local variable can be accesed only with in local scope like function , for loop , if statement 
    i=2
    a=10+i
	b=15
	float = 3.788554
	str="hello there"
	imaginary = complex(3,8)
	d=false
	fmt.Println(a,b)
	fmt.Println(a)
	fmt.Println(float,imaginary,d)
	fmt.Println(str)
}