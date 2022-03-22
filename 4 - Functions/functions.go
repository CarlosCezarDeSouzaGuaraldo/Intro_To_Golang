package main

import "fmt"

func sum(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2

	return soma, subtracao
}

func main() {
	soma := sum(10, 20)
	fmt.Println(soma)

	var f = func() {
		fmt.Println("Function F")
	}

	f()

	resultadoSoma, _ := calculos(10, 15)
	fmt.Println(resultadoSoma)

	resultadoSoma, resultadoSubtracao := calculos(10, 15)
	fmt.Println(resultadoSoma, resultadoSubtracao)
}
