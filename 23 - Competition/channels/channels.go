package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)

	go write("Let's GO!", channel)

	for message := range channel {
		fmt.Println(message)
	}

	fmt.Println("End")
}

func write(text string, channel chan string) {
	for i := 0; i < 5; i++ {
		channel <- text
		time.Sleep(time.Second)
	}

	close(channel)
}
