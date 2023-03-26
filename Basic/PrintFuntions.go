package main
import "fmt"
func main(){
	// print funtions with Printf global function from fmt package
	/* we can use different format string depending on the type we infered or use 
	bool:                    %t
    int, int8 etc.:          %d
    uint, uint8 etc.:        %d, %#x if printed with %#v
    float32, complex64, etc: %g
    string:                  %s
    chan:                    %p
    pointer:                 %p
	default value            %v
	 */
	name:= "james bond"
	salary:=102548.55
	comp := 10 + 32i
	ima :=10+20i
	paytax:=false
	var point *int64
	fmt.Printf("%p\n",point)
	fmt.Println(comp/ima,comp*ima,comp+ima,comp-ima)
	fmt.Printf("%g\n",comp)
    fmt.Printf("%s is %g years profit. he pay tax %t\n", name ,salary ,paytax)
	// %v format will give the default value given 
	fmt.Printf("%v is %v years profit. he pay tax %v\n", name ,salary ,paytax)
	// print function with Print
    movie:="Avatar"
	maked_year:=2009
	directed_by := "James Cameron"
	fmt.Print("movie name is ",movie," out for public at\n", maked_year ," directed by ", directed_by)
	//print function with Println
	fmt.Print("hello checkout with println there is big difference from both print function\n")
	fmt.Println("movie name is ",movie," out for public at", maked_year ," directed by ", directed_by)
}