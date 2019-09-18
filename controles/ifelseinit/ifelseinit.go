package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	nanoSeconds := time.Now().UnixNano()
	source := rand.NewSource(nanoSeconds)
	r := rand.New(source)
	return r.Intn(10)
}

func main() {

	if i := numeroAleatorio(); i > 5 {
		fmt.Println("Ganhou!!!", i)
	} else {
		fmt.Println("Perdeu!!!", i)
	}

	// fmt.Println(i) // não está disponivel

}
