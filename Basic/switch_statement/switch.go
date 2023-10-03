package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	today := time.Now().Weekday()
	//fmt.Println(time.Now())
	fmt.Println(today)
	switch time.Saturday {
	case today + 0:
		fmt.Println("today is saturdat")
	case today + 1:
		fmt.Println("saturday is tommorrow")
	case today + 2:
		fmt.Println("saturday is the day after 2  days ")
	case today + 3:
		fmt.Println("saturday is the day after 3 days ")
	default:
		fmt.Println("saturday is far way ")
	}

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS,x")
	case "linux":
		fmt.Println("linux")
	case "windows":
		fmt.Println("windows")
	default:
		fmt.Printf("your os is %v", os)

	}
}
