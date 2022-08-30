package main

import "fmt"

func general(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	general("String")
	general(2)
	general(4.52)
	general(false)
}
