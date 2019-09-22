package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtd int) {
	for i := 0; i < qtd; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	//fale("Maria", "Pq vc não fala cmg?", 3)
	//fale("João", "Só posso falar depois de você", 1)

	//go fale("Maria", "Pq vc não fala cmg?", 500)
	//go fale("João", "Só posso falar depois de você", 500)

	go fale("Maria", "Entendi!!!", 10)
	fale("João", "Parabéns", 5)

	fmt.Println("Fim")

	time.Sleep(time.Second * 5)
}
