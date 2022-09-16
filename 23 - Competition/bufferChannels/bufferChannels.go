package main

import "fmt"

func main() {
	channel := make(chan string, 2)

	channel <- "Hello"
	channel <- "Hello World"

	message := <- channel
	secondMessage := <- channel

	fmt.Println(message)
	fmt.Println(secondMessage)
}