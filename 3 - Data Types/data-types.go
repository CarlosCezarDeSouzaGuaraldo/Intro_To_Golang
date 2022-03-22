package main

import (
	"errors"
	"fmt"
)

func main() {
	//int8, int16, int32, int64

	number := 100000000000000000
	fmt.Println(number)

	var number2 uint = 15
	fmt.Println(number2)

	//alias
	//int32 = rune
	var number3 rune = 12
	fmt.Println(number3)

	//byte = uint8
	var number4 byte = 123
	fmt.Println(number4)

	var realNumber float32 = 123.45
	var realNumber2 float64 = 123456.789
	fmt.Println(realNumber, realNumber2)

	realNumber3 := 987654.321
	fmt.Println(realNumber3)

	var str string = "tests with string!"
	fmt.Println(str)

	str2 := "test 2"
	fmt.Println(str2)

	//Return the character in ASCII sheet
	char := 'B'
	fmt.Println(char)

	//Fim Strings
	var text string
	fmt.Println(text)	

	var integer int
	fmt.Println(integer)
	
	var boo bool
	fmt.Println(boo)

	var erro error
	fmt.Println(erro)

	var err error = errors.New("Intern Error")
	fmt.Println(err)
}
