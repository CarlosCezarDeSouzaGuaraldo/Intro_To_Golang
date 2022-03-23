package main

import "fmt"

type person struct {
	name string
	lastName string
	age int
	height int

}

type student struct {
	person
	course string
	university string
}

func main() {
	fmt.Println("Heritage")

	student1 := student{person{"Carlos", "Souza", 22, 175}, "ADS", "Fatec"}
	fmt.Println(student1)

	person2 := person{"Carlos", "Souza", 22, 175}
	student2 := student{person2, "ADS", "Fatec"}
	fmt.Println(student2)

	fmt.Println(student2.name)
	fmt.Println(student2.person.name)
	fmt.Println(student2.lastName)
	fmt.Println(student2.age)
}