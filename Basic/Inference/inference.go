package main
import (
	"fmt"
	"os"
)
func main(){
	//inferece or dynamic automatic type inference (:=)
	goos := os.Getenv("GOOS")
	fmt.Printf("the oprating system is : %s\n",goos)
	//inferece or dynamic automatic type inference (:=)
	path:=os.Getenv("PATH")
	fmt.Printf("path is %s\n",path)
}