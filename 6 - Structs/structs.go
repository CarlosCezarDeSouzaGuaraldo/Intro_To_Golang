package main

import "fmt"

type usuario struct {
	name string
	age  int8
	address address
}

type address struct {
	logradouro string
	number uint8
}

func main() {

	fmt.Println("Structs file")

	var user usuario
	user.name = "Carlos"
	user.age = 22
	fmt.Println(user)

	address1 := address{"Street", 123}

	user2 := usuario{"Carlos", 22, address1}
	fmt.Println(user2)

	user3 := usuario{name: "Carlos"}
	fmt.Println(user3)

	user4 := usuario{age: 22}
	fmt.Println(user4)

	user5 := usuario{"Carlos", 22, address{"My street", 0123}}
	fmt.Println(user5)
}
