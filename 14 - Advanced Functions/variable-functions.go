package main

import "fmt"

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func write(text string, numbers ...int){
	for _, number := range numbers{
		fmt.Println(text, number)
	}
}

func main() {
	value := sum(12, 14, 15)
	fmt.Println(value)

	value = sum()
	fmt.Println(value)

	write("Hello!", 12, 10, 1)
}
