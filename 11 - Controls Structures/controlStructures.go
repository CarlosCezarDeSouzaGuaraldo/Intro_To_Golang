package main

import "fmt"

func main() {
	fmt.Println("Control's Structure")

	number := 10
	if number > 15 {
		fmt.Println("More than 15")
	} else {
		fmt.Println("Less than 15")
	}

	if otherNumber := number; otherNumber > 0 {
		fmt.Println("Number more than 0")
	} else if number < -10 {
		fmt.Println("Number less than 0")
	}
}
