package Casting

import "fmt"

type (
	i int
	f float32
	s string
)

func Adhanom() {
	a := 10.55
	b := int(a)
	f := 4141.32
	i := int(f)
	s := "hello there"
	fmt.Println("the value is float:", a)
	fmt.Println("the value is casted in to intger:", b)
	fmt.Println("value of type f is float before casting ", f)
	fmt.Println("value of the type i is interger ", i)
	fmt.Println(s)

}
