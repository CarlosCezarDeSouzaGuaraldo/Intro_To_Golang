package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type dog struct {
	Name string `json:"name"`
	Breed string `json:"breed"`
	Age uint `json:"age"`
}

func main() {
	myDog := dog{"Fred", "PitBull", 2}
	fmt.Println(myDog)

	myDogJson, err := json.Marshal(myDog)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(myDogJson)
	fmt.Println(bytes.NewBuffer(myDogJson))

	otherDog := map[string]string {
		"name": "Toby",
		"breed": "Poodle",
	}

	otherDogJson, err := json.Marshal(otherDog)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(otherDogJson)
	fmt.Println(bytes.NewBuffer(otherDogJson))
}