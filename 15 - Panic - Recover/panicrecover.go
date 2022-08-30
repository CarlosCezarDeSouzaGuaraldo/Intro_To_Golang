package main

import "fmt"

func restore() {
	if r := recover(); r != nil {
		fmt.Println("App restored")
	}
}

func studentQualified(n1, n2 float64) bool {
	defer restore()

	average := (n1 + n2) / 2

	if average > 6 {
		return true
	} else if average < 6 {
		return false
	}

	panic("The average is 6!")
}

func main() {
	fmt.Println(studentQualified(6, 6))
}
