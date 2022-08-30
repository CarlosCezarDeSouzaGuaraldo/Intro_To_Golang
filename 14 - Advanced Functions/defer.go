package main

import "fmt"

func function1() {
	fmt.Println("Executing 1 function")
}

func function2() {
	fmt.Println("Executing 2 function")
}

func studentQualified(n1, n2 float32) bool {
	defer fmt.Println("Result will be returned")

	average := (n1 + n2) / 2

	if average >= 6 {
		return true
	}

	return false
}

func main() {
	defer function1()
	function2()
	fmt.Println(studentQualified(7, 8))
}
