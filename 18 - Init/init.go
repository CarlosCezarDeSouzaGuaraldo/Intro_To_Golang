package main

import "fmt"

var n int

func main() {
	fmt.Println("Running main function")
	fmt.Println(n)
}

func init() {
	fmt.Println("Running init function")
	n = 10
}
