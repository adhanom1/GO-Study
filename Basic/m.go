package main

import "fmt"

var Greeting string

func Hello(name string) string {
	return fmt.Sprintf(Greeting, name)
}
