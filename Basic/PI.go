package main
import "math"
import "fmt"
var pi float64
func init(){
	pi=4*math.Atan(1) //compute the value of 
	
}
func main(){
	two_pi:=2*pi 
	fmt.Printf("the value is %g\n",two_pi)
}