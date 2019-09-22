package main

import (
	"fmt"
	"time"
)

// se estou dentro de uma goroutine o envio só
func rotina(c chan int) {
	time.Sleep(time.Second)
	c <- 1 // operação bloqueante

	fmt.Printf("Só depois que foi lido!")
}

func main() {
	c := make(chan int) // canal sem buffer

	go rotina(c)

	fmt.Println(<-c) // operação bloqueante
	fmt.Printf("Foi lido!\n")

	fmt.Println(<-c) // deadlock (vai esperar pra sempre, pois não tem go routine enviando dados)
	fmt.Printf("Fim")
}
