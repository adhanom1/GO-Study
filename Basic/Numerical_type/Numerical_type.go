package main
import "fmt"
var Integer_8bit uint8 = 255 // for signed integer 0-255
var Integer1_8bit int8 = -10 // for unsigend interger from -128 to 127
var Integer_16bit uint = 45236 // for unsigned integer 0-65536
var Integer1_16bit int = -4536525 // fot signed integer -32768 -> 32767
var Integer_32bit int = 1946565554 // for signed integer (− 2,147,483,648 to 2,147,483,647
var Integer1_32bit uint = 4100255225 // for unsigend integer 0-4,294,967,296 
//var Integer1_64bit int  will be 2 power of 64 signed integer value 
// var Integer1_64 uint will have 2 power 64 unsigned integer value
  var float float32 //(+- 1O-45 -> +- 3.4 * 1038 )
  var float2 float64 //(+- 5 * 10-324 -> 1.7 * 10308 )


	func main(){
		// can test the values from globaly declared variables 
	fmt.Println(Integer_8bit,Integer1_8bit)
	fmt.Println(float)
}