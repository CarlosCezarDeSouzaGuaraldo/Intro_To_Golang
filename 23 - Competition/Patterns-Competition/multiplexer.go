package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	channel := multiplexer(write("Hello!"), write("Bye!"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-channel)
	}
}

func multiplexer(firstInChannel, secInChannel <-chan string) <-chan string {
	outChannel := make(chan string)

	go func() {
		for {
			select {
			case message := <-firstInChannel:
				outChannel <- message
			case message := <-secInChannel:
				outChannel <- message
			}
		}
	}()
	return outChannel
}

func write(text string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("Text received: %s", text)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return channel
}
