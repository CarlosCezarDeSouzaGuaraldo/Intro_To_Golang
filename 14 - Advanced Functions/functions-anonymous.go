package main

import "fmt"

func main() {

	func() {
		fmt.Println("Hey!")
	}()

	func(text string) {
		fmt.Println(text)
	}("Hey!")

	message := func(text string) string {
		return fmt.Sprintf("Received -> %s", text)
	}("Sending parameter")

	fmt.Println(message)
}
