package main
import "fmt"
// pointer declaration
var number,decimal *int

func main(){
// reference data type 
	var flat float64 = 35.877878
	var i int16 = 10
	k:=&flat
	var j = &i
    num := &number
	dec := &decimal
	fmt.Println(flat,k)
	fmt.Println(i,j)
	fmt.Println(num,dec)
	
}