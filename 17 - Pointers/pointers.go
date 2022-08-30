package main

import "fmt"

func changeSignal(number int) int {
	return number * -1
}

func changeSignalWithPointer(number *int) {
	*number = *number * -1
}

func main() {
	number := 20
	invertNumber := changeSignal(number)
	fmt.Println(invertNumber)
	fmt.Println(number)

	newNumber := 40
	fmt.Println(newNumber)
	changeSignalWithPointer(&newNumber)
	fmt.Println(newNumber)
}
