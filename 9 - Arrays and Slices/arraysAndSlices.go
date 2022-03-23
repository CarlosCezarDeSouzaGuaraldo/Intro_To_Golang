package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays and Slices")

	var array1 [5]string
	array1[0] = "First position"
	fmt.Println(array1)

	array2 := [5]string{"First position", ""}
	fmt.Println(array2)

	array3 := [...]int{1,2,3,4,5}
	fmt.Println(array3)

	slice := []int{1,2,3}
	slice = append(slice, 4)
	fmt.Println(slice)

	slice2 := array3[1:3]
	fmt.Println(slice2)

	array3[1] = 12
	fmt.Println(slice2)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	//Intern Arrays
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

}
