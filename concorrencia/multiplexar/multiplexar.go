package main

import (
	"fmt"

	"github.com/cod3rcursos/html"
)

func encaminhar(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem // <-origem é bloqueante, vai aguardar até ter dado para ler do canal
	}
}

func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)

	return c
}

func main() {
	c := juntar(
		html.Titulo("https://www.apple.com", "https://www.google.com"),
		html.Titulo("https://pt.aliexpress.com/", "https://www.udemy.com"),
	)

	fmt.Println(<-c, "|", <-c)
	fmt.Println(<-c, "|", <-c)
}
