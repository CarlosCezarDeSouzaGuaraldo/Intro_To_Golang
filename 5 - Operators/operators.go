package main

import (
	"fmt"
)

func main() {
	sum := 1 + 1
	subtract := 1 - 1
	division := 1 / 1
	multiply := 1 * 1
	rest := 1 % 2

	fmt.Println(sum, subtract, division, multiply, rest)

	var n1 int32 = 10
	var n2 int32 = 25

	soma := n1 + n2
	fmt.Println(soma)

	var variable string = "String"
	variable2 := "String2"

	fmt.Println(variable, variable2)

	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	tru, fals := true, false
	fmt.Println(tru && fals)
	fmt.Println(tru || fals)
	fmt.Println(!tru)

	num := 10
	num++
	num += 15
	fmt.Println(num)

	num--
	num -= 15
	num *= 2
	num /= 2
	fmt.Println(num)

	var text string
	if num > 5 {
		text = "More than 5"
	} else {
		text = "Less than 5"
	}
	fmt.Println(text)
}
