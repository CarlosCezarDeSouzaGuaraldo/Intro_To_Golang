package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Loops")

	i := 0
	for i < 10 {
		fmt.Println("incrementando i")
		time.Sleep(time.Second)
		i++
	}

	fmt.Println(i)

	for j := 0; j < 10; j++ {
		fmt.Println("incrementando j", j)
		time.Sleep(time.Second)
	}

	names := [3]string{"Carlos", "Cezar", "Guaraldo"}

	for index, name := range names {
		fmt.Println(index, name)
	}

	names2 := []string{"Carlos", "Cezar", "Guaraldo"}
	for _, name := range names2 {
		fmt.Println(name)
	}

	for index, letter := range "Word" {
		fmt.Println(index, string(letter))
	}

	user := map[string]string{
		"name":     "Carlos",
		"lastName": "Guaraldo",
	}

	for index, value := range user {
		fmt.Println(index, value)
	}

	//infinite loop
	for {
		fmt.Println("executing infinitely")
		time.Sleep(time.Second)
	}
}
