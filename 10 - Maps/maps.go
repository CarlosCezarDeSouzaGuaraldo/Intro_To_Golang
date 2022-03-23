package main

import "fmt"

func main() {
	fmt.Println("Maps")

	user := map[string]string{
		"name":     "Carlos",
		"lastname": "Guaraldo",
	}

	fmt.Println(user)
	fmt.Println(user["name"])

	user2 := map[string]map[string]string{
		"name": {
			"firstName": "Carlos",
			"lastName":  "Carlos",
		},
		"course": {
			"name": "ADS",
		},
	}
	fmt.Println(user2)

	delete(user2, "name")
	fmt.Println(user2)

	user2["country"] = map[string]string{
		"name": "Brazil",
	}
	fmt.Println(user2)
}
