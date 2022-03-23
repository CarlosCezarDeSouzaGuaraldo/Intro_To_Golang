package main

import "fmt"

func main() {
	fmt.Println("Pointers")

	var variable1 int = 10
	var variable2 int = variable1

	fmt.Println(variable1, variable2)

	variable1++
	fmt.Println(variable1, variable2)

	//pointer is a reference memory
	var variable3 int
	var point *int

	variable3 = 100
	point = &variable3

	fmt.Println(variable3, point)

	variable3 = 150
	fmt.Println(variable3, *point)
}