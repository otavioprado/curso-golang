package main

import (
	"fmt"
	"time"
)

func doisTresQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second * 5)

	c <- 2 * base // enviando dados para o canal (operação bloqueante, só desbloquea depois que alguém ler o valor do canal!)

	time.Sleep(time.Second)
	c <- 3 * base

	fmt.Println("Inserido dois valores!")

	time.Sleep(time.Second)
	c <- 4 * base
}

// Channel é a forma de comunicação entre goroutines
func main() {
	c := make(chan int)
	go doisTresQuatroVezes(2, c)

	fmt.Println("Aguardando...")
	a, b := <-c, <-c // recebendo dados do canal
	fmt.Println("Só vai printar isso aqui quando tiver dois dados para consumir do channel")
	fmt.Println(a, b)
	fmt.Println(<-c)

	time.Sleep(time.Second * 10)
}
