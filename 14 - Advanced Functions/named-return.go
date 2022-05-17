package main

import "fmt"

func mathCal(n1, n2 int) (sum int, subtract int) {
	sum = n1 + n2
	subtract = n1 - n2
	return
}

func main() {
	sum, subtract := mathCal(10, 20)

	fmt.Println(sum, subtract)
}
