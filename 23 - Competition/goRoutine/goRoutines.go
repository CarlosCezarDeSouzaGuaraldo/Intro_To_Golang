package main

import (
	"fmt"
	"time"
)

func main() {
	// competition != parallelism
	go write("Hello World!")
	write("Learning Go!")
}

func write(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
