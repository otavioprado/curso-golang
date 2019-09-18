package main

import (
	"fmt"
)

func main() {
	i := 1

	// Go não tem aritmética de ponteiros
	fmt.Println(i)

	var ponteiro *int = nil
	ponteiro = &i
	fmt.Println(ponteiro)  // endereço de memória
	fmt.Println(*ponteiro) // valor de onde o ponteiro aponta

	*ponteiro++
	i++

	// p++ // não funciona

	fmt.Println(ponteiro, *ponteiro, i)
}
