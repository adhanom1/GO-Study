package main

import "fmt"

type vertex struct {
	first_name string
	last_name  string
	age        int
	salary     float64
}

func main() {

	v := vertex{"adhanom", "kidanmariam", 26, 15000}
	v.salary = 10000
	fmt.Println(v.first_name)
	fmt.Println(v.age)

}
