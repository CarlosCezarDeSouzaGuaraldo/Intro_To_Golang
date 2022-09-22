package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type dog struct {
	Name  string `json:"name"`
	Breed string `json:"breed"`
	Age   uint   `json:"age"`
}

func main() {
	dogJson := `{"name":"Fred","breed":"PitBull","age":2}`

	var myDog dog

	if err := json.Unmarshal([]byte(dogJson), &myDog); err != nil {
		log.Fatal(err)
	}
	fmt.Println(myDog)

	otherDogJson := `{"name":"Toby", "breed":"Poodle"}`

	otherDog := make(map[string]string)
	if error := json.Unmarshal([]byte(otherDogJson), &otherDog); error != nil {
		log.Fatal(error)
	}
	fmt.Println(otherDog)
}
