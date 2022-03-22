package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Writing in main file")
	auxiliar.Write()

	verifyngEmail := checkmail.ValidateFormat("carloscezardesouza@outlook.com")
	fmt.Println(verifyngEmail)
}
